package routers

import (
	"godemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router("/test",&controllers.TestController{})
}
