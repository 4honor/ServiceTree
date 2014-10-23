package controllers

import (
    "ServiceTree/models"
	"github.com/astaxie/beego"
)

// Ns 获取 node space 下对应的资源
type NsController struct {
	beego.Controller
}

func (this *NsController) URLMapping() {
	this.Mapping("Get", this.Get)
}

// @Title Get
// @Description odin兼容接口, 查询 tag kv
// @Param	resource		query 	string	true		"资源名字"
// @Param	subsys		query 	string	true		"子系统"
// @router / [get]
func (this *NsController) Get() {
    var ns []models.Ns
    var subsys *models.Subsys
    var resource models.Resource

	resource_name := this.GetString("resource")
    sys_name := this.GetString("subsys")
    if resource_name == "" || sys_name == "" {
        beego.Warn("query ns with empty resource and subsys ")
        ns = make([]models.Ns, 0)
    }
    subsys, err := models.GetSubsysByName(sys_name) 
    if err == nil {
        resource.SysId = subsys.Id
        //get resource id
        if subsys.Name == "machine"  {
            var resource_model models.Machine
            resource_model, err = models.GetMachineByName(resource_name)              
            if err == nil {
                resource.ResourceId = resource_model.Id 
            }
        }
        if subsys.Name == "monitor" {
            var resource_model models.Monitor
            beego.Trace("get subsys monitor name:", resource_name)
            resource_model, err  = models.GetMonitorByName(resource_name)
            beego.Debug("get resource model:", resource_model)
            if err == nil {
                resource.ResourceId = resource_model.Id 
            }else{
                beego.Warn("get resource failed, with err:", err)
            }
        }
        //get ns with resource  
        ns = models.GetNsWithResource(resource)    
    } else {
        beego.Warn("query ns without any result, with subsys:",sys_name)
        ns = make([]models.Ns,0)
    }
    
    this.Data["json"] = ns
    this.ServeJson()
}
