package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Uid int64 `orm:"auto"`
	Username string `orm:"size(32)"`
	Password string
	Nickname string
	Role string
	Ischeck string
	Photo string
	Email string
	Description string
}

func init()  {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(User))
}