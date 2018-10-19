package parser

import (
	"github.com/factorysh/go-lepsius/model"
)

var Parser map[string]model.Parser

func register(name string, parser model.Parser) {
	if Parser == nil {
		Parser = make(map[string]model.Parser)
	}
	Parser[name] = parser
}
