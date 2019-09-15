package models

type Comments struct {
	Cid     int64
	Gid     int64
	Pid     int64
	Date    int64
	Poster  string
	Comment string
	Mail    string
	Url     string
	Ip      string
	Hide    string
}

func (u *Comments) TableName() string {
	return "e_comment"
}
