package tail

import (
	"errors"
	_tail "github.com/hpcloud/tail"
)

type Input struct {
	tail *_tail.Tail
}

func (i *Input) Lines() chan string {
	lines := make(chan string)
	go func() {
		for line := range i.tail.Lines {
			lines <- line.Text
		}
	}()
	return lines
}

func (i *Input) Configure(conf map[string]interface{}) error {
	path_raw, ok := conf["path"]
	if !ok {
		return errors.New("path key is mandatory")
	}
	path, ok := path_raw.(string)
	if !ok {
		return errors.New("path type must be a string")
	}
	tail, err := _tail.TailFile(path, _tail.Config{Follow: true})
	if err == nil {
		i.tail = tail
	}
	return err
}
