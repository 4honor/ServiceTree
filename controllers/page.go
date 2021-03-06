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
// @Description 构建显示页面
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
            this.Data["Hierarchy"] = getSysHierarchy("machine")
            this.Data["Resource"] = "machine"
            this.Layout = "service-tree/tree-machine-manager.html"
        case "monitor":
            menus[3].Status = "active"
            this.Data["Menus"] = menus
            this.Data["Hierarchy"] = getSysHierarchy("monitor")
            beego.Debug("get monitor hierarchy: ", this.Data["Hierarchy"])
            this.Data["Resource"] = "monitor"
            this.Layout = "service-tree/tree-monitoring-system.html"
        case "login":
            error_msg := this.GetString("error")
            if error_msg == "block" {
                this.Data["Error"] = error_msg
            }else{
                this.Data["Error"] = ""
            }
            this.TplNames = "service-tree/login.html"
        default:
            menus[0].Status = "active"
            this.Data["Menus"] = menus
            this.Data["Hierarchy"] = models.DefaultHierarchy()
            this.Layout = "service-tree/tree-tag-manager.html"
    }
}

func getTags() []models.TagSchema {
    var tags []models.TagSchema
    var tag  models.TagSchema
    keys := models.MustKeys()     
    for _, key := range keys {
        values := models.GetTagValueByKey(key)
        tag = models.TagSchema{Key:key,Value:values}
        tags = append(tags, tag)
    }
    return tags
}

func getSysHierarchy(sys string) string {
    hierarchy := models.GetHierarchy(sys)    
    //empty hierarchy downgrade to default hierarchy
    if hierarchy == "" {
        hierarchy = models.DefaultHierarchy()
    }
    return hierarchy
}
