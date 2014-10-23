package controllers

import (
    "fmt"
	"github.com/astaxie/beego"
    "ServiceTree/models"
    "ServiceTree/libs"
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
// @Param	hierarchy		query 	string	true		"服务树显示层次结构,逗号分隔"
// @Param	resource		query 	string	true		"服务树显示层次结构,逗号分隔"
// @Success 200 {object} models.TreeNode
// @Failure 403 :hierarchy 层次无法构建服务树
// @router /?:hierarchy [get]
func (this *TreeController) Get() {
    var sys_id  int
    var result libs.Result
    var root *models.TreeNode

	hierarchy := this.GetString("hierarchy")
    resource := this.GetString("resource")
    if hierarchy == "" {
        result.Success = false            
        result.Msg = "hierarchy not allow empty"
        this.Data["json"] = result
        this.ServeJson()
    }
    if resource == "" {
        sys_id = 0 
    } else {
        //specify resource, so I try to query sys id
        subsys, err := models.GetSubsysByName(resource)
        if err != nil {
            result.Success = false
            result.Msg = fmt.Sprintf("unable to find resource [%s]", resource)
            this.Data["json"] = result
            this.ServeJson()
        }
        sys_id = subsys.Id
    }

    root = models.BuildTree(hierarchy,sys_id)
    this.Data["json"] = root
	this.ServeJson()
}

