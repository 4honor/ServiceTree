package controllers

import (
    "fmt"
	"ServiceTree/models"
//	"encoding/json"
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
    var resources []models.Resource = []models.Resource{}
	//json.Unmarshal(this.Ctx.Input.RequestBody, &resources)
    fmt.Println(resources)
    this.Data["json"] =  resources
	this.ServeJson()
}
