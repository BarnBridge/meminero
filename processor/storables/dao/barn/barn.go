package barn

import (
	"context"
	"fmt"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

type Storable struct {
	block *types.Block

	logger *logrus.Entry

	processed struct {
		delegateActions []DelegateAction
		delegateChanges []DelegateChange
		locks           []ethtypes.BarnLockEvent
		stakingActions  []StakingAction
	}
}

func New(block *types.Block) *Storable {
	return &Storable{
		block:  block,
		logger: logrus.WithField("module", "storable(barn)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	var barnLogs []gethtypes.Log
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.Barn.Address) {
				barnLogs = append(barnLogs, log)
			}
		}
	}

	if len(barnLogs) == 0 {
		return nil
	}

	err := s.handleDelegateEvents(barnLogs, ctx)
	if err != nil {
		return err
	}

	err = s.handleLockEvents(barnLogs)
	if err != nil {
		return err
	}

	err = s.handleStakingActions(barnLogs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Trace("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done rolling back block")
	}()

	b := &pgx.Batch{}
	tables := []string{"barn_delegate_actions", "barn_delegate_changes", "barn_locks", "barn_staking_actions"}
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
	start := time.Now()
	s.logger.Trace("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done storing")
	}()

	err := s.storeDelegateChanges(ctx, tx)
	if err != nil {
		return err
	}

	err = s.storeDelegateActions(ctx, tx)
	if err != nil {
		return err
	}

	err = s.storeLockEvents(ctx, tx)
	if err != nil {
		return err
	}

	err = s.storeStakingActionsEvents(ctx, tx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
