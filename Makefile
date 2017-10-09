.PHONY: poc
export GOPATH=$(shell pwd)/gopath

build: gp
	cd gopath/src/github.com/bearstech/go-lepsius && go build ./cli/lepsius

gopath/src/github.com/bearstech/go-lepsius:
	mkdir -p gopath/src/github.com/bearstech/
	ln -sf ../../../.. gopath/src/github.com/bearstech/go-lepsius

gp: | vendor

vendor: | gopath/src/github.com/bearstech/go-lepsius
	glide install

test: gp
	go test -v github.com/bearstech/go-lepsius
	go test -v github.com/bearstech/go-lepsius/filter
	go test -v github.com/bearstech/go-lepsius/parser
	go test -v github.com/bearstech/go-lepsius/reader

clean:
	rm -rf gopath vendor

poc: gp
	rm -f gopath/src/github.com/bearstech/go-lepsius
	ln -s /go/ gopath/src/github.com/bearstech/go-lepsius
	docker run -it --rm -v `pwd`:/go -e GOPATH=/go/gopath golang go build -o poc/lepsius ./cli/lepsius
	rm -f gopath/src/github.com/bearstech/go-lepsius
