package sql

import (
	"testing"

	"github.com/devenbhooshan/kitchendb/pkg/sql/executor/ddl"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestRunForCreateTable(t *testing.T) {
	defer ddl.CleanUp("test-storage")
	store, _ := storage.NewPebbleStore("test-storage")

	sqlEngine := NewKitchenSQLEngine(store)
	out, err := sqlEngine.Run("create table foobar (bar text);")
	assert.NotNil(t, out)
	assert.Nil(t, err)
}
