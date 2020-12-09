package ddl

import (
	"errors"
	"fmt"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

type KitchenCreateTableExecutor struct {
	store storage.Store
}

func NewKitchenCreateTableExecutor(store storage.Store) *KitchenCreateTableExecutor {
	return &KitchenCreateTableExecutor{
		store: store,
	}
}

func (ke *KitchenCreateTableExecutor) Run(stmt sqlparser.Statement) ([]byte, error) {
	createTableStmt := stmt.(*sqlparser.CreateTable)

	key := []byte(fmt.Sprintf("table@schema@%s", createTableStmt.NewName.Name))
	value, _ := ke.store.Get(key)
	if value != nil {
		return nil, errors.New("Table already exist")
	}

	ke.store.Set(key, []byte(fmt.Sprintf("%v", createTableStmt.Columns)))
	return []byte("1"), nil
}
