package tick

import (
	"gitlab.bearstech.com/bearstech/go-lepsius/tick/model"
)

type Node struct {
	Input *Input
}

type Input struct {
	Test    bool
	Debug   bool
	Events  chan *model.Line
	Filters []FilterNode
}

func NewInput() *Input {
	return &Input{
		Test:    false,
		Debug:   false,
		Filters: make([]FilterNode, 0),
	}
}

func (i *Input) FromStdin() *FromStdin {
	f := &FromStdin{}
	f.Input = i
	f.New()
	return f
}

func (i *Input) FromChan(c chan *model.Line) *FromChan {
	fc := &FromChan{}
	i.Events = c
	fc.Input = i
	return fc
}

func (i *Input) read() (*model.Line, error) {
	line := <-i.Events
	for _, f := range i.Filters {
		err := f.DoFilter(line)
		if err != nil {
			return nil, err
		}
	}
	return line, nil
}

func (n *Node) Grok() *GrokFilter {
	gf := NewGrokFilter()
	n.Input.Filters = append(n.Input.Filters, gf)
	gf.Input = n.Input
	return gf
}

func (n *Node) Fingerprint() *FingerprintFilter {
	fp := NewFingerprintFilter()
	n.Input.Filters = append(n.Input.Filters, fp)
	fp.Input = n.Input
	return fp
}

func (n *Node) Stdout() *Stdout {
	s := &Stdout{}
	s.Input = n.Input
	return s
}

type Stdout struct {
	Node
}
