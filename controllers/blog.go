package controllers

import (
	"github.com/astaxie/beego"
	"liyangweb/models"
)

type BlogController struct {
	beego.Controller
}
func (c *BlogController) GetOneBlog(gid int64)  {
	models.GetBlogByGid(gid)
}