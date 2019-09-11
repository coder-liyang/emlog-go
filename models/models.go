package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.RegisterModelWithPrefix(
		beego.AppConfig.String("mysqldbprefix"),
		new(User),
		new(Blog),
		new(Options),
		new(Sort),
		new(Twitter),
		)
}
