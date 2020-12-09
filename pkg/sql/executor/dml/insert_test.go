package dml

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devenbhooshan/kitchendb/pkg/sql/executor/ddl"
	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

func TestKitchenInsertExecutorRun(t *testing.T) {
	parser := parser.VSQLParser{}

	store, _ := storage.NewPebbleStore("test-storage-1")
	ddl.CreateTable("create table foobar (bar text);", store)

	stmt, _ := parser.Parse("insert into foobar(bar) values(1)")
	defer ddl.CleanUp("test-storage-1")

	kie := NewKitchenInsertExecutor(store)
	out, err := kie.Run(stmt)
	assert.Nil(t, err)
	assert.Equal(t, []byte("1"), out)
}
