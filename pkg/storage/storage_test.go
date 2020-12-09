package storage

import (
	"log"
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

func TestPebbleStoreGetWithPrefix(t *testing.T) {
	store, _ := NewPebbleStore("test-storage")
	defer store.db.Close()
	defer cleanUp("test-storage")
	keys := []string{"hello", "world", "hello world"}
	values := []string{"YES", "NO", "YES"}
	for i, key := range keys {
		if err := store.Set([]byte(key), []byte(values[i])); err != nil {
			log.Fatal(err)
		}
	}

	output, _ := store.GetWithPrefix([]byte("hello"))

	expectedOutputStr := []string{"YES", "YES"}
	expectedOutput := make([][]byte, len(expectedOutputStr))
	for i, v := range expectedOutputStr {
		expectedOutput[i] = []byte(v)
	}
	
	assert.Equal(t, expectedOutput, output)
}

func cleanUp(dirName string) error {
	return os.RemoveAll(filepath.Join(dirName, ""))
}
