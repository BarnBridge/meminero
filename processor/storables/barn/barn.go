package barn

import (
	"context"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/sirupsen/logrus"
)

type Storable struct {
	block *types.Block

	logger *logrus.Entry

	processed struct {
	}
}

func New(block *types.Block) *Storable {
	return &Storable{
		block:  block,
		logger: logrus.WithField("module", "storable(barn)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	var barnLogs []gethtypes.Log
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.Barn.Address) {
				barnLogs = append(barnLogs, log)
			}
		}
	}

	if len(barnLogs) == 0 {
		log.Debug("no events found")
		return nil
	}

	return nil
}
