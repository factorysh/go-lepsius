export GOPATH=$(shell pwd)/gopath


lib: gopath/src/github.com/bearstech/go-lepsius gopath/src/github.com/vjeantet/grok gopath/src/gopkg.in/mcuadros/go-syslog.v2

gopath/src/github.com/vjeantet/grok:
	go get github.com/vjeantet/grok

gopath/src/gopkg.in/mcuadros/go-syslog.v2:
	go get gopkg.in/mcuadros/go-syslog.v2

gopath/src/github.com/bearstech/go-lepsius:
	mkdir -p gopath/src/github.com/bearstech
	ln -s `pwd` gopath/src/github.com/bearstech/go-lepsius

clean:
	rm -rf gopath
