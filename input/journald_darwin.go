package input

import (
	"github.com/factorysh/go-lepsius/model"
)

type Journald struct {
}

func (i *Journald) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	return lines
}

func (i *Journald) Configure(conf map[string]interface{}) error {
	return nil
}
