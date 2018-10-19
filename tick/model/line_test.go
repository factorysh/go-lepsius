package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	line, err := NewLine("timestamp", 42, "data", "blah blah")
	assert.NoError(t, err)
	fmt.Println(line.Flatten())
}
