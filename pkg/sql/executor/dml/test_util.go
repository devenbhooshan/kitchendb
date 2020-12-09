package dml

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/devenbhooshan/kitchendb/pkg/sql/executor/ddl"
	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

func CreateTestData(store storage.Store) {
	parser := parser.VSQLParser{}

	ddl.CreateTable("create table foobar (bar text);", store)
	kie := NewKitchenInsertExecutor(store)

	for i := 0; i < 3; i++ {
		stmt, _ := parser.Parse(fmt.Sprintf("insert into foobar(bar) values(%d)", i))
		kie.Run(stmt)
	}
}

func CleanUp(dirName string) error {
	return os.RemoveAll(filepath.Join(dirName, ""))
}
