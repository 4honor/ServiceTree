package controllers

import (
    "fmt"
	"ServiceTree/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// id分配接口
type TagKeyController struct {
	beego.Controller
}

func (this *TagKeyController) URLMapping() {
	this.Mapping("Get", this.Alloc)
}

// @Title Post
// @Description create Idalloc
// @Param	body		body 	models.Idalloc	true		"body for Idalloc content"
// @Success 200 {int} models.Idalloc.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TagKeyController) Post() {
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
// @Description 根据 tag key 来获取对应的 tag value
// @Param	tagk		path 	string	false		"调用系统的接口名称"
// @Success 200 {object} models.Idalloc
// @Failure 403 :tagk is empty
// @router /:tagk [get]
func (this *TagKeyController) Alloc() {
    var tag_value []string
	tagk := this.Ctx.Input.Params[":tagk"]
    tag_value = models.MustKeys()
    this.Data["json"]  = tag_value
    fmt.Printf("tag key:%s tag_value: %v", tagk, tag_value)
	this.ServeJson()
}


//@Title Get All
//@Description get all TagKey
//@Successs 200 {object} models.TagValue
