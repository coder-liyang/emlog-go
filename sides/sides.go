package sides

import (
	"github.com/astaxie/beego/orm"
	"liyangweb/models"
	"strconv"
	"time"
)
//汇总所有侧边栏
var Sides = make(map[string][]interface{})

//微语
func WidgetTwitter() (twitters []interface{}) {
	//SELECT option_value FROM e_options where option_name='index_newtwnum' 要取多少条
	indexNewtwnum, _ := models.GetOptionsByOptionName("index_newtwnum") //要取多少条
	offset, _ := strconv.ParseInt(indexNewtwnum.OptionValue, 10, 64)
	twitters,_ = models.GetTwitters(1, offset)
	//fmt.Printf("%T", twitters)
	Sides["twitter"] = twitters
	return
}

//归档
type Archive struct {
	Record string
	Date string
	Lognum int64
}
func WidgetArchive() (archive []interface{}) {
	var (
		l []models.Blog
		dateMonth string
		//num int64 = 0
		)

	/*
	[0]=>
	  array(3) {
	    ["record"]=>
	    string(11) "2019年8月"
	    ["date"]=>
	    string(6) "201908"
	    ["lognum"]=>
	    int(1)
	  }
	 */
	//select date from ' . DB_PREFIX . "blog WHERE hide='n' and checked='y' and type='blog' ORDER BY date DESC
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Blog))
	qs = qs.Filter("hide", "n")
	qs = qs.Filter("checked", "y")
	qs = qs.Filter("type", "blog")
	qs = qs.OrderBy([]string{"date"}...)
	qs = qs.RelatedSel()
	qs.All(&l, []string{}...)

	for _, blog := range l {
		//fmt.Println(blog)
		dateMonth = time.Unix(blog.Date, 0).Format("2006-01")
		archive = append(archive, Archive{
			Record:dateMonth,
			Date:dateMonth,
			Lognum:1,
		})
	}
	Sides["archive"] = archive
	return
}