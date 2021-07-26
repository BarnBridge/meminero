package notifications

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

const (
	SmartYieldTokenBought = "smart-yield-token-bought"
)

type SmartYieldJobData struct {
	StartTime             int64           `json:"startTime"`
	Pool                  smartyield.Pool `json:"pool"`
	HasRewardPool         bool            `json:"hasRewardPool"`
	Buyer                 string          `json:"buyer"`
	Amount                decimal.Decimal `json:"amount"`
	IncludedInBlockNumber int64           `json:"includedInBlockNumber"`
}

func NewSmartYieldTokenBoughtJob(data *SmartYieldJobData) (*Job, error) {
	return NewJob(SmartYieldTokenBought, 0, data.IncludedInBlockNumber, data)
}

func (jd *SmartYieldJobData) ExecuteWithTx(ctx context.Context, tx pgx.Tx) ([]*Job, error) {
	log.Tracef("executing token bought from pool %s by %s", jd.Pool.PoolAddress, jd.Buyer)

	if !jd.HasRewardPool {
		return nil, nil
	}

	err := saveNotification(
		ctx, tx,
		jd.Buyer,
		SmartYieldTokenBought,
		jd.StartTime,
		jd.StartTime+60*60*24,
		fmt.Sprintf("Stake your %s junior tokens to earn extra yield", utils.PrettyToken(jd.Amount, jd.Pool.UnderlyingDecimals)),
		smartYieldMetadata(jd, jd.Pool),
		jd.IncludedInBlockNumber,
	)
	if err != nil {
		return nil, errors.Wrap(err, "save smart yield token bought notification to db")
	}

	return nil, nil
}

func smartYieldMetadata(jd *SmartYieldJobData, pool smartyield.Pool) map[string]interface{} {
	m := make(map[string]interface{})
	m["amount"] = jd.Amount.String()
	m["underlyingSymbol"] = pool.UnderlyingSymbol
	m["protocolId"] = pool.ProtocolId
	m["syPoolAddress"] = jd.Pool.PoolAddress
	return m
}
