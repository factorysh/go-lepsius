package model

import (
	"github.com/bearstech/go-lepsius/conf"
)

type Line struct {
	Values map[string]interface{}
	Keep   bool
}

type Input interface {
	conf.Configurable
	Lines() chan *Line
}

type Parser interface {
	conf.Configurable
	Parse([]byte) (map[string]interface{}, error)
}

type Filter interface {
	conf.Configurable
	Filter(*Line) error
}

type Output interface {
	conf.Configurable
	Read(map[string]interface{}) error
}
