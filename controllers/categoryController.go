package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yikeso/goblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	checkLogin(&c.Controller)
	c.Data["IsCategory"] = true
	categories,err := models.FindAllCategory()
	if err == nil{
		c.Data["Categories"] = categories
	}
	c.TplName = "category.html"
}

func (c *CategoryController) Post() {
	name := c.Input().Get("name")
	err := models.AddCategory(name)
	if err != nil{
		c.Ctx.WriteString("类型名重复，插入失败")
	}else {
		checkLogin(&c.Controller)
		c.Redirect("/category",301)
	}
}

type DeleteCategoryController struct {
	beego.Controller
}

func (c *DeleteCategoryController) Get() {
	checkLogin(&c.Controller)
	c.Data["IsCategory"] = true
	err := models.DeleteById(c.Input().Get("id"))
	if err != nil {
		beego.Error(err.Error())
	}
	categories,err := models.FindAllCategory()
	if err == nil{
		c.Data["Categories"] = categories
	}
	c.TplName = "category.html"
}