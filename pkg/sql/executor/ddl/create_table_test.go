package ddl

import (
	"testing"

	"github.com/devenbhooshan/kitchendb/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestKitchenCreateTableExecutorRunForCreateTable(t *testing.T) {
	defer CleanUp("test-storage")
	store, _ := storage.NewPebbleStore("test-storage")
	out, err := CreateTable("create table foobar (bar text);", store)
	assert.Nil(t, err)
	assert.Equal(t, []byte("1"), out)
}
func TestKitchenCreateTableExecutorShouldReturnErrorWhenSameTableCreatedTwice(t *testing.T) {
	defer CleanUp("test-storage-1")
	store, _ := storage.NewPebbleStore("test-storage-1")

	_, err := CreateTable("create table foobar (bar text);", store)
	assert.Nil(t, err)
	_, err = CreateTable("create table foobar (bar text);", store)
	assert.NotNil(t, err)
}
