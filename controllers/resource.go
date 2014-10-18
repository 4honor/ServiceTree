package controllers

import (
    "fmt"
    "strings"
	"ServiceTree/models"
    "ServiceTree/libs"
	"encoding/json"
	"github.com/astaxie/beego"
)

// id分配接口
type ResourceController struct {
	beego.Controller
}

func (this *ResourceController) URLMapping() {
	this.Mapping("Post", this.Post)
    this.Mapping("Get",  this.Ns)
}

// @Title Post
// @Description 检查资源的有效性
// @Param	body		body 	models.Resource	true		"body for Resource content"
// @Success 200 {int} models.Resource.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ResourceController) Post() {
    var resources []models.Resource = []models.Resource{}

	json.Unmarshal(this.Ctx.Input.RequestBody, &resources)
    fmt.Printf("input is :[%+v]\n", resources)
    this.Data["json"] = models.FilterResources(resources)
	this.ServeJson()
}

// @Title Get
// @Description 获取某个命名空间下的对应的资源
// @Param	ns		path 	string	false		"资源所属空间"
// @Success 200 {object} models.Resource
// @Failure 403 :ns 层次无法构建服务树
// @router /?:ns
func (this *ResourceController) Ns() {
    var ns string
    var result libs.Result
    var tags map[string]string
    var resources []interface{}

	ns = this.Ctx.Input.Params[":ns"]
    if ns == "" {
        result.Success = false 
        result.Msg = "ns is must"
        this.Data["json"] = result
        this.ServeJson()
    }
    tags = parseNs(ns)
    //check resource
    resource, exists := tags["resource"]
    if exists {
        sys_id, err :=  GetSysIdByName(resource)
        if err != nil {
            this.Data["json"] =  make([]string, 0)
            this.ServeJson()
        } else {
            delete(tags,"resource")
            this.Data["json"] = models.GetResourcesWithinNs(tags,sys_id)
            this.ServeJson()
        }
    }else{
        //resource not specify, return all 
        this.Data["json"] = models.GetResourcesWithinNs(tags,0)
        this.ServeJson()
    }
    fmt.Println(resources)
    fmt.Println("tags is: ", tags)
    this.Data["json"] = "resource not exists, get all"
	this.ServeJson()
}

//parse corp:didi,dept:dache to map
//if duplicate key,value use the last one
func parseNs(ns string) map[string]string {
    var tag_pairs []string 
    var tags map[string]string = make(map[string]string)

    ns = strings.Trim(strings.TrimSpace(ns), ",")
    tag_pairs = strings.Split(ns, ",")
    fmt.Println("tag pairs is : ", tag_pairs)

    for _,  pair := range tag_pairs {
        if pair != "" {
            pair = strings.TrimSpace(pair)
            pair_split := strings.Split(pair, ":")
            tagk := strings.TrimSpace(pair_split[0])
            tagv := strings.TrimSpace(pair_split[1])
            tags[tagk] = tagv
        }
    }
    return tags
}

//check if resource exists or not
func GetSysIdByName(name string) (int, error) {
    subsys, err := models.GetSubsysByName(name)
    if err != nil {
        fmt.Println("get subsys failed, ", err)
        return -1, err
    }
    return subsys.Id, nil
}
