package models

import "github.com/astaxie/beego/orm"

type Hotel struct {
	Id    int
	City int
	Name  string
	Address string
	Zip string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Hotel))
}
