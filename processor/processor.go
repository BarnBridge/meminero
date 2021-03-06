package processor

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/state"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/types"
)

type Processor struct {
	Raw   *types.RawData
	Block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	storables []types.Storable
}

func New(raw *types.RawData, state *state.Manager) (*Processor, error) {
	p := &Processor{
		Raw:    raw,
		state:  state,
		logger: logrus.WithField("module", "processor"),
	}

	err := p.preprocess()
	if err != nil {
		return nil, err
	}

	p.registerStorables()

	return p, nil
}

func (p *Processor) rollbackAll(ctx context.Context, db *pgxpool.Pool) error {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "could not start database transaction")
	}

	_, err = tx.Exec(ctx, "delete from blocks where number = $1", p.Block.Number)
	if err != nil {
		return errors.Wrap(err, "could not remove block from database")
	}

	for _, s := range p.storables {
		var log = logrus.WithField("module", fmt.Sprintf("storable(%s)", s.ID()))

		log.Trace("rolling back block")
		start := time.Now()

		err = s.Rollback(ctx, tx)
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}

		recordRollbackDuration(s.ID(), start)
		log.WithField("duration", time.Since(start)).Trace("done rolling back block")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.Wrap(err, "could not commit rollback transaction")
	}

	p.logger.WithField("block", p.Block.Number).Info("removed old version from the db; will be replaced with new version")

	return nil
}

// Store will open a database transaction and execute all the registered Storables in the said transaction
func (p *Processor) Store(ctx context.Context, db *pgxpool.Pool) (bool, error) {
	exists, err := p.checkBlockExists(ctx, db)
	if err != nil {
		return false, err
	}

	if exists && !config.Store.Feature.ReplaceBlocks {
		p.logger.Infof("block %d (%s) already exists in the database; skipping", p.Block.Number, p.Block.BlockHash)
		return false, nil
	}

	if config.Store.Feature.ReplaceBlocks {
		p.logger.WithField("block", p.Block.Number).Warn("removing any old versions of block from db because feature flag is enabled")

		err = p.rollbackAll(ctx, db)
		if err != nil {
			return false, err
		}
	} else {
		reorged, err := p.checkBlockReorged(ctx, db)
		if err != nil {
			return false, err
		}

		if reorged {
			p.logger.WithField("block", p.Block.Number).Warn("detected reorged block")

			err = p.rollbackAll(ctx, db)
			if err != nil {
				return false, err
			}
		}
	}

	start := time.Now()
	p.logger.Info("executing storables")

	wg, _ := errgroup.WithContext(ctx)

	for _, s := range p.storables {
		s := s

		wg.Go(func() error {
			log := logrus.WithField("module", fmt.Sprintf("storable(%s)", s.ID()))

			log.Trace("executing")
			start := time.Now()

			err = s.Execute(ctx)
			if err != nil {
				return err
			}

			recordExecuteDuration(s.ID(), start)
			log.WithField("duration", time.Since(start)).Trace("done executing")

			return nil
		})
	}

	err = wg.Wait()
	if err != nil {
		return false, errors.Wrap(err, "got error executing storables")
	}

	p.logger.WithField("duration", time.Since(start)).Info("done executing storables")

	err = p.storeAll(ctx, db)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Processor) storeAll(ctx context.Context, db *pgxpool.Pool) error {
	start := time.Now()
	p.logger.Info("storing data to database")
	defer func() {
		p.logger.WithField("duration", time.Since(start)).Info("done storing data to database")
	}()

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "could not start database transaction")
	}

	err = p.storeBlock(ctx, tx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	for _, s := range p.storables {
		log := logrus.WithField("module", fmt.Sprintf("storable(%s)", s.ID()))

		log.Trace("saving")

		start := time.Now()

		err = s.SaveToDatabase(ctx, tx)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}

		recordSaveDuration(s.ID(), start)
		log.WithField("duration", time.Since(start)).Trace("done saving")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.Wrap(err, "could not save data to db")
	}

	return nil
}

func (p *Processor) storeBlock(ctx context.Context, tx pgx.Tx) error {
	p.logger.Trace("storing block")
	start := time.Now()
	defer func() { p.logger.WithField("duration", time.Since(start)).Trace("done storing block") }()

	b := p.Block

	_, err := tx.Exec(ctx, "insert into blocks(number,block_hash,parent_block_hash,block_creation_time) values($1,$2,$3,$4)", b.Number, b.BlockHash, b.ParentBlockHash, b.BlockCreationTime)
	if err != nil {
		return err
	}

	return nil
}
