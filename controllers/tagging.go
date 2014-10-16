package controllers

import (
	"ServiceTree/models"
    "ServiceTree/libs"
	"encoding/json"
    "fmt"

	"github.com/astaxie/beego"
)

// 资源打 Tag 接口
type TaggingController struct {
	beego.Controller
}

func (this *TaggingController) URLMapping() {
	this.Mapping("Post", this.Post)
}

// @Title Post
// @Description 对资源打 Tag
// @Param	body		body 	[]models.Tagging	true		"tag内容"
// @Success 200 {int} models.Tagging.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TaggingController) Post() {
	var v []models.Tagging
    var result libs.Result 

	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
    fmt.Printf("start add tagging, with input : %+v\n", v)
	if id := models.AddTagging(v); id != -1 {
        result.Success = true
        result.Msg = "打 Tag成功"
		this.Data["json"] = result
	} else {
        result.Msg = "failed"
        result.Success = false
		this.Data["json"] = result
	}
	this.ServeJson()
}
