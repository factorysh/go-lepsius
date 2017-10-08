package model

import (
	"github.com/bearstech/go-lepsius/conf"
)

type Line struct {
	Message string
	Values  map[string]string
}

type Input interface {
	conf.Configurable
	Lines() chan *Line
}

type Parser interface {
	conf.Configurable
	Parse(string) (map[string]string, error)
}

type Filter interface {
	conf.Configurable
	Filter(map[string]string) (map[string]string, error)
}

type Reader interface {
	conf.Configurable
	Read(map[string]string) error
}
