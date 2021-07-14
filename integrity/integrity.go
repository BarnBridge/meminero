package integrity

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/state"
	"github.com/lacasian/ethwheels/bestblock"
)

type Checker struct {
	db      *pgxpool.Pool
	tracker *bestblock.Tracker
	tm      *state.Manager
	logger  *logrus.Entry
}

func NewChecker(db *pgxpool.Pool, tracker *bestblock.Tracker, tm *state.Manager) *Checker {
	return &Checker{
		db:      db,
		tracker: tracker,
		tm:      tm,
		logger:  logrus.WithField("module", "integrity-checker"),
	}
}

func (c *Checker) Run(ctx context.Context) {
	t := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-t.C:
			err := c.lifecycle()
			if err != nil {
				c.logger.Error(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (c *Checker) getLag() int64 {
	return config.Store.Feature.QueueKeeper.Lag
}

func (c *Checker) lifecycle() error {
	c.logger.Trace("running")
	start := time.Now()
	defer func() {
		c.logger.WithField("duration", time.Since(start)).Trace("done")
	}()

	best := c.tracker.BestBlock()
	checkpoint, err := c.getLastCheckpoint()
	if err != nil {
		return err
	}
	if checkpoint == -1 {
		return nil
	}

	var highestBlock int64
	err = c.db.QueryRow(context.Background(), `select max(number) from blocks;`).Scan(&highestBlock)
	if err != nil {
		return errors.Wrap(err, "could not fetch highest block from database")
	}

	if highestBlock < best-c.getLag()-10 {
		c.logger.WithFields(logrus.Fields{
			"highest-db":    highestBlock,
			"highest-chain": best,
			"diff":          best - highestBlock,
		}).Error("pipeline is falling behind")
	}

	if checkpoint >= highestBlock {
		c.logger.Warn("checkpoint is higher than highest block; there's nothing to check")
		return nil
	}

	missing, err := c.checkMissingBlocks(checkpoint, highestBlock)
	if err != nil {
		return err
	}

	broken, err := c.checkBrokenHashChain(checkpoint, highestBlock)
	if err != nil {
		return err
	}

	all := append(missing, broken...)
	if len(all) == 0 {
		_, err = c.db.Exec(context.Background(), "insert into integrity_checkpoints (number) values($1)", highestBlock)
		if err != nil {
			return errors.Wrap(err, "could not store new integrity checkpoint")
		}

		c.logger.Info("finished checking integrity; all good!")
		return nil
	}

	var uniqueBlocks = make(map[int64]bool)

	for _, block := range all {
		uniqueBlocks[block] = true
	}

	var blocks []int64
	for k := range uniqueBlocks {
		blocks = append(blocks, k)
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i] < blocks[j]
	})

	for _, block := range blocks {
		err = c.tm.AddTaskToQueue(block)
		if err != nil {
			return errors.Wrap(err, "could not queue block for rescrape")
		}
	}

	_, err = c.db.Exec(context.Background(), "insert into integrity_checkpoints (number) values($1)", blocks[0]-1)
	if err != nil {
		return errors.Wrap(err, "could not store new integrity checkpoint")
	}

	c.logger.WithField("count", len(blocks)).Warn("found inconsistent blocks & queued for rescrape")

	return nil
}

func (c *Checker) getLastCheckpoint() (int64, error) {
	var b int64
	err := c.db.QueryRow(context.Background(), `select number from integrity_checkpoints order by created_at desc limit 1`).Scan(&b)
	if err == sql.ErrNoRows {
		err1 := c.db.QueryRow(context.Background(), `select min(number) from blocks`).Scan(&b)
		if err1 == sql.ErrNoRows {
			return -1, nil
		}
		if err1 != nil {
			return 0, errors.Wrap(err, "could not get min block number from db")
		}

		return b, nil
	}
	if err != nil {
		return 0, errors.Wrap(err, "could not get latest integrity checkpoint from db")
	}

	return b, nil
}
