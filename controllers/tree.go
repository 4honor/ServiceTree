package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
    "ServiceTree/models"
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
    var root = models.TreeNode{}

    root.Name = "Root"
    root.Meta = "Root"
    root.IconSkin = "pIcon01"
    root.IsParent = true

    var corp = models.TreeNode{}
    corp.Name = "didi"
    corp.Meta = "corp"
    corp.IconSkin = "pIcon01"
    corp.IsParent = true

    var depart1 = models.TreeNode{}
    depart1.Name = "gs"
    depart1.Meta = "dept"
    depart1.IconSkin = "pIcon1"
    depart1.IsParent = false

    var depart2 = models.TreeNode{}
    depart2.Name = "dache"
    depart2.Meta = "dept"
    depart2.IsParent = false

    corp.Children = append(corp.Children, &depart1)
    corp.Children = append(corp.Children, &depart2)
    root.Children = append(root.Children, &corp)

	hierarchy := this.Ctx.Input.Params[":hierarchy"]
    fmt.Printf("get hierarchy: %s, tree:%+v", hierarchy,root)
    this.Data["json"] = root
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
