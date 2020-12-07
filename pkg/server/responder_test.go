package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponderForSelectStatement(t *testing.T) {
	out, _ := Responder("select foo from bar;")
	assert.Equal(t, []byte("*sqlparser.Select"), out)
}

func TestResponderForInsertStatement(t *testing.T) {
	out, _ := Responder("insert into foo values(1,2);")
	assert.Equal(t, []byte("*sqlparser.Insert"), out)
}
