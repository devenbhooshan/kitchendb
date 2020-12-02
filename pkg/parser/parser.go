package parser;


import (

 "github.com/blastrain/vitess-sqlparser/sqlparser"
)

func Parse(stmt string) sqlparser.Statement {
 out, err := sqlparser.Parse(stmt)
 if err != nil {
	 panic(err)
 }
 return out
}