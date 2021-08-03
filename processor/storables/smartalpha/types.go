package smartalpha

import (
	"math/big"

	"github.com/barnbridge/meminero/types"
)

type Pool struct {
	PoolName           string `json:"poolName"`
	PoolAddress        string `json:"poolAddress"`
	PoolToken          types.Token
	JuniorTokenAddress string `json:"juniorTokenAddress"`
	SeniorTokenAddress string `json:"seniorTokenAddress"`
	OracleAddress      string `json:"oracleAddress"`
	OracleAssetSymbol  string `json:"oracleAssetSymbol"`
	Epoch1Start        int64  `json:"epoch1Start"`
	EpochDuration      int64  `json:"epochDuration"`
	StartAtBlock       int64  `json:"startAtBlock"`
}

const JuniorTranche = "JUNIOR"
const SeniorTranche = "SENIOR"

type TxType string

const (
	JuniorEntry            TxType = "JUNIOR_ENTRY"
	JuniorRedeemTokens     TxType = "JUNIOR_REDEEM_TOKENS"
	JuniorExit             TxType = "JUNIOR_EXIT"
	JuniorRedeemUnderlying TxType = "JUNIOR_REDEEM_UNDERLYING"
	SeniorEntry            TxType = "SENIOR_ENTRY"
	SeniorRedeemTokens     TxType = "SENIOR_REDEEM_TOKENS"
	SeniorExit             TxType = "SENIOR_EXIT"
	SeniorRedeemUnderlying TxType = "SENIOR_REDEEM_UNDERLYING"
	JtokenSend             TxType = "JTOKEN_SEND"
	JtokenReceive          TxType = "JTOKEN_RECEIVE"
	StokenSend             TxType = "STOKEN_SEND"
	StokenReceive          TxType = "STOKEN_RECEIVE"
)

type State struct {
	PoolAddress                string
	QueuedJuniorsUnderlyingIn  *big.Int
	QueuedJuniorsUnderlyingOut *big.Int
	QueuedJuniorTokensBurn     *big.Int
	QueuedSeniorsUnderlyingIn  *big.Int
	QueuedSeniorsUnderlyingOut *big.Int
	QueuedSeniorTokensBurn     *big.Int
	EstimatedJuniorLiquidity   *big.Int
	EstimatedSeniorLiquidity   *big.Int
	EstimatedJuniorTokenPrice  *big.Int
	EstimatedSeniorTokenPrice  *big.Int
}

type EpochInfo struct {
	PoolAddress            string
	Epoch                  *big.Int
	SeniorLiquidity        *big.Int
	JuniorLiquidity        *big.Int
	UpsideExposureRate     *big.Int
	DownsideProtectionRate *big.Int
	JuniorTokenPrice       *big.Int
	SeniorTokenPrice       *big.Int
}
