package easyLexML

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_counter2string(t *testing.T) {
	assert.Equal(t, "0", counter2string(0, 0))
	assert.Equal(t, "1", counter2string(1, 0))
	assert.Equal(t, "2", counter2string(2, 0))
	assert.Equal(t, "42", counter2string(42, 0))
	assert.Equal(t, "42-A", counter2string(42, 1))
	assert.Equal(t, "42-B", counter2string(42, 2))
	assert.Equal(t, "42-Z", counter2string(42, 26))
	assert.Equal(t, "42-AA", counter2string(42, 26+1))
	assert.Equal(t, "42-AB", counter2string(42, 26+2))
	assert.Equal(t, "42-AZ", counter2string(42, 26+26))
	assert.Equal(t, "42-AAA", counter2string(42, 26+26+1))
}
