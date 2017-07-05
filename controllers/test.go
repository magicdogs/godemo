package controllers

import "github.com/astaxie/beego"
import (
	_ "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
	"godemo/models"
	"fmt"
)

type TestController struct {
	beego.Controller
}


func (c *TestController) Get() {
	conn := orm.NewOrm()
	hotel := models.Hotel{Id:1}
	err := conn.Read(&hotel)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(hotel.Id, hotel.Name)
	}
	c.Data["name"] = "zhangsan"
	c.Data["age"] = 18
	//c.TplName = "test.tpl"
	c.TplName = "bootstrap.tpl";
}