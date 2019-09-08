package models

type Sort struct {
	Sid         int64 `orm:"auto;pk"`
	Sortname    string
	Alias       string
	Taxis       int64
	Pid         int64
	Description string
	Template    string
}
