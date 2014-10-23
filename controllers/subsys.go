package controllers

import (
	"fmt"
	"ServiceTree/models"
    "ServiceTree/libs"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 第三方系统接口
type SubsysController struct {
	beego.Controller
}

func (this *SubsysController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.SubsysList)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
	this.Mapping("Get", this.Get)
    this.Mapping("Hierarchy", this.Hierarchy)
}

// @Title Post
// @Description create Subsys
// @Param	body		body 	models.Subsys	true		"body for Subsys content"
// @Success 200 {int} models.Subsys.Id
// @Failure 403 body is empty
// @router / [post]
func (this *SubsysController) Post() {
	var v models.Subsys
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if id, err := models.AddSubsys(&v); err == nil {
		this.Data["json"] = map[string]int64{"id": id}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Get
// @Description get Subsys by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Subsys
// @Failure 403 :id is empty
// @router /:id [get]
func (this *SubsysController) GetOne() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSubsysById(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = v
	}
	this.ServeJson()
}

// @Title Get All
// @Description get Subsys
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Subsys
// @Failure 403
// @router / [get]
func (this *SubsysController) GetAll() {
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

	l, err := models.GetAllSubsys(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJson()
}

// @Title Update
// @Description update the Subsys
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Subsys	true		"body for Subsys content"
// @Success 200 {object} models.Subsys
// @Failure 403 :id is not int
// @router /:id [put]
func (this *SubsysController) Put() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Subsys{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdateSubsysById(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Update
// @Description 更新系统层次接口
// @Param	name		query 	string	true		"系统名称"
// @Param	hierarchy		query 	models.Subsys	true		"更新层次结构"
// @Success 200 {object} models.Subsys
// @Failure 403 :id is not int
// @router /hierarchy [get]
func (this *SubsysController) Hierarchy() {
    var result libs.Result
	name := this.GetString("name")
    hierarchy := this.GetString("hierarchy")
    if name == "" || hierarchy == "" {
        result.Success = false 
        result.Msg = "sys name and hierarchy not allow empty"
        this.Data["json"] = result
        this.ServeJson()
    }
    v := models.Subsys{Name: name, Hierarchy: hierarchy}
	if err := models.UpdateHierarchyByName(&v); err == nil {
        result.Success = true
        result.Msg = "update success"
	} else {
        result.Success = false 
        result.Msg = err.Error()
	}
	this.Data["json"] = result
	this.ServeJson()
}

// @Title Delete
// @Description delete the Subsys
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *SubsysController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSubsys(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}


// @Title Get All
// @Description get Subsys
// @Success 200 {object} models.Subsys
// @Failure 403
// @router / [get]
func (this *SubsysController) SubsysList() {
    subsysList := models.SubsysList()
    for index, sys := range subsysList {
        state := sys["State"] 
        switch state {
            case "0": 
                subsysList[index]["State"] = "审核中"
            case "1":
                subsysList[index]["State"] = "审核通过"
            default:
                subsysList[index]["State"] = "审核中"
       }
    }
    fmt.Println("len of subsysList is : ", len(subsysList))
    if subsysList == nil || len(subsysList) == 0 {
        this.Data["json"]  = make([]int, 0)
    }else{
    	this.Data["json"]   = subsysList
    }
    this.ServeJson()
}
