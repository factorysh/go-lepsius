package tick

import (
	"github.com/factorysh/go-lepsius/tick/filter"
	"github.com/factorysh/go-lepsius/tick/model"
	"github.com/influxdata/kapacitor/tick/ast"
	"github.com/vjeantet/grok"
)

type FilterNode interface {
	Node
	DoFilter(*model.Line) error
}

type GrokFilter struct {
	node
	Source string
	Match  string
	grok   *grok.Grok
}

func NewGrokFilter() *GrokFilter {
	return &GrokFilter{
		Source: "message",
	}
}

func (g *GrokFilter) DoFilter(in *model.Line) error {
	//TODO
	return nil
}

type FingerprintFilter struct {
	node
	Method     string
	SourceList []string `tick:"Source" json:"source"`
	Format     string
	Target     string
}

func NewFingerprintFilter() *FingerprintFilter {
	return &FingerprintFilter{
		Method: "sha1",
		Format: "base64",
	}
}

func (fp *FingerprintFilter) Source(sources ...string) *FingerprintFilter {
	fp.SourceList = append(fp.SourceList, sources...)
	return fp
}

// tick:ignore
func (fp *FingerprintFilter) DoFilter(in *model.Line) error {
	return filter.DoFingerprintFilter(fp.Method, fp.Format, fp.SourceList, fp.Target, in)
}

type WhereFilter struct {
	node
	lambda *ast.LambdaNode
}
