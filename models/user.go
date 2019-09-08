package models

type User struct {
	Uid int64 `orm:"pk"`
	Username string `orm:"size(32)"`
	Password string
	Nickname string
	Role string
	Ischeck string
	Photo string
	Email string
	Description string
}
