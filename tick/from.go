package tick

import (
	"bufio"
	"os"

	"gitlab.bearstech.com/bearstech/go-lepsius/tick/model"
)

type From struct {
	Node
	Parse Parser
}

type FromStdin struct {
	From
}

func (f *FromStdin) New() {
	f.Input.Events = make(chan *model.Line)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			line := scanner.Bytes()
			f.Input.Events <- &model.Line{
				Data: map[string]interface{}{
					"message": line,
				},
			}
		}
	}()
}

type FromChan struct {
	From
}

func (f *FromChan) New() {
	f.Input.Events = make(chan *model.Line)
}
