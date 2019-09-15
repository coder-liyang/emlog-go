package sides

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"liyangweb/models"
	"strconv"
)

//汇总所有侧边栏
var Sides = make(map[string][]interface{})

//微语
func WidgetTwitter() (twitters []interface{}) {
	//SELECT option_value FROM e_options where option_name='index_newtwnum' 要取多少条
	indexNewtwnum, _ := models.GetOptionsByOptionName("index_newtwnum") //要取多少条
	offset, _ := strconv.ParseInt(indexNewtwnum.OptionValue, 10, 64)
	twitters, _ = models.GetTwitters(1, offset)
	//fmt.Printf("%T", twitters)
	Sides["twitter"] = twitters
	return
}

//归档
type Archive struct {
	Ym  string
	Total int64
}
func WidgetArchive() (archive []interface{}) {
	var archives []Archive
	o := orm.NewOrm()
	_, _ = o.Raw("SELECT from_unixtime(`date`, '%Y年%m月') as `ym`, count(*) as total " +
		"FROM e_blog " +
		"WHERE hide = 'n' and checked = 'y' and type = 'blog' " +
		"GROUP BY `ym` " +
		"ORDER BY `ym` DESC").QueryRows(&archives)

	var archivesInterface []interface{}
	for _, v := range archives {
		archivesInterface = append(archivesInterface, v)
		//如果接收item的话,也可以用下面的方法
		//archivesInterface[i] = v
	}
	Sides["archive"] = archivesInterface
	return archivesInterface
}
//友情链接
func WidgetLink() ([]interface{}) {
	//SELECT siteurl,sitename,description FROM e_link WHERE hide='n' ORDER BY taxis ASC
	var linkInterafce []interface{}
	links := models.GetAllLinks("n")
	for _, v := range links {
		linkInterafce = append(linkInterafce, v)
	}
	return linkInterafce
}
//我的信息
func WidgetBlogger() ([]interface{}) {
	user := models.GetUserFromUid(1)
	var usersI = make([]interface{}, 1)
	usersI[0] = user
	return usersI
}
//最新评论
func WidgetNewcomm() (comments []interface{}) {
	/*
	SELECT option_value, option_name
	FROM e_options
	WHERE option_name IN ('index_comnum', 'comment_subnum', 'comment_paging', 'comment_pnum', 'comment_order')
	comment_pnum:每页显示评论条数
	index_comnum:首页最新评论数
	comment_subnum:新近评论截取字节数
	comment_paging:评论是否分页
	comment_order:排序方式 newer older
	 */
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Options))
	qs = qs.Filter("option_name__in", "index_comnum", "comment_subnum", "comment_paging", "comment_pnum", "comment_order")
	var l []models.Options
	qs.All(&l, "option_value", "option_name")
	fmt.Println(l)
	return comments
}
