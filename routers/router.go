package routers

import (
	"original/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.IndexController{},"*:Index")
	beego.Router("/login", &controllers.LoginController{},"*:LoginIndex")
    beego.Router("/user",&controllers.UserController{},"Post:Insert")
	//定制错误页
	beego.ErrorController(&controllers.ErrorController{})

}
