package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	checkLogin(&c.Controller)
	c.TplName = "index.html"
}

func checkLogin(c *beego.Controller){
	username := c.Ctx.GetCookie("username")
	if strings.EqualFold("",username) {
		c.Data["IsLogOut"] = true
	}else {
		c.Data["IsLogOut"] = false
		c.Data["Username"] =  username
	}
}