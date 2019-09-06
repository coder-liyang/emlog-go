package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "liyangweb/routers"
	"time"
)

func init() {
	//设置日志文件
	//beego.SetLogger("file", `{"filename":"logs/beego.log"}`)
	//设置数据库的配制
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqlpoint := beego.AppConfig.String("mysqlpoint")
	mysqldb := beego.AppConfig.String("mysqldb")
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		mysqluser,
		mysqlpass,
		mysqlurls,
		mysqlpoint,
		mysqldb,
	)
	//orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	err := orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}
}

func convertTime (int_time int64) string {
	return time.Unix(int_time, 0).Format("2006-01-02 15:04:05")
}

func main() {
	beego.AddFuncMap("convertTime", convertTime)
	beego.Run()
}

