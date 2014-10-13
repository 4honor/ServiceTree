package controllers

import (
	"ServiceTree/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// TagMeta 操作接口
type TagMetaController struct {
	beego.Controller
}

func (this *TagMetaController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description 创建tag key
// @Param	body		body 	models.TagMeta	true		"tag key信息"
// @Success 200 {int} models.TagMeta.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TagMetaController) Post() {
	var v models.TagMeta
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if id, err := models.AddTagMeta(&v); err == nil {
		this.Data["json"] = map[string]int64{"id": id}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Get
// @Description get TagMeta by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TagMeta
// @Failure 403 :id is empty
// @router /:id [get]
func (this *TagMetaController) GetOne() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTagMetaById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJson()
}

// @Title Get All
// @Description 获取 tag key 列表
// @Param	offset	query	string	false	"当前第几页,默认从0开始计数"
// @Param	limit	query	string	false	"获取 tag key 的个数, 默认10个"
// @Success 200 {object} models.TagMeta
// @Failure 403
// @router / [get]
func (this *TagMetaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// limit: 10 (default is 10)
	if v, err := this.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt("offset"); err == nil {
		offset = v
	}
    sortby = append(sortby, "TagKey")
    order  = append(order, "asc")

	l, err := models.GetAllTagMeta(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJson()
}

// @Title Update
// @Description 更新 tag key 信息
// @Param	id		path 	string	true		"需要更新 tag key 的id"
// @Param	body		body 	models.TagMeta	true		"tag key 更新信息"
// @Success 200 {object} models.TagMeta
// @Failure 403 :id is not int
// @router /:id [put]
func (this *TagMetaController) Put() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.TagMeta{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdateTagMetaById(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Delete
// @Description 删除 tag key 信息
// @Param	id		path 	string	true		"需要删除 tag key 的id"
// @Success 200 {string} 删除 tag key 成功
// @Failure 403 该 tag key 的id 不存在
// @router /:id [delete]
func (this *TagMetaController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTagMeta(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}
