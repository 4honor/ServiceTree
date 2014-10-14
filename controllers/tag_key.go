package controllers 

import (
	"github.com/astaxie/beego"
)

type TagKeyController struct {
	beego.Controller
}

func (this *TagKeyController) Get() {
    this.TplNames = "tree-tag-manager.html"
}
