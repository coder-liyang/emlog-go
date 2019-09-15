package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/techoner/gophp"
	"liyangweb/models"
	"liyangweb/sides"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func setOptions(c *MainController) {
	//配置项
	options, err := models.GetAllOptions()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	for _, item := range options {
		c.Data[item.OptionName] = item.OptionValue
	}
	//侧边栏
	widgets1 := c.Data["widgets1"].(string)

	//widgets1通过unserialize后,变成[]interface{}类型了
	//[]interface {}
	//[search custom_wg_1 newcomm blogger twitter archive random_log calendar link custom_wg_4]
	unserialize, _ := gophp.Unserialize([]byte(widgets1))
	for _, val := range unserialize.([]interface{}) {
		sideName := val.(string)
		if sideName == "twitter" { //微语
			sides.Sides[sideName] = sides.WidgetTwitter()
		} else if sideName == "archive"  { //归档
			sides.Sides["archive"] = sides.WidgetArchive()
		} else if sideName == "link" {
			sides.Sides["link"] = sides.WidgetLink()
		} else if sideName == "blogger" {
			sides.Sides["blogger"] = sides.WidgetBlogger()[0:1]
		} else if sideName == "newcomm" {
			//sides.Sides["newcomm"] =
				sides.WidgetNewcomm()
		} else {
			sides.Sides[sideName] = make([]interface{}, 10)
		}
	}
	c.Data["sides"] = sides.Sides
}

//首页
func (c *MainController) Index() {
	var (
		page int64
	)
	setOptions(c)
	//fmt.Println(c.Data["sides"])
	//文章列表
	//pageStr := c.GetString("p")
	//page, _ = strconv.ParseInt(pageStr, 10, 64)
	page, _ = c.GetInt64("p")
	blogs, _ := models.GetLogsForHome(page)
	c.Data["blogs"] = &blogs
	//分页数据
	count, _ := models.BlogTotal()
	p := pagination.NewPaginator(c.Ctx.Request, 10, count)
	c.Data["paginator"] = p

	c.Layout = "layouts/front.tpl"
	c.TplName = "index.tpl"
}

//文章详情页
func (c *MainController) Content() {
	setOptions(c)
	gid, _ := strconv.ParseInt(c.Ctx.Input.Param(":gid"), 10, 64)
	blog, err := models.GetBlogByGid(gid)
	if err != nil {
		fmt.Println(err.Error())
		c.Ctx.WriteString(err.Error())
	}
	if blog == nil {
		fmt.Println("blog is nil")
	}
	c.Data["blog"] = &blog

	c.Layout = "layouts/front.tpl"
	c.TplName = "single.tpl"
}
