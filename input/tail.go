package input

import (
	"github.com/factorysh/go-lepsius/model"
	"github.com/factorysh/go-lepsius/parser"
	_tail "github.com/hpcloud/tail"
	"github.com/mitchellh/mapstructure"
)

func init() {
	register("tail", &Tail{})
}

type Tail struct {
	tail   *_tail.Tail
	config *TailConf
	parser model.Parser
}

func (i *Tail) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	go func() {
		for line := range i.tail.Lines {
			l, err := i.parser.Parse([]byte(line.Text))
			if err != nil {
				panic(err)
			}
			lines <- &model.Line{
				Values: map[string]interface{}{
					"message": l,
				},
			}
		}
	}()
	return lines
}

type TailConf struct {
	path   string
	parser string
}

func (t *Tail) Configure(conf map[string]interface{}) error {
	cfg := &TailConf{}
	err := mapstructure.Decode(conf, cfg)
	if err != nil {
		return err
	}
	if cfg.parser == "" {
		cfg.parser = "raw"
	}
	t.config = cfg
	t.tail, err = _tail.TailFile(t.config.path, _tail.Config{Follow: true})
	if err != nil {
		return err
	}
	t.parser = parser.Parser[t.config.parser]
	return nil
}
