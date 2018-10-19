package tick

import (
	"fmt"
	"testing"

	"github.com/factorysh/go-lepsius/tick/model"
	"github.com/influxdata/kapacitor/tick"
	"github.com/influxdata/kapacitor/tick/stateful"
	"github.com/stretchr/testify/assert"
)

func TestTick(t *testing.T) {
	script := `
var i = input
	|fromStdin()
		.parse(json)
	|stdout()

var fp = ['name', 'client']
var i2 = input
	|fromChan(chan)
	|grok()
		.source('message')
		.match('%{NUMBER:duration} %{IP:client}')
	|fingerprint()
		.source(fp)
		.target('uid')
	|stdout()

`
	scope := stateful.NewScope()
	input := NewInput()
	input.pipeline.Test = true
	c := make(chan *model.Line, 1)
	c2 := make(chan *model.Line, 1)
	scope.Set("input", input)
	scope.Set("json", JsonParser)
	scope.Set("chan", c)
	scope.Set("chan2", c2)

	r, err := tick.Evaluate(script, scope, nil, false)
	assert.NoError(t, err)
	fmt.Println(r)
	i, err := scope.Get("i")
	assert.NoError(t, err)
	assert.NotNil(t, i)
	s, ok := i.(*Stdout)
	assert.True(t, ok)
	fmt.Println(s)
	fmt.Println(s.Pipeline().Test)

	i2_, err := scope.Get("i2")
	assert.NoError(t, err)
	i2, ok := i2_.(*Stdout)
	assert.True(t, ok)
	assert.Len(t, i2.Pipeline().Filters, 2)

	l, err := model.NewLine("beuha", "aussi")
	assert.NoError(t, err)
	c <- l
	fmt.Println(i2)
	l2, err := i2.Pipeline().read()
	assert.NoError(t, err)
	fmt.Println(l2)

}

func TestLambda(t *testing.T) {
	script := `

var i = input
	|fromChan(chan)
	|where(lambda: "name" == 'robert')
	|stdout()

`
	scope := stateful.NewScope()
	input := NewInput()
	scope.Set("input", input)
	c := make(chan *model.Line, 1)
	scope.Set("chan", c)
	r, err := tick.Evaluate(script, scope, nil, false)
	assert.NoError(t, err)
	fmt.Println(r)
	i, err := scope.Get("i")
	assert.NoError(t, err)
	assert.NotNil(t, i)
	i2, ok := i.(*Stdout)
	assert.True(t, ok)

	l, err := model.NewLine("name", "robert")
	assert.NoError(t, err)
	c <- l
	l2, err := i2.Pipeline().read()
	assert.NoError(t, err)
	fmt.Println("line", l2)

}
