package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"liyangweb/controllers"
	"liyangweb/models"
	"strconv"
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

	beego.Get("/test", func(context *context.Context) {
		orm.Debug = true
		o := orm.NewOrm()
		qs := o.QueryTable(new(models.Blog))
		qs = qs.Filter("type", "blog")
		qs = qs.Filter("hide", "n")
		//qs = qs.RelatedSel()//是否要关联查询
		count,_ := qs.Count()
		context.WriteString(strconv.FormatInt(count, 10))
	})

    //关于我(这个还需要改)
    beego.Router("/about", &controllers.MeController{}, "get:About")
}
