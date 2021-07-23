package notifications

/*
import (
	"context"
	"fmt"

	"github.com/barnbridge/smartbackend/utils"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const (
	SmartYieldTokenBought = "smart-yield-token-bought"
)

type SmartYieldJobData struct {
	StartTime             int64           `json:"startTime"`
	PoolAddress           string          `json:"pool"`
	Buyer                 string          `json:"buyer"`
	Amount                decimal.Decimal `json:"amount"`
	IncludedInBlockNumber int64           `json:"includedInBlockNumber"`
}

func NewSmartYieldTokenBoughtJob(data *SmartYieldJobData) (*Job, error) {
	return NewJob(SmartYieldTokenBought, 0, data.IncludedInBlockNumber, data)
}

func (jd *SmartYieldJobData) ExecuteWithTx(ctx context.Context, tx pgx.Tx) ([]*Job, error) {
	log.Tracef("executing token bought from pool %s by %s", jd.PoolAddress, jd.Buyer)

	syPool := state.PoolBySmartYieldAddress(jd.PoolAddress)
	rewardsPool := state.RewardPoolForSYAddress(jd.PoolAddress)
	if rewardsPool == nil {
		return nil, nil
	}

	err := saveNotification(
		ctx, tx,
		jd.Buyer,
		SmartYieldTokenBought,
		jd.StartTime,
		jd.StartTime+60*60*24,
		fmt.Sprintf("Stake your %s junior tokens to earn extra yield", utils.PrettyToken(jd.Amount, syPool.UnderlyingDecimals)),
		smartYieldMetadata(jd, syPool),
		jd.IncludedInBlockNumber,
	)
	if err != nil {
		return nil, errors.Wrap(err, "save smart yield token bought notification to db")
	}

	return nil, nil
}

func smartYieldMetadata(jd *SmartYieldJobData, pool *types.SYPool) map[string]interface{} {
	m := make(map[string]interface{})
	m["amount"] = jd.Amount.String()
	m["underlyingSymbol"] = pool.UnderlyingSymbol
	m["protocolId"] = pool.ProtocolId
	m["syPoolAddress"] = jd.PoolAddress
	return m
}
*/
