package rewards

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SmartAlpha.RewardPoolByAddress(log.Address.String()) != nil {
				err := s.processRewardPoolEvent(log)
				if err != nil {
					return errors.Wrapf(err, "could not process reward pool event %s")
				}
			}
		}
	}

	return nil
}

func (s *Storable) processRewardPoolEvent(log gethtypes.Log) error {
	poolSingle := *ethtypes.RewardPoolSingle
	poolMulti := *ethtypes.RewardPoolMulti

	if poolSingle.IsClaimEvent(&log) {
		e, err := poolSingle.ClaimEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode PoolSingle.Claim event")
		}

		s.processed.Claims = append(s.processed.Claims, e)

		return nil
	}

	if poolMulti.IsClaimRewardTokenEvent(&log) {
		e, err := poolMulti.ClaimRewardTokenEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode PoolMulti.ClaimRewardToken event")
		}

		s.processed.ClaimsMulti = append(s.processed.ClaimsMulti, e)

		return nil
	}

	if poolSingle.IsDepositEvent(&log) {
		e, err := poolSingle.DepositEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode PoolSingle.Deposit event")
		}

		s.processed.Deposits = append(s.processed.Deposits, e)

		return nil
	}

	if poolSingle.IsWithdrawEvent(&log) {
		e, err := poolSingle.WithdrawEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode PoolSingle.Withdraw event")
		}

		s.processed.Withdrawals = append(s.processed.Withdrawals, e)

		return nil
	}

	return nil
}
