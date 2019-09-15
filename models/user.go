package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Uid         int64  `orm:"pk"`
	Username    string `orm:"size(32)"`
	Password    string
	Nickname    string
	Role        string
	Ischeck     string
	Photo       string
	Email       string
	Description string
}

func GetUserFromUid(uid int64) User {
	user := User{
		Uid: 1,
	}
	o := orm.NewOrm()
	o.Read(&user)

	fmt.Println(user)
	return user
}
