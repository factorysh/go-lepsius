package flux

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/interpreter"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/flux/semantic"
	"github.com/influxdata/flux/values"
)

func TestFlux(t *testing.T) {
	ql := `
	import "input"
	a = 1+1
	b = a *2
	//a |> yield()
	`
	pkg := parser.ParseSource(ql)
	if ast.Check(pkg) > 0 {
		t.Fatal(ast.GetError(pkg))
	}
	graph, err := semantic.New(pkg)
	if err != nil {
		t.Fatal(err)
	}

	// Create new interpreter for each test case
	itrp := interpreter.NewInterpreter()
	var testScope = interpreter.NewNestedScope(nil, values.NewObjectWithValues(
		map[string]values.Value{
			"true":  values.NewBool(true),
			"false": values.NewBool(false),
		}))
	sideEffects, err := itrp.Eval(graph, testScope, flux.StdLib())
	assert.NoError(t, err)
	fmt.Println(sideEffects)
	fmt.Println(testScope)
	testScope.Range(func(k string, v values.Value) {
		fmt.Println(k, v)
	})
}
