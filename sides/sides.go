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
	Date  string
	Total int64
}
func WidgetArchive() (archive []interface{}) {
	var archives []Archive
	o := orm.NewOrm()
	_, _ = o.Raw("SELECT from_unixtime(date, '%Y年%m月') as `date`, count(*) as total " +
		"FROM e_blog " +
		"WHERE hide = 'n' and checked = 'y' and type = 'blog' " +
		"GROUP BY `date` " +
		"ORDER BY `date` DESC").QueryRows(&archives)

	var archivesInterface []interface{}
	for _, v := range archives {
		archivesInterface = append(archivesInterface, v)
		//如果接收item的话,也可以用下面的方法
		//archivesInterface[i] = v
	}
	Sides["archive"] = archivesInterface
	return archivesInterface
}
