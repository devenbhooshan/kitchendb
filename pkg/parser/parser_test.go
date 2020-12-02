package parser;
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	assert.NotNil(t, Parse("select now();"), "they should be equal")
}