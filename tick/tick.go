package tick

import (
	"bufio"
	"os"
)

type Input struct {
	Test  bool
	Debug bool
}

func (i *Input) FromStdin() *FromStdin {
	f := &FromStdin{}
	f.Input = i
	f.New()
	return f
}

type Line map[string]interface{}

type Node struct {
	Input  *Input
	Events chan *Line
}

func (n *Node) Stdout() *Stdout {
	s := &Stdout{}
	s.Events = n.Events
	s.Input = n.Input
	return s
}

type Eventsable interface {
	Events() chan *Line
}

type OutAble interface {
	Out() chan *Line
}

type FromStdin struct {
	Node
}

func (f *FromStdin) New() {
	f.Events = make(chan *Line)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			line := scanner.Bytes()
			f.Events <- &Line{
				"message": line,
			}
		}
	}()
}

type Stdout struct {
	Node
}
