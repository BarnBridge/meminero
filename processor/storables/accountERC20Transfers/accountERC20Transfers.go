package accountERC20Transfers

import (
	"database/sql"

	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/state"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Storable struct {
	block   *types.Block
	ethConn *ethclient.Client

	processed struct {
		transfers      []ethtypes.ERC20TransferEvent
		blockNumber    int64
		blockTimestamp int64
	}
}

func New(block *types.Block ,ethConn *ethclient.Client) *Storable {
	return &Storable{
		block:   block,
		ethConn: ethConn,
	}

}

func (s *Storable) Execute() error {
	var logs [] gethtypes.Log
	erc20Decoder := ethtypes.NewERC20Decoder()

	for _, tx := range s.block.Txs{
		for _, log := range tx.LogEntries {
			if erc20Decoder.IsERC20TransferEvent(&log) && state.IsMonitoredAccount(log) {
				err := s.checkTokenExists(tx, utils.NormalizeAddress(log.Address.String()))
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

func (s *Storable) Rollback(tx *sql.Tx) error {
	return nil
}

func (s *Storable) SaveToDatabase(tx *sql.Tx) error {
	err := s.storeTransfers(tx)
	if err != nil {
		return errors.Wrap(err, "could not store erc20transfers")
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}