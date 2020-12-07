package parser

import (
	"testing"

	"github.com/blastrain/vitess-sqlparser/sqlparser"

	"github.com/stretchr/testify/assert"
)

func TestParserShouldReturnNonNilValue(t *testing.T) {
	parser := VSQLParser{}
	out, _ := parser.Parse("select now()")
	assert.NotNil(t, out)
	assert.IsType(t, &sqlparser.Select{}, out)
}

func TestParserShouldReturnErrorWhenTheParsingFails(t *testing.T) {
	parser := VSQLParser{}
	_, err := parser.Parse("select")
	assert.Error(t, err)
}

func TestParserForCreateTableStatement(t *testing.T) {
	parser := VSQLParser{}
	out, _ := parser.Parse("create table foobar (bar text);")
	ctable := out.(*sqlparser.CreateTable)
	assert.IsType(t, &sqlparser.CreateTable{}, out)
	assert.Equal(t, 1, len(ctable.Columns))
	assert.Equal(t, "text", ctable.Columns[0].Type)
}
