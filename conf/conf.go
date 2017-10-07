package conf

import ()

type Book struct {
	Input struct {
		Name string                 `yaml:"name"`
		Args map[string]interface{} `yaml:"args"`
	} `yaml:"input"`
	Parser struct {
		Name string
		Args map[string]interface{}
	}
	Reader struct {
		Name string
		Args map[string]interface{}
	}
}

type Configurable interface {
	Configure(conf map[string]interface{}) error
}
