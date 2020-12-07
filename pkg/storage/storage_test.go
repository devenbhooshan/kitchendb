package storage

import (
	"testing"

	"os"
	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func TestPebbleStoreSet(t *testing.T) {
	store, err := NewPebbleStore("test-storage")
	defer store.db.Close()
	defer cleanUp("test-storage")
	key := []byte("test-key")
	expectedValue := []byte("test-value")

	assert.Nil(t, err)

	assert.Nil(t, store.Set(key, expectedValue))

	value, err := store.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, expectedValue, value)
}

func cleanUp(dirName string) error {
	return os.RemoveAll(filepath.Join(dirName, ""))
}
