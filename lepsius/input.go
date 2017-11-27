package lepsius

import (
	"fmt"
	"github.com/bearstech/go-lepsius/conf"
	_input "github.com/bearstech/go-lepsius/input"
	"github.com/bearstech/go-lepsius/model"
	_parser "github.com/bearstech/go-lepsius/parser"
)

type Lepsius struct {
	input   model.Input
	parser  model.Parser
	filters []model.Filter
	reader  model.Reader
}

func LepsiusFromBook(_conf *conf.Book) (*Lepsius, error) {
	var input model.Input
	switch _conf.Input.Name {
	case "tail":
		input = &_input.Tail{}
	default:
		return nil, fmt.Errorf("Input %s not found", _conf.Input.Name)
	}
	err := input.Configure(_conf.Input.Args)
	if err != nil {
		return nil, fmt.Errorf("Section: Input Conf: %s %s", _conf.Input.Args,
			err.Error())
	}
	var parser model.Parser
	switch _conf.Parser.Name {
	case "grok":
		parser = &_parser.Grok{}
	default:
		return nil, fmt.Errorf("Parser %s not not found", _conf.Parser.Name)
	}
	err = parser.Configure(_conf.Parser.Args)
	if err != nil {
		return nil, fmt.Errorf("Section: Parser Conf:%s %s", _conf.Parser.Args,
			err.Error())
	}

	return &Lepsius{
		input:  input,
		parser: parser,
	}, nil
}

func (l *Lepsius) Serve() error {
	for line := range l.input.Lines() {
		event, err := l.parser.Parse(line.Message)
		if err != nil {
			// log something
		} else {
			if line.Values != nil {
				for k, v := range line.Values {
					event[k] = v
				}
			}
			err = l.reader.Read(event)
			if err != nil {
				// log something
			}
			// the line is correct
		}
	}
	return nil
}
