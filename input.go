package lepsius

import (
	"fmt"
	"github.com/bearstech/go-lepsius/conf"
	_input "github.com/bearstech/go-lepsius/input"
	_parser "github.com/bearstech/go-lepsius/parser"
)

type Input interface {
	conf.Configurable
	Lines() chan string
}

type Parser interface {
	conf.Configurable
	Parse(string) (*map[string]string, error)
}

type Reader interface {
	conf.Configurable
	Read(*map[string]string) error
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
		input = &_input.Tail{}
	} else {
		return nil, fmt.Errorf("Input %s not found", _conf.Input.Name)
	}
	err := input.Configure(_conf.Input.Args)
	if err != nil {
		return nil, err
	}
	var parser Parser
	if _conf.Parser.Name == "grok" {
		parser = &_parser.Grok{}
	} else {
		return nil, fmt.Errorf("Filter %s not not found", _conf.Parser.Name)
	}
	err = parser.Configure(_conf.Parser.Args)
	if err != nil {
		return nil, err
	}

	return &Lepsius{
		input:  input,
		parser: parser,
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
