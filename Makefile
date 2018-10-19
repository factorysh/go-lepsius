.PHONY: poc

build: | vendor bin
	go build -o bin/lepsius ./cli/lepsius

bin:
	mkdir -p bin

vendor:
	dep ensure

test: | vendor
	go test -v github.com/factorysh/go-lepsius
	go test -v github.com/factorysh/go-lepsius/filter
	go test -v github.com/factorysh/go-lepsius/parser
	go test -v github.com/factorysh/go-lepsius/output
	go test -v github.com/factorysh/go-lepsius/input

src/logstash-patterns-core:
	mkdir -p src
	cd src && git clone https://github.com/logstash-plugins/logstash-patterns-core.git

clean:
	rm -rf vendor

poc: | vendor
	rm -f gopath/src/github.com/factorysh/go-lepsius
	ln -s /go/ gopath/src/github.com/factorysh/go-lepsius
	docker run -it --rm -v `pwd`:/go -e GOPATH=/go/gopath golang go build -o poc/lepsius ./cli/lepsius
	rm -f gopath/src/github.com/factorysh/go-lepsius

linux: | vendor
	mkdir -p bin/linux
	docker run --rm \
	-v `pwd`:/go/src/github.com/factorysh/go-lepsius \
	-w /go/src/github.com/factorysh/go-lepsius \
	lepsius-dev \
	go build -o bin/linux/lepsius ./cli/lepsius

image-tool:
	docker build -t lepsius-dev contrib
