package rewards

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.isFactory(log.Address.String()) {
				if ethtypes.SmartYieldPoolFactorySingle.IsPoolCreatedEvent(&log) {
					p, err := ethtypes.SmartYieldPoolFactorySingle.PoolCreatedEvent(log)
					if err != nil {
						return errors.Wrap(err, "could not decode PoolSingle.PoolCreated event")
					}

					err = s.processNewPoolSingle(ctx, p)
					if err != nil {
						return errors.Wrap(err, "could not process new PoolSingle")
					}
				}

				if ethtypes.SmartYieldPoolFactoryMulti.IsPoolMultiCreatedEvent(&log) {
					p, err := ethtypes.SmartYieldPoolFactoryMulti.PoolMultiCreatedEvent(log)
					if err != nil {
						return errors.Wrap(err, "could not decode PoolMulti.PoolMultiCreatedEvent")
					}

					err = s.processNewPoolMulti(ctx, p)
					if err != nil {
						return errors.Wrap(err, "could not process new PoolMulti")
					}
				}
			}

			if s.state.SmartYield.RewardPoolByAddress(log.Address.String()) != nil {
				err := s.processRewardPoolEvent(log)
				if err != nil {
					return errors.Wrapf(err, "could not process reward pool event %s")
				}
			}
		}
	}

	return nil
}

func (s *Storable) processNewPoolSingle(ctx context.Context, log ethtypes.SmartYieldPoolFactorySinglePoolCreatedEvent) error {
	poolAddr := utils.NormalizeAddress(log.Pool.String())

	if s.state.SmartYield.RewardPoolByAddress(poolAddr) != nil {
		return nil
	}

	a := *ethtypes.SmartYieldPoolSingle.ABI

	var rewardToken, poolToken common.Address

	wg, _ := errgroup.WithContext(ctx)
	wg.Go(eth.CallContractFunction(a, poolAddr, "rewardToken", []interface{}{}, &rewardToken))
	wg.Go(eth.CallContractFunction(a, poolAddr, "poolToken", []interface{}{}, &poolToken))

	err := wg.Wait()
	if err != nil {
		return errors.Wrap(err, "could not call contract functions")
	}

	var p smartyield.RewardPool
	p.PoolType = smartyield.PoolTypeSingle
	p.PoolAddress = poolAddr
	p.RewardTokenAddresses = []string{utils.NormalizeAddress(rewardToken.String())}
	p.PoolTokenAddress = utils.NormalizeAddress(poolToken.String())
	p.StartAtBlock = s.block.Number

	s.processed.Pools = append(s.processed.Pools, p)

	s.state.SmartYield.CacheRewardPool(p)

	return nil
}

func (s *Storable) processNewPoolMulti(ctx context.Context, log ethtypes.SmartYieldPoolFactoryMultiPoolMultiCreatedEvent) error {
	poolAddr := utils.NormalizeAddress(log.Pool.String())

	if s.state.SmartYield.RewardPoolByAddress(poolAddr) != nil {
		return nil
	}

	a := *ethtypes.SmartYieldPoolMulti.ABI

	var numRewardTokens *big.Int

	err := eth.CallContractFunction(a, poolAddr, "numRewardTokens", []interface{}{}, &numRewardTokens)()
	if err != nil {
		return errors.Wrap(err, "could not get numRewardTokens")
	}

	var rewardTokens = make([]common.Address, numRewardTokens.Int64())
	var poolToken common.Address

	wg, _ := errgroup.WithContext(ctx)

	for i := int64(0); i < numRewardTokens.Int64(); i++ {
		wg.Go(eth.CallContractFunction(a, poolAddr, "rewardTokens", []interface{}{big.NewInt(i)}, &rewardTokens[i]))
	}
	wg.Go(eth.CallContractFunction(a, poolAddr, "poolToken", []interface{}{}, &poolToken))

	err = wg.Wait()
	if err != nil {
		return errors.Wrap(err, "could not get pool info")
	}

	var rewardTokensString []string
	for _, rt := range rewardTokens {
		addr := utils.NormalizeAddress(rt.String())
		rewardTokensString = append(rewardTokensString, addr)

		err := s.checkTokenExists(ctx, addr)
		if err != nil {
			return errors.Wrap(err, "could not ensure reward token exists")
		}
	}

	var p smartyield.RewardPool
	p.PoolType = smartyield.PoolTypeMulti
	p.PoolAddress = poolAddr
	p.RewardTokenAddresses = rewardTokensString
	p.PoolTokenAddress = utils.NormalizeAddress(poolToken.String())
	p.StartAtBlock = s.block.Number

	s.processed.Pools = append(s.processed.Pools, p)

	s.state.SmartYield.CacheRewardPool(p)

	return nil
}

func (s *Storable) processRewardPoolEvent(log gethtypes.Log) error {
	poolSingle := *ethtypes.SmartYieldPoolSingle
	poolMulti := *ethtypes.SmartYieldPoolMulti

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
