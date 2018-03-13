package output

import (
	"github.com/bearstech/go-lepsius/model"
)

var Output map[string]model.Output

func register(name string, output model.Output) {
	if Output == nil {
		Output = make(map[string]model.Output)
	}
	Output[name] = output
}
