package routers

import (
	"liyangweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //首页
    beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("page", &controllers.MainController{}, "get:Index")
    //关于我
    beego.Router("/about", &controllers.MeController{}, "get:About")
}
