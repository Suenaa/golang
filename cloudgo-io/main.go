package main

import (
	flag "github.com/spf13/pflag"
	"github.com/Suenaa/golang/cloudgo-io/service"
)

const (
	PORT string = "8080"
)

func main() {
	var port string
	var p string
	flag.StringVarP(&p, "port", "p", PORT, "port for httpd listening")
	flag.Parse()

	if len(p) != 0 {
		port = p
	}

	m := service.NewServer(port)
	service.RunServer(m)
}