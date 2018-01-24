package routers

import (
	"demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.Router("/doLogin", &controllers.UserController{}, "POST:DoLogin")
	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
}
