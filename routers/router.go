package routers

import (
	"liyangweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //首页
    //http://localhost:8080/
    beego.Router("/", &controllers.MainController{}, "get:Index")
	//http://localhost:8080/?p=2
    beego.Router("page", &controllers.MainController{}, "get:Index")

    //详情页
    //http://localhost:8080/service/342.html
    beego.Router("/:sort:string/?:gid([0-9]+).html", &controllers.MainController{}, "get:Content")
    //http://localhost:8080/1.html
	beego.Router("/:gid([0-9]+).html", &controllers.MainController{}, "get:Content")



    //关于我(这个还需要改)
    beego.Router("/about", &controllers.MeController{}, "get:About")
}
