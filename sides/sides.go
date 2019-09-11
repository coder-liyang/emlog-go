package sides

import (
	"liyangweb/models"
	"strconv"
)
//汇总所有侧边栏
var Sides = make(map[string][]interface{})

func Twitter() (twitters []interface{}) {
	//SELECT option_value FROM e_options where option_name='index_newtwnum' 要取多少条
	indexNewtwnum, _ := models.GetOptionsByOptionName("index_newtwnum") //要取多少条
	pageSize, _ := strconv.ParseInt(indexNewtwnum.OptionValue, 10, 64)
	twitters,_ = models.GetTwitters(1, pageSize)
	//fmt.Printf("%T", twitters)
	Sides["twitter"] = twitters
	return
}