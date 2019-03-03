package flux

import (
	"context"
	"io"
	"math"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/control"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/flux/semantic"
	_ "github.com/influxdata/flux/stdlib/universe" // uinverse flux
)

func BuildPackage(source string) (*semantic.Package, error) {
	pkg := parser.ParseSource(source)
	if ast.Check(pkg) > 0 {
		return nil, ast.GetError(pkg)
	}
	return semantic.New(pkg)
}

type Querier struct {
	C *control.Controller
}

func New() *Querier {
	cfg := control.Config{
		ConcurrencyQuota: 1,
		MemoryBytesQuota: math.MaxInt64,
	}
	return &Querier{
		C: control.New(cfg),
	}

}

// Query does query
func (q *Querier) Query(ctx context.Context, w io.Writer, c flux.Compiler, d flux.Dialect) (int64, error) {
	query, err := q.C.Query(ctx, c)
	if err != nil {
		return 0, err
	}
	results := flux.NewResultIteratorFromQuery(query)
	defer results.Release()

	encoder := d.Encoder()
	return encoder.Encode(w, results)
}
