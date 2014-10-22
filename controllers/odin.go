package controllers

import (
    "fmt"
    "ServiceTree/models"
    "encoding/json"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/httplib"
)

// Odin 兼容接口
type OdinController struct {
	beego.Controller
}

type Tagkvs struct {
    Checked bool `json:"checked"`
    Name    string `json:"name"`
    Value   []string `json:"value"`
}

func (this *OdinController) URLMapping() {
	this.Mapping("Get", this.Get)
}

// @Title Get
// @Description odin兼容接口, 查询 tag kv
// @Param	metric		query 	string	true		"采集项名字"
// @router /api/tagkv [get]
func (this *OdinController) Get() {
    var tagkvs []Tagkvs
	metric := this.GetString("metric")
    ns := this.GetString("ns")
    query := fmt.Sprintf("http://10.231.146.171/api/tagkv?metric=%s@host=", metric)
    result, err := httplib.Get(query).String()
    if err == nil {
        json.Unmarshal([]byte(result), &tagkvs) 
    }
    if len(tagkvs) == 0 {
        tagkvs = make([]Tagkvs, 0)
        beego.Warn("query tagkvs from odin and got nothing, may the input is wrong.")
    }
    this.Data["json"] = replace_host(tagkvs, ns)
    this.ServeJson()
}

//small the host only the ns 
func replace_host(tagkvs []Tagkvs, ns string) []Tagkvs {
    var ns_tagkvs []Tagkvs
    var tmp_tagkvs []Tagkvs
    var machines []models.Machine
    beego.Trace("name space is : ", ns)
    for _, tagkv := range tagkvs {
        if tagkv.Name != "host" {
            tmp_tagkvs = append(tmp_tagkvs, tagkv)
            continue
        }
        //用户选择对应的机器, 不需要过滤
        if tagkv.Checked {
            tmp_tagkvs = append(tmp_tagkvs, tagkv)
        }else{
        //用户默认全选, 需要过滤机器到对应的 ns 下面
            httpport := beego.AppConfig.String("httpport")
            if httpport == "" {
              //try 8080 as default if not config
              httpport = "8080"
            }
            query := fmt.Sprintf("http://127.0.0.1:%s/v1/resource/%s",httpport,ns)
            result, err := httplib.Get(query).String()
            if err == nil {
                err = json.Unmarshal([]byte(result), &machines)
                if err == nil {
                    var machine_names []string
                    for _, machine := range machines {
                        machine_names = append(machine_names, machine.Name)
                    }
                    tagkv.Value = machine_names
                }
            }else{
                beego.Debug("query resource from ns with error:", err)
            }
            ns_tagkvs = append(ns_tagkvs, tagkv)
        }
        ns_tagkvs  = append(ns_tagkvs, tmp_tagkvs...)
    } 
    return ns_tagkvs
}

