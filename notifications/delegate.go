package notifications

import (
	"context"
	"fmt"

	"github.com/barnbridge/smartbackend/utils"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const (
	DelegateStart = "delegate-start"
)

type DelegateJobData struct {
	StartTime             int64           `json:"startTime"`
	From                  string          `json:"from"`
	To                    string          `json:"to"`
	Amount                decimal.Decimal `json:"amount"`
	IncludedInBlockNumber int64           `json:"includedInBlockNumber"`
}

func NewDelegateStartJob(data *DelegateJobData) (*Job, error) {
	return NewJob(DelegateStart, 0, data.IncludedInBlockNumber, data)
}

func (jd *DelegateJobData) ExecuteWithTx(ctx context.Context, tx pgx.Tx) ([]*Job, error) {
	log.Tracef("executing delegate start form %s to %s", jd.From, jd.To)

	err := saveNotification(
		ctx, tx,
		jd.To,
		DelegateStart,
		jd.StartTime,
		jd.StartTime+60*60*24,
		fmt.Sprintf("%s vBOND has been delegated to you from %s", utils.PrettyBond(jd.Amount), jd.From),
		delegateMetadata(jd),
		jd.IncludedInBlockNumber,
	)
	if err != nil {
		return nil, errors.Wrap(err, "save delegated notification to db")
	}

	return nil, nil
}

func delegateMetadata(jd *DelegateJobData) map[string]interface{} {
	m := make(map[string]interface{})
	m["from"] = jd.From
	m["to"] = jd.To
	m["amount"] = jd.Amount.String()
	return m
}
