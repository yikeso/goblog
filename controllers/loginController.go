package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get()  {
	c.TplName = "log.html"
}

func (c *LoginController) Post()  {
	data := c.Input()
	maxAge := 1<<10
	c.Ctx.SetCookie("username",data.Get("username"),maxAge,"/")
	c.Redirect("/",301)
}
