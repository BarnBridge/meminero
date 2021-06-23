package processor

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/types"
)

type Processor struct {
	Raw   *types.RawData
	Block *types.Block

	logger *logrus.Entry

	storables []types.Storable
}

func New(raw *types.RawData) (*Processor, error) {
	p := &Processor{
		Raw:    raw,
		logger: logrus.WithField("module", "processor"),
	}

	err := p.preprocess()
	if err != nil {
		return nil, err
	}

	p.registerStorables()

	return p, nil
}

func (p *Processor) rollbackAll(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "could not start database transaction")
	}

	_, err = tx.Exec("delete from blocks where number = $1", p.Block.Number)
	if err != nil {
		return errors.Wrap(err, "could not remove block from database")
	}

	for _, s := range p.storables {
		err = s.Rollback(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "could not commit rollback transaction")
	}

	p.logger.WithField("block", p.Block.Number).Info("removed old version from the db; will be replaced with new version")

	return nil
}

// Store will open a database transaction and execute all the registered Storables in the said transaction
func (p *Processor) Store(db *sql.DB) error {
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

func (p *Processor) storeAll(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "could not start database transaction")
	}

	err = p.storeBlock(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, s := range p.storables {
		err = s.SaveToDatabase(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "could not save data to db")
	}

	return nil
}

func (p *Processor) storeBlock(tx *sql.Tx) error {
	p.logger.Trace("storing block")
	start := time.Now()
	defer func() { p.logger.WithField("duration", time.Since(start)).Debug("done storing block") }()

	stmt, err := tx.Prepare(pq.CopyIn("blocks", "number", "block_hash", "parent_block_hash", "block_creation_time"))
	if err != nil {
		return err
	}

	b := p.Block

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
