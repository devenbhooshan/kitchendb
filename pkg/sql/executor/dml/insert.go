package dml

import (
	"errors"
	"fmt"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
	uuid "github.com/satori/go.uuid"
)

type KitchenInsertExecutor struct {
	store storage.Store
}

func NewKitchenInsertExecutor(store storage.Store) *KitchenInsertExecutor {
	return &KitchenInsertExecutor{
		store: store,
	}
}

func (ke *KitchenInsertExecutor) Run(stmt sqlparser.Statement) ([]byte, error) {
	insertStmt := stmt.(*sqlparser.Insert)
	tblName := insertStmt.Table.Name
	key := []byte(fmt.Sprintf("table@schema@%s", tblName))
	value, _ := ke.store.Get(key)
	if value != nil {
		u2 := uuid.NewV4()
		values := insertStmt.Rows.(sqlparser.Values)[0][0]
		val := (values.(sqlparser.Expr)).(*sqlparser.SQLVal)

		key := []byte(fmt.Sprintf("rows@%s@%s", tblName, u2))

		ke.store.Set(key, []byte(string(val.Val[0])))
		return []byte("1"), nil
	}
	return nil, errors.New("Table doesn't exist")

}
