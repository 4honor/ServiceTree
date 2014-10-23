package controllers

import (
	"ServiceTree/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 机器管理接口
type MachineController struct {
	beego.Controller
}

func (this *MachineController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Get
// @Description 通过机器 id 来获取机器信息
// @Param      id              path    string  true            "查询机器 id"
// @Success 200 {object} models.Machine
// @Failure 403 :id is empty
// @router /:id [get]
func (this *MachineController) GetOne() {
       idStr := this.Ctx.Input.Params[":id"]
       id, _ := strconv.Atoi(idStr)
       v, err := models.GetMachineById(id)
       if err != nil {
               this.Data["json"] = err.Error()
       } else {
               this.Data["json"] = v
       }
       this.ServeJson()
}

// @Title Post
// @Description 新增机器
// @Param	body		body 	models.Machine	true		"新增机器内容"
// @Success 200 {int} models.Machine.Id
// @Failure 403 提交数据为空
// @Failure 500 后端处理错误
// @router / [post]
func (this *MachineController) Post() {
	var v models.Machine
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if id, err := models.AddMachine(&v); err == nil {
		this.Data["json"] = map[string]int64{"id": id}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Get All
// @Description 批量获取机器信息
// @Param	ns		path 	string	true		"查询机器的命名空间(ns), eg: corp:xiaoju,depart:dache"
// @Success 200 {object} models.Machine
// @Failure 403
// @router /:ns:string [get]
func (this *MachineController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := this.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := this.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				this.Data["json"] = errors.New("Error: invalid query key/value pair")
				this.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllMachine(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJson()
}

// @Title Update
// @Description 更新某台机器信息
// @Param	id		path 	string	true		"想要更新机器 id"
// @Param	body		body 	models.Machine	true		"body for Machine content"
// @Success 200 {object} models.Machine
// @Failure 403 :id is not int
// @router /:id [put]
func (this *MachineController) Put() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Machine{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdateMachineById(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Delete
// @Description 删除某台机器
// @Param	id		path 	string	true		"删除机器 id"
// @Success 200 {string} 删除成功!
// @Failure 403 机器id 不存在
// @router /:id [delete]
func (this *MachineController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMachine(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}
