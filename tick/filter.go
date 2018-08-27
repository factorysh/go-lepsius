package tick

import (
	"github.com/vjeantet/grok"
)

type FilterNode interface {
	DoFilter(*Line) (*Line, error)
}

type GrokFilter struct {
	Node
	Source string
	Match  string
	grok   *grok.Grok
}

func (gf *GrokFilter) DoFilter(in *Line) (*Line, error) {
	//TODO
	return in, nil
}

func NewGrokFilter() *GrokFilter {
	return &GrokFilter{
		Source: "message",
	}
}

type FingerprintFilter struct {
	Node
	Method     string
	SourceList []string `tick:"Source" json:"source"`
	Target     string
}

func (fp *FingerprintFilter) DoFilter(in *Line) (*Line, error) {
	//TODO
	return in, nil
}

func (fp *FingerprintFilter) Source(sources ...string) *FingerprintFilter {
	fp.SourceList = append(fp.SourceList, sources...)
	return fp
}

func NewFingerprintFilter() *FingerprintFilter {
	return &FingerprintFilter{
		Method: "sha256",
	}
}
