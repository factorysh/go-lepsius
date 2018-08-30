package tick

import (
	"github.com/vjeantet/grok"
	"gitlab.bearstech.com/bearstech/go-lepsius/tick/filter"
	"gitlab.bearstech.com/bearstech/go-lepsius/tick/model"
)

type FilterNode interface {
	DoFilter(*model.Line) error
}

type GrokFilter struct {
	Node
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
	Node
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

func (fp *FingerprintFilter) DoFilter(in *model.Line) error {
	return filter.DoFingerprintFilter(fp.Method, fp.Format, fp.SourceList, fp.Target, in)
}
