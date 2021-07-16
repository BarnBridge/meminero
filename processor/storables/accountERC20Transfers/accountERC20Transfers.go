package accountERC20Transfers

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/eth"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/state"
	"github.com/barnbridge/smartbackend/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry
	ctx    context.Context

	processed struct {
		transfers []ethtypes.ERC20TransferEvent

	}
}

func New(block *types.Block, state *state.Manager, ctx context.Context) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		ctx: ctx,
		logger: logrus.WithField("module", "storable(account_erc20_transfers)"),
	}
}

func (s *Storable) Execute() error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if ethtypes.ERC20.IsERC20TransferEvent(&log) {
				erc20Transfer, err := ethtypes.ERC20.ERC20TransferEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode erc20 transfer")
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

					err = s.state.StoreToken(s.ctx,*token)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (s *Storable) Rollback(tx pgx.Tx) error {
	_, err := tx.Exec(s.ctx, `delete from account_erc20_transfers where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(tx pgx.Tx) error {
	err := s.storeTransfers(tx)
	if err != nil {
		return errors.Wrap(err, "could not store erc20transfers")
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
