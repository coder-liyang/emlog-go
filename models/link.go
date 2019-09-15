package models

import "github.com/astaxie/beego/orm"

type Link struct {
	Id          int64
	Sitename    string
	Siteurl     string
	Description string
	Hide        string
	Taxis       string
}
//所有友情链接
func GetAllLinks(hide string) ([]Link) {
	//SELECT siteurl,sitename,description FROM e_link WHERE hide='n' ORDER BY taxis ASC
	o := orm.NewOrm()
	qs := o.QueryTable(new(Link))
	if hide != "" {
		qs = qs.Filter("hide", hide)
	}
	qs = qs.OrderBy("taxis")
	var l []Link
	_, _ = qs.All(&l)
	return l
}
