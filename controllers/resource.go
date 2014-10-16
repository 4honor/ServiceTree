package controllers

import (
	"ServiceTree/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// id分配接口
type ResourceController struct {
	beego.Controller
}

func (this *ResourceController) URLMapping() {
	this.Mapping("Post", this.Post)
}

// @Title Post
// @Description 检查资源的有效性
// @Param	body		body 	models.Resource	true		"body for Resource content"
// @Success 200 {int} models.Resource.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ResourceController) Post() {
	var resource models.Resource
    var resources []models.Resource
    resource.SysId = 2
    resource.Name = "web01.qq"
    resource.ResourceId = 1
	json.Unmarshal(this.Ctx.Input.RequestBody, &resource)
    resources = append(resources, resource)
    this.Data["json"] =  resources
	this.ServeJson()
}
