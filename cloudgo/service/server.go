package service

import "github.com/go-martini/martini"

func NewServer(port string) {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello Golang!"
	})
	m.RunOnAddr(":" + port)
}
