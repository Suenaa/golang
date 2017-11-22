package service

import(
	"github.com/go-martini/martini"
	"net/http"
	"github.com/unrolled/render"
	"fmt"
)

var port string

func NewServer(Port string) *martini.ClassicMartini {
	m := martini.Classic()
	port = Port
	return m
}

func RunServer(m *martini.ClassicMartini) {
	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
		IndentJSON: true,
		})
	m.Use(martini.Static("assets"))
	m.Get("/json", func(res http.ResponseWriter, req *http.Request) {
		formatter.JSON(res, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
			} {ID: "101010", Content: "Hello from Go!"})
		})
	m.Get("/login", func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.HTML(res, http.StatusOK, "login", " ")
		})
	m.Post("/login", func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		
		formatter.HTML(res, http.StatusOK, "success", struct{
			Username string
			Password string
			} {
				Username: req.Form["username"][0],
				Password: req.Form["password"][0]})

		fmt.Println("username:", req.Form["username"][0])
		fmt.Println("password:", req.Form["password"][0])
		})
	m.Get("/unknow", func(res http.ResponseWriter, req *http.Request) {
		http.Error(res, "501 page not implemented", http.StatusNotImplemented)
		})
	m.RunOnAddr(":" + port)
}
