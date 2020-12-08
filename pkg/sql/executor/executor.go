package executor

import (
	"errors"
	"fmt"
	"log"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

type Executor interface {
	Run(sqlparser.Statement) ([]byte, error)
}

type KitchenExecutor struct {
	store storage.Store
}

func (ke *KitchenExecutor) Run(stmt sqlparser.Statement) ([]byte, error) {
	switch stmt.(type) {
	case *sqlparser.CreateTable:
		{
			// table exist already
			// if not, create an key value pair like this
			// key: table_schema_<table_name>
			// value: columns object (in bytes)
			// ke.store.Get()
			createTableStmt := stmt.(*sqlparser.CreateTable)
			log.Println(createTableStmt.NewName.Name)

			key := []byte(fmt.Sprintf("table_schema_%s", createTableStmt.NewName.Name))
			value, _ := ke.store.Get(key)
			if value != nil {
				return nil, errors.New("Table already exist")
			}

			ke.store.Set(key, []byte(fmt.Sprintf("%v", createTableStmt.Columns)))
			return []byte("1"), nil
		}
	}
	return nil, nil
}
