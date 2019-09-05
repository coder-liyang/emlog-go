package controllers

import "github.com/astaxie/beego"

type MeController struct {
	beego.Controller
}

func (c *MeController) About() {
	c.TplName = "about.tpl"
}