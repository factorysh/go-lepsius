package tick

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"io"

	"github.com/vjeantet/grok"
)

var hashes map[string]func() hash.Hash

func init() {
	hashes = map[string]func() hash.Hash{
		"sha1": sha1.New,
	}
}

type FilterNode interface {
	DoFilter(*Line) error
}

type GrokFilter struct {
	Node
	Source string
	Match  string
	grok   *grok.Grok
}

func (gf *GrokFilter) DoFilter(in *Line) error {
	//TODO
	return nil
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
	Format     string
	Target     string
}

func (fp *FingerprintFilter) DoFilter(in *Line) error {
	//TODO
	h, ok := hashes[fp.Method]
	if !ok {
		return fmt.Errorf("Hash method not found : %s", fp.Method)
	}
	hh := h()
	for _, s := range fp.SourceList {
		io.WriteString(hh, fmt.Sprintf("%v", in.Data[s]))
	}
	var v interface{}
	switch fp.Format {
	case "base64":
		v = base64.StdEncoding.EncodeToString(hh.Sum(nil))
	default:
		v = hh.Sum(nil)
	}
	in.Data[fp.Target] = v
	return nil
}

func (fp *FingerprintFilter) Source(sources ...string) *FingerprintFilter {
	fp.SourceList = append(fp.SourceList, sources...)
	return fp
}

func NewFingerprintFilter() *FingerprintFilter {
	return &FingerprintFilter{
		Method: "sha1",
		Format: "base64",
	}
}
