package input

import (
	"github.com/bearstech/go-lepsius/model"
)

var Input map[string]model.Input

func register(name string, input model.Input) {
	if Input == nil {
		Input = make(map[string]model.Input)
	}
	Input[name] = input
}
