package controllers

import (
	"demo/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	c.Layout = "common/layout.tpl"
	c.TplName = "user/login.tpl"
	c.Data["formSubmitUrl"] = c.URLFor("UserController.DoLogin")
	c.Data["homeUrl"] = c.URLFor("HomeController.Index")
}

func (c *UserController) DoLogin() {
	user := new(models.User)

	user.Name = c.GetString("name")
	inputPwd := c.GetString("pwd")
	rt := user.CheckLogin(inputPwd)
	c.Data["json"] = models.Error{Code: rt.Code, Msg: rt.Msg}
	c.ServeJSON()
}

/*func (c *UserController) Add() {
	user := &models.User{Name: "Guodawn", Pwd: "123"}
	rt := user.Add()
	if rt {

	}
	c.ServeJSON()
}*/
