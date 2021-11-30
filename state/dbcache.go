package state

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

const JuniorAPYAvgRefreshInterval = 30 * time.Minute

func (m *Manager) refreshDBCache(ctx context.Context) error {
	err := m.refreshJuniorAvgAPY(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) refreshJuniorAvgAPY(ctx context.Context) error {
	var lastUpdated time.Time
	err := m.db.QueryRow(ctx, `select last_updated from smart_yield.junior_apy_30d_avg limit 1`).Scan(&lastUpdated)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrap(err, "could not scan junior_apy_30d_avg last_updated")
	}

	if err == pgx.ErrNoRows || lastUpdated.Before(time.Now().Add(-JuniorAPYAvgRefreshInterval)) {
		_, err := m.db.Exec(ctx, `refresh materialized view concurrently smart_yield.junior_apy_30d_avg`)
		if err != nil {
			return errors.Wrap(err, "could not refresh smart_yield.junior_apy_30d_avg")
		}
	} else {
		m.logger.Trace("no need to refresh smart_yield.junior_apy_30d_avg")
	}

	return nil
}
