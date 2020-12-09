package sql

import (
	"github.com/devenbhooshan/kitchendb/pkg/sql/executor"

	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

type SQLEngine interface {
	Run(stm string) ([][]byte, error)
}

type KitchenSQLEngine struct {
	parser   parser.Parser
	executor executor.Executor
}

func (ksqle *KitchenSQLEngine) Run(stm string) ([][]byte, error) {
	ast, err := ksqle.parser.Parse(stm)
	if err != nil {
		return nil, err
	}
	return ksqle.executor.Run(ast)
}

func NewKitchenSQLEngine(store storage.Store) *KitchenSQLEngine {
	return &KitchenSQLEngine{
		parser:   &parser.VSQLParser{},
		executor: executor.NewKitchenExecutor(store),
	}
}
