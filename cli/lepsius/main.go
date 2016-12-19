package main

import (
	"github.com/bearstech/go-lepsius"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	handler, err := lepsius.NewHandler("%{HAPROXYHTTPDIRECT}")
	if err != nil {
		panic(err)
	}

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	err = server.ListenUDP("0.0.0.0:1514")
	if err != nil {
		panic(err)
	}
	server.Boot()
	server.Wait()
}
