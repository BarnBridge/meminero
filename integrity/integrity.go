package integrity

import (
	"context"
	"sort"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/state"
	"github.com/lacasian/ethwheels/bestblock"
)

var (
	metricsHighestDBBlock = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "scraper_highest_db_block",
		Help: "Highest block number saved in the db",
	})
	metricsHighestChainBlock = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "scraper_highest_chain_block",
		Help: "Head of chain block number",
	})
	metricsCheckpointBlock = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "scraper_integrity_checkpoint_block",
		Help: "Block up to which integrity of chain is confirmed",
	})
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
			err := c.lifecycle(ctx)
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

func (c *Checker) lifecycle(ctx context.Context) error {
	c.logger.Trace("running")
	start := time.Now()
	defer func() {
		c.logger.WithField("duration", time.Since(start)).Trace("done")
	}()

	best := c.tracker.BestBlock()
	checkpoint, err := c.getLastCheckpoint(ctx)
	if err != nil {
		return err
	}
	if checkpoint == -1 {
		return nil
	}

	var highestBlock int64
	err = c.db.QueryRow(ctx, `select max(number) from blocks;`).Scan(&highestBlock)
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

	metricsCheckpointBlock.Set(float64(checkpoint))
	metricsHighestChainBlock.Set(float64(best))
	metricsHighestDBBlock.Set(float64(highestBlock))

	if checkpoint >= highestBlock {
		c.logger.Warn("checkpoint is higher than highest block; there's nothing to check")
		return nil
	}

	missing, err := c.checkMissingBlocks(ctx, checkpoint, highestBlock)
	if err != nil {
		return err
	}

	broken, err := c.checkBrokenHashChain(ctx, checkpoint, highestBlock)
	if err != nil {
		return err
	}

	all := append(missing, broken...)
	if len(all) == 0 {
		_, err = c.db.Exec(ctx, "insert into public.integrity_checkpoints (number) values($1)", highestBlock)
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

	_, err = c.db.Exec(ctx, "insert into public.integrity_checkpoints (number) values($1)", blocks[0]-1)
	if err != nil {
		return errors.Wrap(err, "could not store new integrity checkpoint")
	}

	c.logger.WithField("count", len(blocks)).Warn("found inconsistent blocks & queued for rescrape")

	return nil
}

func (c *Checker) getLastCheckpoint(ctx context.Context) (int64, error) {
	var b int64
	err := c.db.QueryRow(ctx, `select number from public.integrity_checkpoints order by created_at desc limit 1`).Scan(&b)
	if err == pgx.ErrNoRows {
		err1 := c.db.QueryRow(ctx, `select min(number) from blocks`).Scan(&b)
		if err1 == pgx.ErrNoRows {
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
