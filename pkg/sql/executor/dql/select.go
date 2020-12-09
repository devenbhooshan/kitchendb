package dql

import (
	"errors"
	"fmt"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

type KitchenSelectExecutor struct {
	store storage.Store
}

func NewKitchenSelectExecutor(store storage.Store) *KitchenSelectExecutor {
	return &KitchenSelectExecutor{
		store: store,
	}
}

func (ke *KitchenSelectExecutor) Run(stmt sqlparser.Statement) ([][]byte, error) {
	selectStmt := stmt.(*sqlparser.Select)
	tbl := selectStmt.From[0].(*sqlparser.AliasedTableExpr)
	tblName := (tbl.Expr.(sqlparser.TableName)).Name

	key := []byte(fmt.Sprintf("table@schema@%s", tblName))
	value, _ := ke.store.Get(key)
	if value != nil {

		keyForRows := []byte(fmt.Sprintf("rows@%s", tblName))
		allRows, err := ke.store.GetWithPrefix(keyForRows)
		if err != nil {
			return nil, err
		}
		return allRows, nil
	}
	return nil, errors.New("Table doesn't exist")

}
