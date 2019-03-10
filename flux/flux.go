package flux

import (
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/flux/semantic"
)

func BuildPackage(source string) (*semantic.Package, error) {
	pkg := parser.ParseSource(source)
	if ast.Check(pkg) > 0 {
		return nil, ast.GetError(pkg)
	}
	return semantic.New(pkg)
}
