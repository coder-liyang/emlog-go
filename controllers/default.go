package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"liyangweb/models"
)

type MainController struct {
	beego.Controller
}


//首页
func (c * MainController) Index() {
	var (
		page int64
	)
	//配置项
	options, err := models.GetAllOptions()
	//c.Data["Options"] = options
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	//文章列表
	//pageStr := c.GetString("p")
	//page, _ = strconv.ParseInt(pageStr, 10, 64)
	page, _ = c.GetInt64("p")
	blogs, err := models.GetLogsForHome(page)
	for _, item := range options {
		c.Data[item.OptionName] = item.OptionValue
	}
	c.Data["blogs"] = &blogs
	//分页数据
	count, err := models.BlogTotal()
	p := pagination.NewPaginator(c.Ctx.Request, 10, count)
	c.Data["paginator"] = p

	c.TplName = "index.tpl"
}
