package controllers

import (
	"hello/models"

	beego "github.com/beego/beego/v2/server/web"
)

type DemoController struct {
	beego.Controller
}

func (this *DemoController) GetHello() {
	this.Ctx.WriteString("demo-hello")
}

func (this *DemoController) GetUsername() {
	var (
		id    int
		err   error
		title string
		user  models.Users
	)
	id, err = this.GetInt("id")
	user, err = models.UserInfo(id)
	if err != nil {
		title = "404"
	} else {
		title = user.Name
	}

	this.Ctx.WriteString(title)
}
