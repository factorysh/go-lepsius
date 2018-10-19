package main

import (
	"os"

	"github.com/factorysh/go-lepsius/conf"
	"github.com/factorysh/go-lepsius/lepsius"
)

func main() {
	book, err := conf.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	l, err := lepsius.LepsiusFromBook(book)
	if err != nil {
		panic(err)
	}

	err = l.Serve()
	if err != nil {
		panic(err)
	}

}
