package flux

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/factorysh/go-lepsius/flux/output" // output flux
	"github.com/factorysh/go-lepsius/flux/query"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/ast"
	_ "github.com/influxdata/flux/builtin"
	"github.com/influxdata/flux/interpreter"
	"github.com/influxdata/flux/lang"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/flux/semantic"
	"github.com/influxdata/flux/values"
)

func TestFlux(t *testing.T) {
	ql := `
	a = 1+1
	b = a *2
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

func TestVanilla(t *testing.T) {
	ql := `
		import "csv"
		csv.from(file: "toto.csv")
	`
	c := lang.FluxCompiler{
		Query: ql,
	}

	//querier := cmd.NewQuerier()
	querier := query.New(os.Stdout)
	//result, err := querier.Query(context.Background(), c)
	err := querier.Query(context.Background(), c)
	assert.NoError(t, err)
	//defer result.Release()

	//encoder := csv.NewMultiResultEncoder(csv.DefaultEncoderConfig())
	//_, err = encoder.Encode(os.Stdout, result)

	//assert.NoError(t, err)
}
