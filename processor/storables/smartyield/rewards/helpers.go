package rewards

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) isFactory(addr string) bool {
	addr = utils.NormalizeAddress(addr)
	for _, v := range s.factories {
		if addr == utils.NormalizeAddress(v) {
			return true
		}
	}

	return false
}

func (s *Storable) checkTokenExists(tokenAddress string) error {
	if s.state.CheckTokenExists(tokenAddress) {
		return nil
	}

	token, err := eth.GetERC20TokenFromChain(tokenAddress)
	if err != nil {
		return errors.Wrap(err, "could not get erc20 info from chain")
	}

	err = s.state.StoreToken(context.Background(), *token)
	if err != nil {
		return errors.Wrap(err, "could not store token to state")
	}

	return nil
}

func (s *Storable) txHistory(user string, amount decimal.Decimal, tranche string, txType smartyield.TxType, raw gethtypes.Log) []interface{} {
	rp := s.state.SmartYield.RewardPoolByAddress(raw.Address.String())
	p := s.state.SmartYield.PoolByAddress(rp.PoolTokenAddress)

	return []interface{}{
		p.ProtocolId,
		p.PoolAddress,
		p.UnderlyingAddress,
		utils.NormalizeAddress(user),
		amount,
		tranche,
		txType,
		s.block.BlockCreationTime,
		s.block.Number,
		utils.NormalizeAddress(raw.TxHash.String()),
		raw.TxIndex,
		raw.Index,
	}
}
