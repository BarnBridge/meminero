package smartalpha

import (
	"context"

	"github.com/pkg/errors"
)

func (sa *SmartAlpha) HasEpochInfo(ctx context.Context, poolAddress string, epochId int64) (bool, error) {
	var count int
	err := sa.db.QueryRow(
		ctx,
		`select count(*) from smart_alpha.pool_epoch_info where pool_address = $1 and epoch_id = $2`,
		poolAddress, epochId,
	).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "could not check database if epoch info exists")
	}

	return count > 0, nil
}
