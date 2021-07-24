package accounterc20transfers

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

var (
	metricsTransfersProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "scraper_erc20_transfers_total",
		Help: "Number of ERC20 transfers detected",
	})
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		transfers []ethtypes.ERC20TransferEvent
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(account_erc20_transfers)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if len(log.Topics) == 3 && ethtypes.ERC20.IsERC20TransferEvent(&log) {
				erc20Transfer, err := ethtypes.ERC20.ERC20TransferEvent(log)
				if err != nil {
					return errors.Wrapf(err, "could not decode erc20 transfer in tx %s", log.TxHash.String())
				}

				if !s.state.IsMonitoredAccount(erc20Transfer.From.String()) &&
					!s.state.IsMonitoredAccount(erc20Transfer.To.String()) {
					continue
				}

				s.processed.transfers = append(s.processed.transfers, erc20Transfer)

				exists := s.state.CheckTokenExists(log.Address.String())
				if !exists {
					token, err := eth.GetERC20TokenFromChain(log.Address.String())
					if err != nil {
						return err
					}

					err = s.state.StoreToken(ctx, *token)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Debug("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from account_erc20_transfers where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.storeTransfers(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store erc20transfers")
	}

	metricsTransfersProcessed.Add(float64(len(s.processed.transfers)))
	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
