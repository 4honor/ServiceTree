package controllers

import (
	"ServiceTree/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// id分配接口
type IdallocController struct {
	beego.Controller
}

func (this *IdallocController) URLMapping() {
	this.Mapping("Get", this.Alloc)
}

// @Title Post
// @Description create Idalloc
// @Param	body		body 	models.Idalloc	true		"body for Idalloc content"
// @Success 200 {int} models.Idalloc.Id
// @Failure 403 body is empty
// @router / [post]
func (this *IdallocController) Post() {
	var v models.Idalloc
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if id, err := models.AddIdalloc(&v); err == nil {
		this.Data["json"] = map[string]int64{"id": id}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Get
// @Description 获取一个新自增 id
// @Param	api		path 	string	false		"调用系统的接口名称"
// @Success 200 {object} models.Idalloc
// @Failure 403 :api is empty
// @router /:api [get]
func (this *IdallocController) Alloc() {
	api := this.Ctx.Input.Params[":api"]
    var v models.Idalloc
    v.Api = api
    if id, err := models.AddIdalloc(&v); err == nil {
        this.Data["json"]  = map[string]int64{"id":id}
    } else {
        this.Data["json"] = err.Error()
    }
	this.ServeJson()
}
