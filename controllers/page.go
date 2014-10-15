package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
)

// id分配接口
type PageController struct {
	beego.Controller
}

func (this *PageController) URLMapping() {
	this.Mapping("Get", this.Get)
}

// @Title Get
// @Description 获取一个新自增 id
// @Param	subsys		path 	string	false		"调用系统的接口名称"
// @router /:subsys [get]
func (this *PageController) Get() {
	subsys := this.Ctx.Input.Params[":subsys"]
    fmt.Println("subsys is:", subsys)
    switch subsys {
        case "subsys":
            this.TplNames = "service-tree/tree-register.html"
        case "machine":
            this.TplNames = "service-tree/tree-machine-manager.html"
        case "tag":
            this.TplNames = "service-tree/tree-tag-manager.html"
        case "monitor":
            this.TplNames = "service-tree/tree-monitoring-system.html"
        default:
            this.TplNames = "service-tree/tree-tag-manager.html"
    }
}
