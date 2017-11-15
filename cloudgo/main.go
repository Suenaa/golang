package main

import (
	"os"
	flag "github.com/spf13/pflag"
	"github.com/Suenaa/golang/cloudgo/service"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	var p string
	flag.StringVarP(&p, "port", "p", PORT, "port for httpd listening")
	flag.Parse()

	if len(p) != 0 {
		port = p
	}

	service.NewServer(port)
}