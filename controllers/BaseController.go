package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare(){
	c.CheckLogin()
}

func (c *BaseController ) CheckLogin() {
	if false {
		loginUrl := c.URLFor("UserController.Login")
		c.Redirect(loginUrl, 302)
		c.StopRun()
	}
} 