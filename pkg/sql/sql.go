package sql

import (
	"reflect"

	"github.com/devenbhooshan/kitchendb/pkg/sql/executor"

	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
)

type SQLEngine interface {
	Run(stm string) ([]byte, error)
}

type KitchenSQLEngine struct {
	parser   parser.Parser
	executor executor.Executor
}

func (ksqle *KitchenSQLEngine) Run(stm string) ([]byte, error) {
	vsqlparser := parser.VSQLParser{}
	ast, err := vsqlparser.Parse(stm)
	if err != nil {
		return nil, err
	}
	return []byte(reflect.TypeOf(ast).String()), nil
}

func NewKitchenSQLEngine() *KitchenSQLEngine {
	return &KitchenSQLEngine{
		parser:   &parser.VSQLParser{},
		executor: &executor.KitchenExecutor{},
	}
}
