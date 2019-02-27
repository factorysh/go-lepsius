package flux

import (
	"errors"
	"fmt"

	"github.com/influxdata/flux/parser"
	"github.com/influxdata/flux/semantic"
)

func Parse(source string) (*semantic.Package, error) {
	p := parser.ParseSource(source)
	if len(p.Errors) > 0 {
		fmt.Println(p.Errors)
		return nil, errors.New("Parsing error")
	}

	return semantic.New(p)
}
