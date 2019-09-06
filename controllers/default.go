package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"liyangweb/models"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

func (c * MainController) Index() {
	//if options, err := models.GetAllOptions(query, fields, sortby, order, 10, 10); err != nil {
	//	c.Ctx.WriteString(err.Error())
	//}
	//配置项
	options, err := models.GetAllOptions()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	//文章列表
	blogs, err := models.GetLogsForHome()
	fmt.Println(blogs)

	//c.Data["Options"] = options
	for _, item := range options {
		c.Data[item.OptionName] = item.OptionValue
	}
	c.Data["blogs"] = blogs

	c.TplName = "index.tpl"
}
