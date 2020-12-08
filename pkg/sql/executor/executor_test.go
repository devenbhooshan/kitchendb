package executor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

func TestKitchenExecutorRunForCreateTable(t *testing.T) {
	parser := parser.VSQLParser{}
	stmt, _ := parser.Parse("create table foobar (bar text);")
	store, _ := storage.NewPebbleStore("test-storage")

	defer cleanUp("test-storage")

	ke := KitchenExecutor{
		store: store,
	}
	out, err := ke.Run(stmt)
	assert.Nil(t, err)
	assert.Equal(t, []byte("1"), out)
}

func cleanUp(dirName string) error {
	return os.RemoveAll(filepath.Join(dirName, ""))
}
