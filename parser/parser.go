package parser

import (
	"github.com/bearstech/go-lepsius/model"
)

var Parser map[string]model.Parser

func register(name string, parser model.Parser) {
	if Parser == nil {
		Parser = make(map[string]model.Parser)
	}
	Parser[name] = parser
}
