package tick

import (
	"fmt"
	"testing"

	"github.com/influxdata/kapacitor/tick"
	"github.com/influxdata/kapacitor/tick/stateful"
	"github.com/stretchr/testify/assert"
)

func TestTick(t *testing.T) {
	script := `
var i = input
	|fromStdin()
	|stdout()
`
	scope := stateful.NewScope()
	input := &Input{
		Test: true,
	}
	scope.Set("input", input)

	r, err := tick.Evaluate(script, scope, nil, false)
	assert.NoError(t, err)
	fmt.Println(r)
	i, err := scope.Get("i")
	assert.NoError(t, err)
	s, ok := i.(*Stdout)
	assert.True(t, ok)
	fmt.Println(s)
	fmt.Println(s.Input.Test)

}
