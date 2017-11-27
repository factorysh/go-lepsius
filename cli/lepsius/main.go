package main

import (
	"github.com/bearstech/go-lepsius/conf"
	"github.com/bearstech/go-lepsius/lepsius"
	"os"
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
