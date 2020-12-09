package dql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devenbhooshan/kitchendb/pkg/sql/executor/dml"
	"github.com/devenbhooshan/kitchendb/pkg/sql/parser"
	"github.com/devenbhooshan/kitchendb/pkg/storage"
)

func TestKitchenSelectExecutor(t *testing.T) {
	parser := parser.VSQLParser{}
	store, _ := storage.NewPebbleStore("test-storage-1")
	dml.CreateTestData(store)

	stmt, _ := parser.Parse("select * from foobar;")

	defer dml.CleanUp("test-storage-1")

	ke := KitchenSelectExecutor{
		store: store,
	}
	out, err := ke.Run(stmt)
	assert.Nil(t, err)
	expectedOutputStr := []string{"0", "1", "2"}
	outputStr := make([]string, len(out))
	for i, v := range out {
		outputStr[i] = string(v)
	}
	assert.Equal(t, true, testEq(expectedOutputStr, outputStr))

}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}
	foundCount := 0

	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				foundCount++
				break
			}
		}
	}
	return foundCount == len(a)
}
