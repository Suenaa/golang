package entities

import(
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {
	e, err := xorm.NewEngine("mysql", "root:
		@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	err = e.Sync2(new(UserInfo))
	checkErr(err)
	engine = e
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}