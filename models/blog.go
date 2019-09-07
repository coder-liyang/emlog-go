package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Blog struct {
	//Gid          int64 `orm:"auto"`
	//Gid         int64
	Gid          int64 `orm:"auto;pk"`
	Title       string `orm:"size(128)"`
	Date        int64
	Content     string `orm:"size(128)"`
	Excerpt     string `orm:"size(128)"`
	Alias       string `orm:"size(128)"`
	Sortid      int
	Type        string `orm:"size(128)"`
	Views       int
	Comnum      int
	Attnum      int
	Top         string `orm:"size(128)"`
	Sortop      string `orm:"size(128)"`
	Hide        string `orm:"size(128)"`
	Checked     string `orm:"size(128)"`
	AllowRemark string `orm:"size(128)"`
	Password    string `orm:"size(128)"`
	Template    string `orm:"size(128)"`
	//Author     *User `json:"user_id";orm:"rel(fk)"`
	User        *User  `orm:"rel(fk);column(author)" json:"author"`
}

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(Blog))
}

// AddBlog insert a new Blog into database and returns
// last inserted Gid on success.
func AddBlog(m *Blog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBlogByGid retrieves Blog by Gid. Returns error if
// Gid doesn't exist
func GetBlogByGid(id int64) (v *Blog, err error) {
	o := orm.NewOrm()
	v = &Blog{Gid: id}
	if err = o.QueryTable(new(Blog)).Filter("Gid", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBlog retrieves all Blog matches certain condition. Returns empty list if
// no records exist
func GetAllBlog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Blog))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Blog
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBlog updates Blog by Gid and returns error if
// the record to be updated doesn't exist
func UpdateBlogByGid(m *Blog) (err error) {
	o := orm.NewOrm()
	v := Blog{Gid: m.Gid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBlog deletes Blog by Gid and returns error if
// the record to be deleted doesn't exist
func DeleteBlog(id int64) (err error) {
	o := orm.NewOrm()
	v := Blog{Gid: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Blog{Gid: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
//获取前台显示的文章
func GetLogsForHome(page int64) (ml []interface{}, err error) {
	var query = make(map[string]string)
	query["hide"] = "n"
	query["checked"] = "y"

	var limit int64 = 10
	offset := (page - 1) * limit
	return GetAllBlog(
		query,//make(map[string]string)
		[]string{}, //var fields []string
		[]string{"top", "date"},
		[]string{"desc", "desc"},
		offset,
		10,
		)
}

func BlogTotal() (int64, error) {
	var query = make(map[string]string)
	query["hide"] = "n"
	query["checked"] = "y"

	o := orm.NewOrm()
	qs := o.QueryTable(new(Blog))
	return qs.Count()
}