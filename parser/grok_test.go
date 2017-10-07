package parser

import (
	"github.com/bearstech/go-lepsius/conf"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestGrok(t *testing.T) {

	var book conf.Book
	err := yaml.Unmarshal([]byte(`---

filter:
  name: grok
  args:
    pattern: %{HAPROXYHTTP}

`), &book)
	if err != nil {
		t.Error(err)
	}
}
