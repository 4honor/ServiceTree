package controllers

import (
	"github.com/astaxie/beego"
)

// 服务树请求接口
type TreeController struct {
	beego.Controller
}

func (this *TreeController) URLMapping() {
	this.Mapping("Get", this.Get)
    this.Mapping("Ns", this.Ns)
}

// @Title Get
// @Description 获取服务树
// @Param	hierarchy		path 	string	false		"服务树显示层次结构,逗号分隔"
// @Success 200 {object} models.Tree
// @Failure 403 :hierarchy 层次无法构建服务树
// @router /?:hierarchy [get]
func (this *TreeController) Get() {
	hierarchy := this.Ctx.Input.Params[":hierarchy"]
    this.Data["json"] = hierarchy
	this.ServeJson()
}

// @Title Get
// @Description 获取服务树节点下的
// @Param	ns		path 	string	false		"资源所属空间"
// @Success 200 {object} models.Tree
// @Failure 403 :hierarchy 层次无法构建服务树
// @router /ns/:ns [get]
func (this *TreeController) Ns() {
	hierarchy := this.Ctx.Input.Params[":ns"]
    this.Data["json"] = hierarchy
	this.ServeJson()
}
