package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id int
	Name string 
	Pwd string
	Status int
}

func init (){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "maxscale_route:U5lW0HGg1@tcp(192.168.10.239:4006)/demo?charset=utf8")
	orm.RegisterModel(new (User))
}

func (u *User) TableName() string{
	return "user"
}

func (u *User) Add() bool {
	o := orm.NewOrm()
	o.Using("default")
	user := &User{Name:"Guodawn", Pwd:"123"}
	_, err := o.Insert(user)
	if err != nil { 
		return false
	}
	return true
}