package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Options struct {
	OptionId    int64  `orm:"auto"`
	OptionName  string `orm:"size(128)"`
	OptionValue string `orm:"size(128)"`
}

// AddOptions insert a new Options into database and returns
// last inserted OptionId on success.
func AddOptions(m *Options) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOptionsByOptionId retrieves Options by OptionId. Returns error if
// OptionId doesn't exist
func GetOptionsByOptionId(id int64) (v *Options, err error) {
	o := orm.NewOrm()
	v = &Options{OptionId: id}
	if err = o.QueryTable(new(Options)).Filter("OptionId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOptions retrieves all Options matches certain condition. Returns empty list if
// no records exist
func GetAllOptions() (ml []Options, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Options))
	var l []Options

	if _, err = qs.All(&l, []string{}...); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

// UpdateOptions updates Options by OptionId and returns error if
// the record to be updated doesn't exist
func UpdateOptionsByOptionId(m *Options) (err error) {
	o := orm.NewOrm()
	v := Options{OptionId: m.OptionId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOptions deletes Options by OptionId and returns error if
// the record to be deleted doesn't exist
func DeleteOptions(id int64) (err error) {
	o := orm.NewOrm()
	v := Options{OptionId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Options{OptionId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//从配置列表中找到相应的属性值
func GetOptionFromOptions(options []Options, optionName string) string {
	for _, item := range options {
		if item.OptionName == optionName {
			return item.OptionValue
		}
	}
	return ""
}

// GetOptionsByOptionName retrieves Options by OptionName. Returns error if
// OptionName doesn't exist
func GetOptionsByOptionName(optionName string) (v *Options, err error) {
	o := orm.NewOrm()
	v = &Options{OptionName: optionName}
	if err = o.QueryTable(new(Options)).Filter("OptionName", optionName).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//通过配置名找到对应的配置值
func GetOptionValueByName(optionName string) (OptionValue string) {
	o := orm.NewOrm()
	v := &Options{OptionName: optionName}
	if err := o.QueryTable(new(Options)).Filter("OptionName", optionName).RelatedSel().One(v); err == nil {
		return v.OptionValue
	}
	return ""
}
