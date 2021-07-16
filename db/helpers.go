package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"

	"github.com/barnbridge/smartbackend/config"
)

type ErrorLineExtract struct {
	LineNum   int    // Line number starting with 1
	ColumnNum int    // Column number starting with 1
	Text      string // Text of the line without a new line character.
}

// ExtractErrorLine takes source and character position extracts the line
// number, column number, and the line of text.
//
// The first character is position 1.
func ExtractErrorLine(source string, position int) (ErrorLineExtract, error) {
	ele := ErrorLineExtract{LineNum: 1}
	if position > len(source) {
		return ele, fmt.Errorf("position (%d) is greater than source length (%d)", position, len(source))
	}

	lines := strings.SplitAfter(source, "\n")
	for _, ele.Text = range lines {
		if position-len(ele.Text) < 1 {
			ele.ColumnNum = position
			break
		}
		ele.LineNum += 1
		position -= len(ele.Text)
	}

	ele.Text = strings.TrimSuffix(ele.Text, "\n")

	return ele, nil
}

func (db *DB) pgxPoolConfig() (*pgxpool.Config, error) {
	pgxCfg, err := pgxpool.ParseConfig(config.Store.Database.ConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse db config")
	}

	pgxCfg.ConnConfig.Logger = logrusadapter.NewLogger(logrus.WithField("module", "pgx"))
	pgxCfg.ConnConfig.LogLevel = pgx.LogLevelWarn
	pgxCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &shopspring.Numeric{},
			Name:  "numeric",
			OID:   pgtype.NumericOID,
		})

		return nil
	}

	return pgxCfg, nil
}
