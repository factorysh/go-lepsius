package tick

import (
	"bufio"
	"os"

	"github.com/factorysh/go-lepsius/tick/model"
)

type From struct {
	node
	Parse Parser
}

type FromStdin struct {
	From
}

func (f *FromStdin) New() {
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			line := scanner.Bytes()
			f.Pipeline().Events <- &model.Line{
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
}
