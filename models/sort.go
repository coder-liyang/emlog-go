package models

import "github.com/astaxie/beego/orm"

type Sort struct {
	Sid         int64 `orm:"auto;pk"`
	Sortname    string
	Alias       string
	Taxis       int64
	Pid         int64
	Description string
	Template    string
	ChildSort	[]*Sort `orm:"-"`
}

var sortWithLevel []Sort
func GetAllSortWithLevel(pid int64) (sorts []*Sort) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sort))
	qs = qs.Filter("pid", pid)
	num, _ := qs.All(&sortWithLevel)
	if num > 0 {
		for _, sort := range sortWithLevel {
			sort.ChildSort = GetAllSortWithLevel(sort.Sid)
		}
	}
	return sorts
}