package barn

import (
	"context"
	"fmt"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Storable struct {
	block *types.Block

	logger *logrus.Entry

	processed struct {
		delegateActions []DelegateAction
		delegateChanges []DelegateChange
		locks           []ethtypes.BarnLockEvent
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
		s.logger.WithField("handler", "barn").Debug("no events found")
		return nil
	}

	err := s.handleDelegateEvents(barnLogs, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	b := &pgx.Batch{}
	tables := [7]string{"barn_delegate_actions", "barn_delegate_changes", "barn_locks", "barn_staking_actions"}
	for _, t := range tables {
		query := fmt.Sprintf(`delete from governance.%s where included_in_block = $1`, t)
		b.Queue(query, s.block.Number)
	}

	br := tx.SendBatch(ctx, b)
	_, err := br.Exec()
	if err != nil {
		return err
	}

	err = br.Close()
	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}