package filter

import (
	"github.com/bearstech/go-lepsius/model"
)

var Filter map[string]model.Filter

func register(name string, filter model.Filter) {
	if Filter == nil {
		Filter = make(map[string]model.Filter)
	}
	Filter[name] = filter
}
