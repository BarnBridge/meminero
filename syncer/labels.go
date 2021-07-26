package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/utils"
)

type Label struct {
	Address string `json:"address"`
	Label   string `json:"label"`
}

type Labels []Label

func (l Labels) Sync(tx pgx.Tx) error {
	if len(l) == 0 {
		return nil
	}

	start := time.Now()
	log.WithField("count", len(l)).Info("syncing labels")
	defer func() {
		log.WithField("duration", time.Since(start)).Info("done syncing labels")
	}()

	for _, a := range l {
		_, err := tx.Exec(context.Background(), `
			insert into labels (address, label) values ($1, $2) on conflict (address) do update set label = $2 
		`, utils.NormalizeAddress(a.Address), a.Label)
		if err != nil {
			return errors.Wrap(err, "could not insert label")
		}
	}

	return nil
}
