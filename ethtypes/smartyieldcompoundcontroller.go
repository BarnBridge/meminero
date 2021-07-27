// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethtypes

import (
	"math/big"

	web3types "github.com/alethio/web3-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lacasian/ethwheels/ethgen"
	"github.com/shopspring/decimal"
)

// Reference imports to suppress errors
var (
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = web3types.Log{}
	_ = decimal.NewFromBigInt
)

const SmartYieldCompoundControllerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"smartYield_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bondModel_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"uniswapPath_\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compRewardTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compRewardSold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingPoolShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"harvestCost\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCKS_PER_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOND_LIFE_MAX\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOND_MAX_RATE_PER_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOUBLE_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXP_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_BUY_JUNIOR_TOKEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_REDEEM_SENIOR_BOND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HARVEST_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSED_BUY_JUNIOR_TOKEN\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSED_BUY_SENIOR_BOND\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_ROUTER_V2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevCTokenBalance_\",\"type\":\"uint256\"}],\"name\":\"_afterCTokenBalanceChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_beforeCTokenBalanceChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cTokens_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate_\",\"type\":\"uint256\"}],\"name\":\"cTokensToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeDistributionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeSupplyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulatives\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cTokenBalance_\",\"type\":\"uint256\"},{\"internalType\":\"uint224\",\"name\":\"nowSupplyStateIndex_\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"prevSupplyStateIndex_\",\"type\":\"uint224\"}],\"name\":\"expectedDistributeSupplierComp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxCompAmount_\",\"type\":\"uint256\"}],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"compGot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingHarvestReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestedLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevCompSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevCumulationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevExchnageRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerRatePerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"compIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeElapsed_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"prevUniswapPriceCumulatives_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nowUniswapPriceCumulatives_\",\"type\":\"uint256[]\"}],\"name\":\"quoteCompToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"compIn_\",\"type\":\"uint256\"}],\"name\":\"quoteSpotCompToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newVal_\",\"type\":\"uint16\"}],\"name\":\"setBondLifeMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVal_\",\"type\":\"uint256\"}],\"name\":\"setBondMaxRatePerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVal_\",\"type\":\"address\"}],\"name\":\"setBondModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dao_\",\"type\":\"address\"}],\"name\":\"setDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVal_\",\"type\":\"uint256\"}],\"name\":\"setFeeBuyJuniorToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVal_\",\"type\":\"uint256\"}],\"name\":\"setFeeRedeemSeniorBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVal_\",\"type\":\"address\"}],\"name\":\"setFeesOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian_\",\"type\":\"address\"}],\"name\":\"setGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setHarvestCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVal_\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"buyJToken_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"buySBond_\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newUniswapPath_\",\"type\":\"address[]\"}],\"name\":\"setUniswapPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smartYield\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spotDailyDistributionRateProvider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spotDailyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spotDailySupplyRateProvider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prevPriceCumulative_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nowPriceCumulative_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeElapsed_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"}],\"name\":\"uniswapAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uniswapPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uniswapPath\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"priceKey_\",\"type\":\"uint8\"}],\"name\":\"uniswapPriceCumulativeNow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uniswapPriceCumulatives\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapPriceCumulativesNow\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uniswapPriceKeys\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController_\",\"type\":\"address\"}],\"name\":\"yieldControllTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var SmartYieldCompoundController = NewSmartYieldCompoundControllerDecoder()

type SmartYieldCompoundControllerDecoder struct {
	*ethgen.Decoder
}

func NewSmartYieldCompoundControllerDecoder() *SmartYieldCompoundControllerDecoder {
	dec := ethgen.NewDecoder(SmartYieldCompoundControllerABI)
	return &SmartYieldCompoundControllerDecoder{
		dec,
	}
}

type SmartYieldCompoundControllerHarvestEvent struct {
	Caller              common.Address
	CompRewardTotal     *big.Int
	CompRewardSold      *big.Int
	UnderlyingPoolShare *big.Int
	UnderlyingReward    *big.Int
	HarvestCost         *big.Int
	Raw                 types.Log
}

func (e *SmartYieldCompoundControllerHarvestEvent) CompRewardTotalDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CompRewardTotal, exp)
}

func (e *SmartYieldCompoundControllerHarvestEvent) CompRewardSoldDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.CompRewardSold, exp)
}

func (e *SmartYieldCompoundControllerHarvestEvent) UnderlyingPoolShareDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.UnderlyingPoolShare, exp)
}

func (e *SmartYieldCompoundControllerHarvestEvent) UnderlyingRewardDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.UnderlyingReward, exp)
}

func (e *SmartYieldCompoundControllerHarvestEvent) HarvestCostDecimal(exp int32) decimal.Decimal {
	return decimal.NewFromBigInt(e.HarvestCost, exp)
}

func (d *SmartYieldCompoundControllerDecoder) HarvestEventID() common.Hash {
	return common.HexToHash("0xeaedd1267621f4a8ee97011faf0e67800ad3063714bd179f078c980bb9a3fa3d")
}

func (d *SmartYieldCompoundControllerDecoder) IsHarvestEvent(log *types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.HarvestEventID()
}

func (d *SmartYieldCompoundControllerDecoder) IsHarvestEventW3(log *web3types.Log) bool {
	if len(log.Topics) == 0 {
		return false
	}
	return log.Topics[0] == d.HarvestEventID().String()
}

func (d *SmartYieldCompoundControllerDecoder) HarvestEventW3(w3l web3types.Log) (SmartYieldCompoundControllerHarvestEvent, error) {
	l, err := ethgen.W3LogToLog(w3l)
	if err != nil {
		return SmartYieldCompoundControllerHarvestEvent{}, err
	}

	return d.HarvestEvent(l)
}

func (d *SmartYieldCompoundControllerDecoder) HarvestEvent(l types.Log) (SmartYieldCompoundControllerHarvestEvent, error) {
	var out SmartYieldCompoundControllerHarvestEvent
	if !d.IsHarvestEvent(&l) {
		return out, ethgen.ErrMismatchingEvent
	}
	err := d.UnpackLog(&out, "Harvest", l)
	out.Raw = l
	return out, err
}
