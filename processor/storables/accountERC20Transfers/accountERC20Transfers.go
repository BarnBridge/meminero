package accountERC20Transfers

import (
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/state"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Storable struct {
	block   *types.Block
	ethConn *ethclient.Client
	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		transfers      []ethtypes.ERC20TransferEvent
		blockNumber    int64
		blockTimestamp int64
	}
}

func New(block *types.Block ,ethConn *ethclient.Client,state  *state.Manager) *Storable {
	return &Storable{
		block:   block,
		ethConn: ethConn,
		state: state,
		logger: logrus.WithField("module", "storable(account_erc20_transfers)"),
	}

}

func (s *Storable) Execute() error {
	var logs [] gethtypes.Log
	erc20Decoder := ethtypes.NewERC20Decoder()
	for _, tx := range s.block.Txs{
		for _, log := range tx.LogEntries {
			if erc20Decoder.IsERC20TransferEvent(&log) && s.state.IsMonitoredAccount(log) {
				err := s.checkTokenExists(utils.NormalizeAddress(log.Address.String()))
				if err != nil {
					continue
				}

				logs = append(logs, log)
			}
		}
	}
	err := s.decodeLogs(logs,erc20Decoder)
	if err != nil {
		return errors.Wrap(err,"could not decode erc20 transfers logs")
	}
	return nil
}

func (s *Storable) Rollback(pgx  pgx.Tx) error {
	return nil
}

func (s *Storable) SaveToDatabase(tx  pgx.Tx) error {
	err := s.storeTransfers(tx)
	if err != nil {
		return errors.Wrap(err, "could not store erc20transfers")
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}