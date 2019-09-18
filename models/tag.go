package models

type Tag struct {
	Tid int64 `orm:"auto;pk"`
	Tagname string
	Gid string
}