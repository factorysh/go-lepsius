package parser

import (
	"github.com/bearstech/go-lepsius/conf"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestGrok(t *testing.T) {

	var book conf.Book
	err := yaml.Unmarshal([]byte(`---

filters:
 - name: grok
   args:
     pattern: "%{HAPROXYHTTP}"

`), &book)
	if err != nil {
		t.Error(err)
	}
	g := &Grok{}
	err = g.Configure(book.Filters[0].Args)
	if err != nil {
		t.Error(err)
	}
	v, err := g.Parse(`78.40.125.71:36602 [29/Oct/2015:23:59:29.957] http-in~ httpd/backend1 2488/0/0/1313/3801 200 423 - - ---- 1/1/0/1/0 0/0 "GET /test.php HTTP/1.1"`)
	if err != nil {
		t.Error(err)
	}
	t.Log("Grok value:", v)
}
