package input

import (
	"bufio"
	"github.com/bearstech/go-lepsius/model"
	"os"
)

type Stdin struct {
}

func (s *Stdin) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			lines <- &model.Line{
				Message: line,
			}
		}
	}()
	return lines
}

func (s *Stdin) Configure(conf map[string]interface{}) error {
	return nil
}
