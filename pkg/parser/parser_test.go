package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserShouldReturnNonNilValue(t *testing.T) {
	out, _ := Parse("select now()")
	assert.NotNil(t, out)
}

func TestParserShouldReturnErrorWhenTheParsingFails(t *testing.T) {
	_, err := Parse("select")
	assert.Error(t, err)
}
