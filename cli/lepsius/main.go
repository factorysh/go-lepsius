package main

import (
	"fmt"

	"github.com/bearstech/go-lepsius"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {

	handler := lepsius.NewHandler("%{HAPROXYHTTP}")

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	err := server.ListenUDP("0.0.0.0:1514")
	if err != nil {
		fmt.Println("Oups", err)
	}
	server.Boot()
	server.Wait()
}
