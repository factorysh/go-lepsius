package lepsius

import (
	"testing"

	"github.com/factorysh/go-lepsius/conf"
	"gopkg.in/yaml.v2"
)

func TestFromBook(t *testing.T) {
	var book conf.Book
	err := yaml.Unmarshal([]byte(`---
input:
  - tail:
      path: /tmp/toto
      parser:
        grok:
          pattern: "%{HAPROXYHTTP}"

output:
  - stdout:

`), &book)
	if err != nil {
		t.Error(err)
	}
	t.Log(book)
	l, err := LepsiusFromBook(&book)
	if err != nil {
		t.Error(err)
	}
	t.Log(l)
}
