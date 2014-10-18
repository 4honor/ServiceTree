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
}

// @Title Get
// @Description 获取服务树
// @Param	hierarchy		path 	string	false		"服务树显示层次结构,逗号分隔"
// @Success 200 {object} models.TreeNode
// @Failure 403 :hierarchy 层次无法构建服务树
// @router /?:hierarchy [get]
func (this *TreeController) Get() {
	hierarchy := this.Ctx.Input.Params[":hierarchy"]
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
    depart1.IconSkin = "pIcon01"
    depart1.IsParent = true

    var module1 = models.TreeNode{}
    module1.Name = "api"
    module1.Meta = "module"
    module1.IsParent = false
    module1.IconSkin = "icon01"

    var module2 = models.TreeNode{}
    module2.Name = "pay"
    module2.Meta = "module"
    module2.IsParent = false
    module2.IconSkin = "icon01"

    var depart2 = models.TreeNode{}
    depart2.Name = "dache"
    depart2.Meta = "dept"
    depart2.IconSkin = "pIcon01"
    depart2.IsParent = true

    var module3 = models.TreeNode{}
    module3.Name = "api"
    module3.Meta = "module"
    module3.IsParent = false
    module3.IconSkin = "icon01"

    var module4 = models.TreeNode{}
    module4.Name =  hierarchy
    module4.Meta = "module"
    module4.IsParent = false
    module4.IconSkin = "icon01"

    var depart3 = models.TreeNode{}
    depart3.Name = "testing"
    depart3.Meta = "dept"
    depart3.IconSkin = "icon01"
    depart3.IsParent = false

    depart1.Children = append(depart1.Children, &module1)
    depart1.Children = append(depart1.Children, &module2)
    depart2.Children = append(depart2.Children, &module3)
    depart2.Children = append(depart2.Children, &module4)
    corp.Children = append(corp.Children, &depart1)
    corp.Children = append(corp.Children, &depart2)
    corp.Children = append(corp.Children, &depart3)
    root.Children = append(root.Children, &corp)

    fmt.Printf("get hierarchy: %s, tree:%+v", hierarchy,root)
    this.Data["json"] = root
	this.ServeJson()
}

