.PHONY: poc

build: | vendor bin
	go build -o bin/lepsius ./cli/lepsius

bin:
	mkdir -p bin

vendor:
	dep ensure

test: | vendor
	go test -v github.com/bearstech/go-lepsius
	go test -v github.com/bearstech/go-lepsius/filter
	go test -v github.com/bearstech/go-lepsius/parser
	go test -v github.com/bearstech/go-lepsius/reader

src/logstash-patterns-core:
	mkdir -p src
	cd src && git clone https://github.com/logstash-plugins/logstash-patterns-core.git

clean:
	rm -rf vendor

poc: | vendor
	rm -f gopath/src/github.com/bearstech/go-lepsius
	ln -s /go/ gopath/src/github.com/bearstech/go-lepsius
	docker run -it --rm -v `pwd`:/go -e GOPATH=/go/gopath golang go build -o poc/lepsius ./cli/lepsius
	rm -f gopath/src/github.com/bearstech/go-lepsius

linux: | vendor
	mkdir -p bin/linux
	docker run --rm \
	-v `pwd`:/go/src/github.com/bearstech/go-lepsius \
	-w /go/src/github.com/bearstech/go-lepsius \
	lepsius-dev \
	go build -o bin/linux/lepsius ./cli/lepsius

image-tool:
	docker build -t lepsius-dev contrib
