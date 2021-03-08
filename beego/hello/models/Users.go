package models

import (
	"github.com/astaxie/beego/orm"
	//"github.com/beego/beego/v2/client/orm"
)

type Users struct {
	Id      int
	Name    string
	AddTime int64
	Status  int
	Mobile  string
	Avatar  string
}

func init() {
	orm.RegisterModel(new(Users))
}

// 获取记录
func UserInfo(id int) (Users, error) {
	var (
		err error
	)

	o := orm.NewOrm()
	user := Users{Id: id}
	err = o.Read(&user)
	return user, err
}
