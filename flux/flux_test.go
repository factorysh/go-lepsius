package flux

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/factorysh/go-lepsius/flux/output" // output flux
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/interpreter"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/flux/semantic"
	_ "github.com/influxdata/flux/stdlib" // universe flux
	"github.com/influxdata/flux/values"
)

func init() {
	flux.FinalizeBuiltIns()
}
func TestFlux(t *testing.T) {
	ql := `
	import "input"
	import "output"
	a = 1+1
	b = a *2
	p = input.pipe(path:"/tmp/lepsius")
	p |> output.spew()
	`
	pkg := parser.ParseSource(ql)
	if ast.Check(pkg) > 0 {
		t.Fatal(ast.GetError(pkg))
	}
	graph, err := semantic.New(pkg)
	assert.NoError(t, err)

	// Create new interpreter for each test case
	itrp := interpreter.NewInterpreter()
	var testScope = interpreter.NewNestedScope(nil, values.NewObjectWithValues(
		map[string]values.Value{
			//"true":  values.NewBool(true),
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
