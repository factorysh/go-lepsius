.PHONY: poc
export GOPATH=$(shell pwd)/gopath

build: lib
	go build ./cli/lepsius

lib: gopath/src/github.com/bearstech/go-lepsius gopath/src/github.com/vjeantet/grok gopath/src/gopkg.in/mcuadros/go-syslog.v2

gopath/src/github.com/vjeantet/grok:
	go get github.com/vjeantet/grok

gopath/src/gopkg.in/mcuadros/go-syslog.v2:
	go get gopkg.in/mcuadros/go-syslog.v2

gopath/src/github.com/bearstech/go-lepsius:
	mkdir -p gopath/src/github.com/bearstech
	ln -s `pwd` gopath/src/github.com/bearstech/go-lepsius

test:
	go test -v github.com/bearstech/go-lepsius

clean:
	rm -rf gopath

poc:
	rm -f gopath/src/github.com/bearstech/go-lepsius
	ln -s /go/ gopath/src/github.com/bearstech/go-lepsius
	docker run -it --rm -v `pwd`:/go -e GOPATH=/go/gopath golang go build -o poc/lepsius ./cli/lepsius
	rm -f gopath/src/github.com/bearstech/go-lepsius
