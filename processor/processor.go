package processor

import (
	"context"
	"database/sql"
	"time"

	"github.com/barnbridge/smartbackend/state"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/types"
)

type Processor struct {
	Raw   *types.RawData
	Block *types.Block

	eth   *ethclient.Client

	logger *logrus.Entry

	storables []types.Storable
}

func New(raw *types.RawData, eth *ethclient.Client) (*Processor, error) {
	p := &Processor{
		Raw:    raw,
		eth: eth,
		logger: logrus.WithField("module", "processor"),
	}

	err := p.preprocess()
	if err != nil {
		return nil, err
	}

	err = state.Refresh()
	if err != nil {
		return nil, errors.Wrap(err, "could not refresh state")
	}
	p.registerStorables()

	return p, nil
}

func (p *Processor) rollbackAll(db  *pgxpool.Pool) error {
	tx, err := db.BeginTx(context.Background(),pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "could not start database transaction")
	}

	_, err = tx.Exec(context.Background(),"delete from blocks where number = $1", p.Block.Number)
	if err != nil {
		return errors.Wrap(err, "could not remove block from database")
	}

	for _, s := range p.storables {
		err = s.Rollback(tx)
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not commit rollback transaction")
	}

	p.logger.WithField("block", p.Block.Number).Info("removed old version from the db; will be replaced with new version")

	return nil
}

// Store will open a database transaction and execute all the registered Storables in the said transaction
func (p *Processor) Store(db  *pgxpool.Pool) error {
	exists, err := p.checkBlockExists(db)
	if err != nil {
		return err
	}

	if exists && !config.Store.Feature.ReplaceBlocks {
		p.logger.Info("block already exists in the database; skipping")
		return nil
	}

	if config.Store.Feature.ReplaceBlocks {
		p.logger.WithField("block", p.Block.Number).Warn("removing any old versions of block from db because feature flag is enabled")

		err = p.rollbackAll(db)
		if err != nil {
			return err
		}
	} else {
		reorged, err := p.checkBlockReorged(db)
		if err != nil {
			return err
		}

		if reorged {
			p.logger.WithField("block", p.Block.Number).Warn("detected reorged block")

			err = p.rollbackAll(db)
			if err != nil {
				return err
			}
		}
	}

	for _, s := range p.storables {
		err = s.Execute()
		if err != nil {
			return err
		}
	}

	err = p.storeAll(db)
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) storeAll(db *pgxpool.Pool) error {
	tx, err := db.BeginTx(context.Background(),pgx.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "could not start database transaction")
	}

	err = p.storeBlock(tx)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	for _, s := range p.storables {
		err = s.SaveToDatabase(tx)
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not save data to db")
	}

	return nil
}

func (p *Processor) storeBlock(tx pgx.Tx) error {
	p.logger.Trace("storing block")
	start := time.Now()
	defer func() { p.logger.WithField("duration", time.Since(start)).Debug("done storing block") }()

	stmt, err := tx.Prepare(context.Background(),pq.CopyIn("blocks", "number", "block_hash", "parent_block_hash", "block_creation_time"))
	if err != nil {
		return err
	}

	b := p.Block
	err := tx.
	_, err = stmt.Exec(b.Number, b.BlockHash, b.ParentBlockHash, b.BlockCreationTime)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}
