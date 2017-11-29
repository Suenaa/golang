package entities

import "github.com/go-xorm/xorm"

var engine *xorm.Engine

func init() {
	e, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3305)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	engine = e
	engine.Sync2(new(UserInfo))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}