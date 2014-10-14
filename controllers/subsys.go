package controllers

import (
	"github.com/astaxie/beego"
)

type SubsysController struct {
	beego.Controller
}

func (this *SubsysController) Get() {
	this.TplNames = "service-tree/tree-register.html"
}
