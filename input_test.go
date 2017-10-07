package lepsius

import (
	"github.com/bearstech/go-lepsius/conf"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestFromBook(t *testing.T) {
	var book conf.Book
	err := yaml.Unmarshal([]byte(`---

input:
  name: tail
  args:
    path: /tmp/toto

parser:
  name: grok
  args:
    pattern: |
      %{HAPROXYHTTP}

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
