package input

import (
	"context"
	"fmt"
	"syscall"
	"testing"

	"github.com/containerd/fifo"
	"github.com/davecgh/go-spew/spew"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/interpreter"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/flux/semantic"
	"github.com/influxdata/flux/values"
	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	ctx := context.Background()
	f, err := fifo.OpenFifo(ctx, "/tmp/lepsius", syscall.O_CREAT+syscall.O_RDWR,
		0660)
	assert.NoError(t, err)
	ql := `
	import "input"
	p = input.pipe(path:"/tmp/lepsius")
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
			"true":  values.NewBool(true),
			"false": values.NewBool(false),
		}))
	sideEffects, err := itrp.Eval(graph, testScope, flux.StdLib())
	assert.NoError(t, err)
	fmt.Println(sideEffects)
	spew.Dump(testScope)
	testScope.Range(func(k string, v values.Value) {
		fmt.Println(k, "Type: ", v.Type(), v.Type().Nature())
		if v.Type().Nature() == semantic.Object {
			fmt.Println("\tobject: ", v.Object())
			v.Object().Range(func(kk string, vv values.Value) {
				fmt.Println("\t", kk, vv)
				spew.Dump(vv)
			})
		}
	})
	f.Write([]byte("Hello world"))
}
