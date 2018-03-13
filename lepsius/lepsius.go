package lepsius

import (
	"errors"
	"fmt"
	"github.com/bearstech/go-lepsius/conf"
	_filter "github.com/bearstech/go-lepsius/filter"
	_input "github.com/bearstech/go-lepsius/input"
	"github.com/bearstech/go-lepsius/model"
	_output "github.com/bearstech/go-lepsius/output"
)

type Lepsius struct {
	input  model.Input
	filter []model.Filter
	output []model.Output
}

func LepsiusFromBook(_conf *conf.Book) (*Lepsius, error) {
	lepsius := &Lepsius{
		filter: make([]model.Filter, 0),
		output: make([]model.Output, 0),
	}

	if len(_conf.Input) != 1 {
		return nil, errors.New("You need one input")
	}
	i := _conf.Input[0]
	input, ok := _input.Input[i.Name()]
	if !ok {
		return nil, fmt.Errorf("Unknown input : %s", i.Name())
	}

	err := input.Configure(i.Args())
	if err != nil {
		return nil, fmt.Errorf("Section: Input Conf: %s %s", i.Args(),
			err.Error())
	}
	lepsius.input = input
	for _, u := range _conf.Filter {
		err = u.Validate()
		if err != nil {
			return nil, err
		}
		f, ok := _filter.Filter[u.Name()]
		if !ok {
			return nil, fmt.Errorf("Input unknown : %s", u.Name())
		}
		err = f.Configure(u.Args())
		if err != nil {
			return nil, fmt.Errorf("Input args for %s : %v", u.Name(), u.Args())
		}
		lepsius.filter = append(lepsius.filter, f)
	}
	if len(_conf.Output) == 0 {
		return nil, errors.New("Without output, you will ear nothing")
	}
	for _, u := range _conf.Output {
		err = u.Validate()
		if err != nil {
			return nil, err
		}
		f, ok := _output.Output[u.Name()]
		if !ok {
			return nil, fmt.Errorf("Output unknown : %s", u.Name())
		}
		err = f.Configure(u.Args())
		if err != nil {
			return nil, fmt.Errorf("Output args for %s : %v", u.Name(), u.Args())
		}
		lepsius.output = append(lepsius.output, f)
	}
	return lepsius, nil
}

func (l *Lepsius) Serve() error {
	for line := range l.input.Lines() {
		for _, r := range l.output {
			err := r.Read(line.Values)
			if err != nil {
				fmt.Println(err)
				// log something
			}
			// the line is correct
		}
	}
	return nil
}
