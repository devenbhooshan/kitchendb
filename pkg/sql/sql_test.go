package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunForSelectStatement(t *testing.T) {
	sqlEngine := NewKitchenSQLEngine()
	out, _ := sqlEngine.Run("select foo from bar;")
	assert.Equal(t, []byte("*sqlparser.Select"), out)
}

func TestRunForInsertStatement(t *testing.T) {
	sqlEngine := NewKitchenSQLEngine()
	out, _ := sqlEngine.Run("insert into foo values(1,2);")
	assert.Equal(t, []byte("*sqlparser.Insert"), out)
}
