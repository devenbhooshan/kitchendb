package parser

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func Parse(stmt string) (sqlparser.Statement, error) {
	out, err := sqlparser.Parse(stmt)
	if err != nil {
		return nil, err
	}
	return out, nil
}
