package parser

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

type Parser interface {
	Parse(stmt string) (sqlparser.Statement, error)
}

type VSQLParser struct {
}

func (vp *VSQLParser) Parse(stmt string) (sqlparser.Statement, error) {
	out, err := sqlparser.Parse(stmt)
	if err != nil {
		return nil, err
	}
	return out, nil
}
