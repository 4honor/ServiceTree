package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
    menus := [...]Menu{
        {Name:"tag",DisplayName:"Tag 管理",Status:""},
        {Name:"subsys",DisplayName:"系统注册",Status:""},
        {Name:"machine",DisplayName:"机器管理",Status:""},
        {Name:"monitor",DisplayName:"监控系统",Status:""},
    }
    menus[0].Status = "active"
    this.Data["Menus"] = menus
    this.TplNames =  "service-tree/menu.tpl"
    this.Layout = "service-tree/tree-tag-manager.html"
}
