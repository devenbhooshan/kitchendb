package ddl

import (
	"os"
	"path/filepath"

	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

func CreateTable(createTableStmt string, store storage.Store) ([]byte, error) {
	parser := parser.VSQLParser{}
	stmt, _ := parser.Parse(createTableStmt)

	//"create table foobar (bar text)
	ke := NewKitchenCreateTableExecutor(store)
	return ke.Run(stmt)
}

func CleanUp(dirName string) error {
	return os.RemoveAll(filepath.Join(dirName, ""))
}
