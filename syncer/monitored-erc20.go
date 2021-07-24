package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/utils"
)

type MonitoredERC20 []string

func (m MonitoredERC20) Sync(tx pgx.Tx) error {
	if len(m) == 0 {
		return nil
	}

	start := time.Now()
	log.WithField("count", len(m)).Info("syncing monitored erc20")
	defer func() {
		log.WithField("duration", time.Since(start)).Info("done syncing monitored erc20")
	}()

	for _, acc := range m {
		_, err := tx.Exec(context.Background(), `
			insert into monitored_erc20 (address)
			values ($1) 
			on conflict do nothing 
		`, utils.NormalizeAddress(acc))
		if err != nil {
			return errors.Wrap(err, "could not insert monitored erc20")
		}
	}

	return nil
}
