package rewards

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) processNewPoolSingle(ctx context.Context, log ethtypes.SmartYieldPoolFactorySinglePoolCreatedEvent) error {
	poolAddr := utils.NormalizeAddress(log.Pool.String())

	if s.state.SmartYield.RewardPoolByAddress(poolAddr) != nil {
		return nil
	}

	a := *ethtypes.SmartYieldPoolFactorySingle.ABI

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

	s.processed.PoolsSingle = append(s.processed.PoolsSingle, p)

	s.state.SmartYield.CacheRewardPool(p)

	return nil
}

func (s *Storable) processNewPoolMulti(ctx context.Context, log ethtypes.SmartYieldPoolFactoryMultiPoolMultiCreatedEvent) error {
	poolAddr := utils.NormalizeAddress(log.Pool.String())

	if s.state.SmartYield.RewardPoolByAddress(poolAddr) != nil {
		return nil
	}

	a := *ethtypes.SmartYieldPoolFactoryMulti.ABI

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

		err := s.checkTokenExists(addr)
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

	s.processed.PoolsMulti = append(s.processed.PoolsMulti, p)

	s.state.SmartYield.CacheRewardPool(p)

	return nil
}
