package controllers

import (
    "fmt"
    "ServiceTree/models"
	"github.com/astaxie/beego"
)

// id分配接口
type PageController struct {
	beego.Controller
}

type Menu struct {
    Name          string
    DisplayName   string
    Status        string
}

func (this *PageController) URLMapping() {
	this.Mapping("Get", this.Get)
}

// @Title Get
// @Description 获取一个新自增 id
// @Param	subsys		path 	string	false		"调用系统的接口名称"
// @router /:subsys [get]
func (this *PageController) Get() {
	subsys := this.Ctx.Input.Params[":subsys"]
    menus := [...]Menu{
        {Name:"tag",DisplayName:"Tag 管理",Status:""},
        {Name:"subsys",DisplayName:"系统注册",Status:""},
        {Name:"machine",DisplayName:"机器管理",Status:""},
        {Name:"monitor",DisplayName:"监控系统",Status:""},
    }
    this.TplNames =  "service-tree/menu.tpl"
    switch subsys {
        case "tag":
            menus[0].Status = "active"
            this.Data["Menus"] = menus
            tags := getTags() 
            this.Data["Tags"] = tags
            this.Data["Hierarchy"] = models.DefaultHierarchy()
            this.Layout = "service-tree/tree-tag-manager.html"
        case "init":
            menus[0].Status = "active"
            this.Data["Menus"] = menus
            tags := getTags() 
            this.Data["Tags"] = tags
            this.Data["OptionalKeys"] = models.OptionalKeys()
            this.Data["Hierarchy"] = models.DefaultHierarchy()
            fmt.Println("get optional key list is : " , models.OptionalKeys())
            this.Layout = "service-tree/tree-tag-manager-init.html"
        case "subsys":
            menus[1].Status = "active"
            this.Data["Menus"] = menus
            this.Data["Hierarchy"] = models.DefaultHierarchy()
            this.Layout = "service-tree/tree-register.html"
        case "machine":
            menus[2].Status = "active"
            this.Data["Menus"] = menus
            this.Layout = "service-tree/tree-machine-manager.html"
        case "monitor":
            menus[3].Status = "active"
            this.Data["Menus"] = menus
            this.Layout = "service-tree/tree-monitoring-system.html"
        default:
            menus[0].Status = "active"
            this.Data["Menus"] = menus
            this.Data["Hierarchy"] = models.DefaultHierarchy()
            this.Layout = "service-tree/tree-tag-manager.html"
    }
}

func getTags() []models.Tag {
    var tags []models.Tag
    var tag  models.Tag
    keys := models.MustKeys()     
    for _, key := range keys {
        values := models.GetTagValueByKey(key)
        tag = models.Tag{Key:key,Value:values}
        tags = append(tags, tag)
    }
    return tags
}
