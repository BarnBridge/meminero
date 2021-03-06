package scraper

import (
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/alethio/web3-go/ethrpc/provider/httprpc"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/types"

	"github.com/alethio/web3-go/ethrpc"
	"github.com/sirupsen/logrus"
)

var (
	metricsTaskDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "scraper_get_task_duration",
		Help: "Duration of different tasks executed by the scraper",
	}, []string{"task"})
)

type Scraper struct {
	conn   *ethrpc.ETH
	logger *logrus.Entry
}

func New() (*Scraper, error) {
	batchLoader, err := httprpc.NewBatchLoader(config.Store.ETH.MaxBatch, 4*time.Millisecond)
	if err != nil {
		return nil, errors.Wrap(err, "could not init batch loader")
	}

	provider, err := httprpc.NewWithLoader(config.Store.ETH.Config.HTTP, batchLoader)
	if err != nil {
		return nil, errors.Wrap(err, "could not init httprpc provider")
	}
	provider.SetHTTPTimeout(5000 * time.Millisecond)

	c, err := ethrpc.New(provider)
	if err != nil {
		return nil, errors.Wrap(err, "could not init ethrpc")
	}

	return &Scraper{
		conn:   c,
		logger: logrus.WithField("module", "scraper"),
	}, nil
}

// Exec does the JSONRPC calls necessary for scraping a given block and returns the raw data
// It:
// - scrapes the block using eth_getBlockByNumber
// - for each transaction in the block, scrapes the receipts using eth_getTransactionReceipt
// - for each uncle in the block, scrapes the data using eth_getUncleByBlockHashAndIndex
func (s *Scraper) Exec(block int64) (*types.RawData, error) {
	var log = s.logger.WithField("block", block)

	b := &types.RawData{}

	log.Debug("getting block")
	start := time.Now()
	dataBlock, err := s.conn.GetBlockByNumber("0x" + strconv.FormatInt(block, 16))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	b.Block = dataBlock
	log.WithField("duration", time.Since(start)).Debug("got block")
	recordDuration("block", start)

	log.Debug("getting receipts")
	start = time.Now()

	var wg sync.WaitGroup
	var errs []error
	var mu sync.Mutex
	for _, tx := range dataBlock.Transactions {
		wg.Add(1)
		txCopy := tx

		go func() {
			defer wg.Done()

			dataReceipt, err := s.conn.GetTransactionReceipt(txCopy.Hash)
			if err != nil {
				errs = append(errs, err)
				return
			}

			mu.Lock()
			b.Receipts = append(b.Receipts, dataReceipt)
			mu.Unlock()
		}()
	}
	wg.Wait()
	sort.Sort(b.Receipts)

	log.WithField("duration", time.Since(start)).Debugf("got %d receipts", len(b.Receipts))
	if len(errs) > 0 {
		return nil, errs[0]
	}
	recordDuration("receipts", start)

	log.Debug("done scraping block")

	return b, nil
}

func recordDuration(task string, start time.Time) {
	d := float64(time.Since(start) / time.Millisecond)
	metricsTaskDuration.WithLabelValues(task).Observe(d)
}
