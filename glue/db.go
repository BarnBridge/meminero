package glue

import "database/sql"

// getHighestBlock returns the highest block inserted into the database
func (g *Glue) getHighestBlock() (int64, error) {
	var block int64

	err := g.db.QueryRow("select number from blocks order by number desc limit 1").Scan(&block)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return block, nil
}
