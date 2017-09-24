.PHONY: poc
export GOPATH=$(shell pwd)/gopath

build: vendor gp
	go build ./cli/lepsius

gopath/src/github.com/bearstech/go-lepsius:
	mkdir -p gopath/src/github.com/bearstech/
	ln -s ../../../.. gopath/src/github.com/bearstech/go-lepsius

gp: gopath/src/github.com/bearstech/go-lepsius vendor

vendor:
	glide install

test: gp
	go test -v github.com/bearstech/go-lepsius

clean:
	rm -rf gopath vendor

poc: gp
	rm -f gopath/src/github.com/bearstech/go-lepsius
	ln -s /go/ gopath/src/github.com/bearstech/go-lepsius
	docker run -it --rm -v `pwd`:/go -e GOPATH=/go/gopath golang go build -o poc/lepsius ./cli/lepsius
	rm -f gopath/src/github.com/bearstech/go-lepsius
