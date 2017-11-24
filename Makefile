.PHONY: poc

build: | vendor bin
	go build -o bin/lepsius ./cli/lepsius

bin:
	mkdir -p bin

vendor:
	glide install

test: | vendor
	go test -v github.com/bearstech/go-lepsius
	go test -v github.com/bearstech/go-lepsius/filter
	go test -v github.com/bearstech/go-lepsius/parser
	go test -v github.com/bearstech/go-lepsius/reader

clean:
	rm -rf vendor

poc: | vendor
	rm -f gopath/src/github.com/bearstech/go-lepsius
	ln -s /go/ gopath/src/github.com/bearstech/go-lepsius
	docker run -it --rm -v `pwd`:/go -e GOPATH=/go/gopath golang go build -o poc/lepsius ./cli/lepsius
	rm -f gopath/src/github.com/bearstech/go-lepsius
