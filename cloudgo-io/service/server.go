package service

import(
	"github.com/go-martini/martini"
	"net/http"
	"github.com/unrolled/render"
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
		Extensions: []string{".html"},
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
		var rt string
		if req.Form["username"][0] == "abc" && req.Form["password"][0] == "123" {
			rt = "login successfully"
		} else {
			rt = "username or password is wrong"
		}
		formatter.HTML(res, http.StatusOK, "success", struct {
			Content string `json:"content"`
			} {Content: rt})
		})
	m.Get("/unknow", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(501)
		})
	m.RunOnAddr(":" + port)
}
