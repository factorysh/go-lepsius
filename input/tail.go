package input

import (
	"github.com/bearstech/go-lepsius/model"
	_tail "github.com/hpcloud/tail"
	"github.com/mitchellh/mapstructure"
)

type Tail struct {
	tail *_tail.Tail
}

func (i *Tail) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	go func() {
		for line := range i.tail.Lines {
			lines <- &model.Line{
				Message: line.Text,
			}
		}
	}()
	return lines
}

type TailConf struct {
	path string
}

func (i *Tail) Configure(conf map[string]interface{}) error {
	var tailconf TailConf
	err := mapstructure.Decode(conf, &tailconf)
	if err != nil {
		return err
	}
	tail, err := _tail.TailFile(tailconf.path, _tail.Config{Follow: true})
	if err == nil {
		i.tail = tail
	}
	return err
}
