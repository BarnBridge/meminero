package integrity

import "github.com/pkg/errors"

func (c *Checker) checkMissingBlocks(start, end int64) ([]int64, error) {
	rows, err := c.db.Query(`
		select x.number
		from generate_series($1::bigint, $2::bigint) as x(number)
				 left join (select number from blocks where number between $1 and $2) b on x.number = b.number
		where b.number is null
		order by number;
	`, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "could not query database for missing blocks")
	}

	var blocks []int64
	for rows.Next() {
		var b int64

		err = rows.Scan(&b)
		if err != nil {
			return nil, errors.Wrap(err, "could not scan missing block from db")
		}

		blocks = append(blocks, b)
	}

	return blocks, nil
}
