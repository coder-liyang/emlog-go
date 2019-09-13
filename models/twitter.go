package models

import (
	"github.com/astaxie/beego/orm"
)

type Twitter struct {
	Id       int64
	Content  string
	Img      string
	Author   int64
	Date     int64
	Replynum int64
}

func GetTwitters(p int64, offset int64) (ml []interface{}, err error) {
	if p <= 0 {
		p = 1
	}
	o := orm.NewOrm()
	var l []Twitter
	qs := o.QueryTable(new(Twitter))
	qs = qs.OrderBy([]string{"-id"}...)
	if _, err = qs.Limit(offset, (p-1)*offset).All(&l, []string{}...); err == nil {
		for _, v := range l {

			//m := make(map[string]interface{})
			//val := reflect.ValueOf(v)
			//for _, fname := range fields {
			//	m[fname] = val.FieldByName(fname).Interface()
			//}
			//ml = append(ml, m)

			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}
