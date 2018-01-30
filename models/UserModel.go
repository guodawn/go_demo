package models

import (
	"demo/errnums"

	"demo/utils"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int
	Name   string
	Pwd    string
	Status int8
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "user"
}

/*func (u *User) Add() bool {
	o := orm.NewOrm()
	o.Using("default")
	user := &User{Name: "Guodawn", Pwd: "123"}
	_, err := o.Insert(user)
	if err != nil {
		return false
	}
	return true
}*/

func (u *User) CheckLogin(inputPwd string) (rt Error) {
	o := orm.NewOrm()
	err := o.Read(u, "Name")
	if err == orm.ErrNoRows {
		rt.Code = errnums.UserNotExist
		rt.Msg = "用户不存在"
	} else if err == orm.ErrMissPK {
		rt.Code = errnums.MissingPrimaryKey
		rt.Msg = "找不到主键"
	} else if u.Pwd != utils.String2Md5(inputPwd) {
		rt.Code = errnums.UserPwdErr
		rt.Msg = "密码错误"
	} else {
		rt.Code = errnums.Success
		rt.Msg = "登录成功"
	}
	return
}
