package tick

import (
	"gitlab.bearstech.com/bearstech/go-lepsius/tick/model"
)

type Node interface {
	setPipeline(p *Pipeline)
	Pipeline() *Pipeline
	linkFilter(f FilterNode)
}

type node struct {
	pipeline *Pipeline
}

func (n *node) setPipeline(p *Pipeline) {
	n.pipeline = p
}

func (n *node) Pipeline() *Pipeline {
	return n.pipeline
}

type Pipeline struct {
	Test    bool
	Debug   bool
	Events  chan *model.Line
	Filters []FilterNode
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Test:    false,
		Debug:   false,
		Filters: make([]FilterNode, 0),
		Events:  make(chan *model.Line),
	}
}

type Input struct {
	pipeline *Pipeline
}

type Output interface {
	Pipeline() *Pipeline
}

func (n *node) linkFilter(filter FilterNode) {
	filter.setPipeline(n.pipeline)
	n.pipeline.Filters = append(n.pipeline.Filters, filter)
}

func NewInput() *Input {
	return &Input{
		pipeline: NewPipeline(),
	}
}

func (i *Input) FromStdin() *FromStdin {
	f := &FromStdin{}
	f.setPipeline(i.pipeline)
	//f.New()
	return f
}

func (i *Input) FromChan(c chan *model.Line) *FromChan {
	fc := &FromChan{}
	i.pipeline.Events = c
	fc.setPipeline(i.pipeline)
	return fc
}

func (p *Pipeline) read() (*model.Line, error) {
	line := <-p.Events
	for _, f := range p.Filters {
		err := f.DoFilter(line)
		if err != nil {
			return nil, err
		}
	}
	return line, nil
}

func (n *node) Grok() *GrokFilter {
	gf := NewGrokFilter()
	n.linkFilter(gf)
	return gf
}

func (n *node) Fingerprint() *FingerprintFilter {
	fp := NewFingerprintFilter()
	n.linkFilter(fp)
	return fp
}

func (n *node) Stdout() *Stdout {
	return &Stdout{
		pipeline: n.pipeline,
	}
}

type Stdout struct {
	pipeline *Pipeline
}

func (s *Stdout) Pipeline() *Pipeline {
	return s.pipeline
}
