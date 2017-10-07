package lepsius

import (
	"fmt"
	"github.com/bearstech/go-lepsius/conf"
	"github.com/bearstech/go-lepsius/tail"
)

type Input interface {
	conf.Configurable
	Lines() chan string
}

type Parser interface {
	conf.Configurable
	Parse(string) (Event, error)
}

type Event map[string]string

type Reader interface {
	conf.Configurable
	Read(Event) error
}

type Lepsius struct {
	input  Input
	parser Parser
	reader Reader
}

func New(input Input, parser Parser, reader Reader) *Lepsius {
	return &Lepsius{
		input,
		parser,
		reader,
	}
}

func LepsiusFromBook(_conf *conf.Book) (*Lepsius, error) {
	var input Input
	if _conf.Input.Name == "tail" {
		input = &tail.Input{}
	} else {
		return nil, fmt.Errorf("Input %s not found", _conf.Input)
	}
	err := input.Configure(_conf.Input.Args)
	if err != nil {
		return nil, err
	}
	return &Lepsius{
		input: input,
	}, nil
}

func (l *Lepsius) Serve() error {
	for line := range l.input.Lines() {
		event, err := l.parser.Parse(line)
		if err != nil {
			// log something
		} else {
			err = l.reader.Read(event)
			if err != nil {
				// log something
			}
			// the line is correct
		}
	}
	return nil
}
