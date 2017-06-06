package routers

import (
	"github.com/yikeso/goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.Router("/deleteCategory",&controllers.DeleteCategoryController{})
}
