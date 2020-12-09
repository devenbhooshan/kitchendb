package executor

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/devenbhooshan/kitchendb/pkg/sql/executor/ddl"
	"github.com/devenbhooshan/kitchendb/pkg/sql/executor/dml"
	"github.com/devenbhooshan/kitchendb/pkg/sql/executor/dql"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

type Executor interface {
	Run(sqlparser.Statement) ([][]byte, error)
}

type KitchenExecutor struct {
	store storage.Store
}

func (ke *KitchenExecutor) Run(stmt sqlparser.Statement) ([][]byte, error) {
	switch stmt.(type) {
	case *sqlparser.CreateTable:
		{
			kcte := ddl.NewKitchenCreateTableExecutor(ke.store)
			output, err := kcte.Run(stmt)
			return [][]byte{output}, err
		}
	case *sqlparser.Select:
		{
			kse := dql.NewKitchenSelectExecutor(ke.store)
			return kse.Run(stmt)
		}

	case *sqlparser.Insert:
		{
			kse := dml.NewKitchenInsertExecutor(ke.store)
			output, err := kse.Run(stmt)
			return [][]byte{output}, err
		}
	}
	return nil, nil
}

func NewKitchenExecutor(store storage.Store) *KitchenExecutor {
	return &KitchenExecutor{
		store: store,
	}
}
