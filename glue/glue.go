package glue

import (
	"context"
	"sync"
	"time"

	"github.com/alethio/web3-go/validator"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/processor"
	"github.com/barnbridge/meminero/scraper"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

var (
	metricsBlocksProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "scraper_processed_blocks",
		Help: "Number of blocks scraped and saved",
	})
	metricsBlocksSkipped = promauto.NewCounter(prometheus.CounterOpts{
		Name: "scraper_skipped_blocks",
		Help: "Number of blocks that have been processed but not saved",
	})
	metricsBlocksErrored = promauto.NewCounter(prometheus.CounterOpts{
		Name: "scraper_errored_blocks",
		Help: "Number of blocks errored and re-queued",
	})
	metricsScrapingDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "scraper_scrape_duration_ms",
		Help:    "How long did it take to scrape the data from the node",
		Buckets: []float64{100, 500, 1000, 2000, 4000, 8000},
	})
	metricsProcessingDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "scraper_processing_duration_ms",
		Help:    "How long did it take to process the data",
		Buckets: []float64{1, 10, 50, 100, 500, 1000, 2000, 4000},
	})
)

type Glue struct {
	state   *state.Manager
	scraper *scraper.Scraper
	db      *pgxpool.Pool
	logger  *logrus.Entry

	stopMu sync.Mutex
}

func New(db *pgxpool.Pool, state *state.Manager) (*Glue, error) {
	logger := logrus.WithField("module", "glue")

	s, err := scraper.New()
	if err != nil {
		return nil, errors.Wrap(err, "could not init scraper")
	}

	return &Glue{
		state:   state,
		scraper: s,
		db:      db,
		logger:  logger,
	}, nil
}

func (g *Glue) ScrapeSingleBlock(ctx context.Context, b int64) (bool,error) {
	log := g.logger.WithField("block", b)
	log.Info("processing block")

	start := time.Now()
	blk, err := g.scraper.Exec(b)
	if err != nil {
		return false, errors.Wrap(err, "could not scrape block")
	}

	_, err = g.validateBlock(log, blk)
	if err != nil {
		return false, errors.Wrap(err, "could not validate block")
	}

	log.Debug("block is valid; processing")

	metricsScrapingDuration.Observe(float64(time.Since(start) / time.Millisecond))

	startProcessing := time.Now()

	log.Debug("updating state cache")
	err = g.state.RefreshCache(ctx)
	if err != nil {
		log.Fatal(err)
	}

	p, err := processor.New(blk, g.state)
	if err != nil {
		return false, errors.Wrap(err, "could not init processor")
	}

	savedBlock, err := p.Store(ctx, g.db)
	if err != nil {
		return false, errors.Wrap(err, "could not store block")
	}

	metricsProcessingDuration.Observe(float64(time.Since(startProcessing) / time.Millisecond))
	log.WithField("duration", time.Since(start)).Info("done processing block")

	return savedBlock, nil
}

func (g *Glue) Run(ctx context.Context) {
	for {
		b, err := g.state.NextTask(ctx)
		if err != nil && err != context.Canceled {
			g.logger.Fatal(err)
		} else if err == context.Canceled {
			return
		}

		acquired, err := g.state.LockBlock(b)
		if err != nil {
			g.logger.Fatal(err)
		}
		if !acquired {
			// could not get lock, already being worked on, skip
			continue
		}

		g.stopMu.Lock()

		savedBlock, err := g.ScrapeSingleBlock(ctx, b)
		if err != nil {
			g.logger.Error(err)
			err = g.state.UnlockBlock(b)
			if err != nil {
				g.logger.Fatal(err)
			}
			g.mustRequeueTask(b)
			metricsBlocksErrored.Inc()
		} else {
			err = g.state.UnlockBlock(b)
			if err != nil {
				g.logger.Fatal(err)
			}
			if savedBlock {
				metricsBlocksProcessed.Inc()
			} else {
				metricsBlocksSkipped.Inc()
			}
		}

		g.stopMu.Unlock()
	}
}

func (g *Glue) mustRequeueTask(b int64) {
	err := g.state.AddTaskToQueue(b)
	if err != nil {
		g.logger.Fatal(err)
	}
}

func (g *Glue) validateBlock(log *logrus.Entry, blk *types.RawData) (bool, error) {
	log.Debug("validating block")

	v := validator.New()
	v.LoadBlock(blk.Block)
	v.LoadReceipts(blk.Receipts)

	return v.Run()
}

func (g *Glue) Close() {
	g.stopMu.Lock()
	defer g.stopMu.Unlock()
}
