package routers

import (
	"hello/controllers"

	//"github.com/astaxie/beego"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/hello", &controllers.MainController{}, "get:GetHello")
	beego.Include(&controllers.UserController{})
}
