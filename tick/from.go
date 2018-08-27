package tick

import (
	"bufio"
	"os"
)

type From struct {
	Parse  Parser
	Events chan *Line
}

type FromStdin struct {
	Node
	From
}

func (f *FromStdin) New() {
	f.Events = make(chan *Line)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			line := scanner.Bytes()
			f.Events <- &Line{
				Data: map[string]interface{}{
					"message": line,
				},
			}
		}
	}()
}

type FromChan struct {
	Node
	From
}

func (f *FromChan) New() {
	f.Events = make(chan *Line)
}
