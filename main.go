package main

import (
	_ "godemo/routers"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"godemo/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "mscnut:hackerlusy@tcp(www.mscnut.top:3306)/test?charset=utf8",maxIdle,maxConn)
	orm.Debug=true
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	hotel := new(models.Hotel)
	hotel.Address = "sss"
	hotel.City = 652
	hotel.Name = "citys"
	hotel.Zip = "zip"

	//fmt.Println(o.Insert(hotel))
	beego.Run()
	if( 1 == 2) {
		doc, _ := goquery.NewDocument("http://sports.sina.com.cn")
		dhead := doc.Find("head")
		dcharset := dhead.Find("meta[http-equiv]")
		charset, _ := dcharset.Attr("content")
		fmt.Println(charset)
		req := httplib.Get("http://beego.me/")
		str, err := req.String()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	}
}

