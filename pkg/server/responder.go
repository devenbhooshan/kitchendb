package server

import (
	"reflect"

	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
)

func Responder(stm string) ([]byte, error) {
	ast, err := parser.Parse(stm)
	if err != nil {
		return nil, err
	}
	return []byte(reflect.TypeOf(ast).String()), nil
}
