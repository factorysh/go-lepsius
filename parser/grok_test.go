package parser

import (
	"github.com/bearstech/go-lepsius/conf"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestGrokHAproxy(t *testing.T) {

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
	v, err := g.Parse([]byte(`78.40.125.71:36602 [29/Oct/2015:23:59:29.957] http-in~ httpd/backend1 2488/0/0/1313/3801 200 423 - - ---- 1/1/0/1/0 0/0 "GET /test.php HTTP/1.1"`))
	if err != nil {
		t.Error(err)
	}
	for k, vv := range v {
		t.Logf("%v => %+v", k, vv)
	}
}

func TestGrokTraefikCLF(t *testing.T) {
	var book conf.Book
	err := yaml.Unmarshal([]byte(`---

filters:
 - name: grok
   args:
     pattern: "%{TRAEFIKCLF}"

`), &book)
	if err != nil {
		t.Error(err)
	}
	g := &Grok{}
	err = g.Configure(book.Filters[0].Args)
	if err != nil {
		t.Error(err)
	}
	v, err := g.Parse([]byte(`151.127.44.139 - - [14/Feb/2018:14:21:21 +0000] "GET /images/logo.png HTTP/2.0" 200 77420 "https://preprod.docs.factory.sh/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:59.0) Gecko/20100101 Firefox/59.0" 4 "Host-preprod-docs-factory-sh-0" "http://172.18.0.2:8000" 1ms`))
	if err != nil {
		t.Error(err)
	}
	for k, vv := range v {
		t.Logf("%v => %+v", k, vv)
	}
	//assert.Equal(t, v["client_ip"], "151.127.44.139", "")
	assert.Equal(t, "1ms", v["chrono"], "")
	assert.Equal(t, "200", v["response"], "")
	assert.Equal(t, "/images/logo.png", v["request"], "")
	assert.Equal(t, `"https://preprod.docs.factory.sh/"`, v["referrer"])
}
