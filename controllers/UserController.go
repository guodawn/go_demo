package controllers

import (
	"demo/models"
)

type UserController struct {
	BaseController
}

type Ret struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
}



func (c *UserController) Login() {
	c.Layout = "common/layout.tpl"
	c.TplName = "user/login.tpl"
	c.Data["formSubmitUrl"] = c.URLFor("UserController.DoLogin")
	c.Data["homeUrl"] = c.URLFor("HomeController.Index")
}

func (c *UserController) DoLogin() {
	c.Data["json"] = Ret{Code : 0, Msg : "登录成功"}
	c.ServeJSON()
}

func (c *UserController) Add() {
	user := &models.User{Name:"Guodawn", Pwd:"123"}
	rt := user.Add()
	if rt {

	}
	c.ServeJSON()
}