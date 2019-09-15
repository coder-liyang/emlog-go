package sides

import (
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
//原系统还查询了父评论,这里我就简单处理了,不查询父评论了
func WidgetNewcomm() (commentsInterface []interface{}) {
	var (
		o orm.Ormer
		//[{0 comment_order newer} {0 comment_paging y} {0 comment_pnum 15} {0 comment_subnum 20} {0 index_comnum 10}]
		qs orm.QuerySeter
		l []models.Options
		comments []models.Comments
	)
	/*
	查看配置信息
	SELECT option_value, option_name
	FROM e_options
	WHERE option_name IN ('index_comnum', 'comment_subnum', 'comment_paging', 'comment_pnum', 'comment_order')
	comment_pnum:每页显示评论条数
	index_comnum:首页最新评论数
	comment_subnum:新近评论截取字节数
	comment_paging:评论是否分页
	comment_order:排序方式 newer older
	 */
	o = orm.NewOrm()
	qs = o.QueryTable(new(models.Options))
	qs = qs.Filter("option_name__in", "index_comnum", "comment_subnum", "comment_paging", "comment_pnum", "comment_order")
	_, _ = qs.All(&l, "option_value", "option_name")

	//map[comment_order:newer comment_paging:y comment_pnum:15 comment_subnum:20 index_comnum:10]
	//var optionsMap = make(map[string]string)
	optionsMap := map[string]string{}
	for _, v := range l {
		//fmt.Printf("OptionName:%s-%T;OptioinValue:%s:%T\n", v.OptionName,v.OptionName, v.OptionValue, v.OptionValue)
		optionsMap[v.OptionName] = v.OptionValue
	}
	//查评论
	//SELECT * FROM e_comment WHERE hide='n' ORDER BY date DESC LIMIT 0, 10
	//qs = o.QueryTable(new(models.Comments))
	//qs = o.QueryTable("comment")
	//qs = o.QueryTable("e_comment")
	//qs = qs.Filter("hide", "n")
	//
	//qs.All(comments)
	//for _, comment := range comments{
	//	commentsInterface = append(commentsInterface, comment)
	//}
	oo := orm.NewOrm()
	_, _ = oo.Raw("SELECT * FROM e_comment WHERE hide='n' ORDER BY date DESC LIMIT 0, ?", optionsMap["index_comnum"]).QueryRows(&comments)
	for _, comment := range comments{
		commentsInterface = append(commentsInterface, comment)
	}
	return commentsInterface
}
