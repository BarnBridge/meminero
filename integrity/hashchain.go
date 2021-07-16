package integrity

import (
	"context"

	"github.com/pkg/errors"
)

func (c *Checker) checkBrokenHashChain(ctx context.Context,start, end int64) ([]int64, error) {
	rows, err := c.db.Query(ctx, `
		with a as (
			select number
			from blocks as t1
			where t1.number between $1 and $2
			  and (select block_hash from blocks as t2 where t2.number = t1.number - 1) != t1.parent_block_hash
		)
		select number
		from a
		union all
		select number - 1
		from a
		order by number;
	`, start-100, end)
	if err != nil {
		return nil, errors.Wrap(err, "could not query database for broken hash chain")
	}

	var blocks []int64
	for rows.Next() {
		var b int64

		err = rows.Scan(&b)
		if err != nil {
			return nil, errors.Wrap(err, "could not scan inconsistent block from db")
		}

		blocks = append(blocks, b)
	}

	return blocks, nil
}
