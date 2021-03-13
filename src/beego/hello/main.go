package main

import (
	_ "hello/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//"github.com/beego/beego/v2/client/orm"
	//beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb)
}

func main() {

	beego.Run(":8080")
}
