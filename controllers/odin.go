package controllers

import (
    "fmt"
    "strings"
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
    query := fmt.Sprintf("http://10.231.146.171/api/tagkv?metric=%s@host=", metric)
    result, err := httplib.Get(query).String()
    if err == nil {
        json.Unmarshal([]byte(result), &tagkvs) 
    }
    if len(tagkvs) == 0 {
        tagkvs = make([]Tagkvs, 0)
        beego.Warn("query tagkvs from odin and got nothing, may the input is wrong.")
    }
    this.Data["json"] = replace_host(metric, tagkvs)
    this.ServeJson()
}

// get metric ns
func getMetricNs(metric string) string {
    var ns_list []models.Ns
    var ns models.Ns
    var ns_string []string

    httpport := beego.AppConfig.String("httpport")
    query := fmt.Sprintf("http://127.0.0.1:%s/v1/ns?resource=%s&subsys=%s",httpport,metric,"monitor")
    result, err := httplib.Get(query).String()
    if err == nil {
        err = json.Unmarshal([]byte(result), &ns_list)
        beego.Debug("get ns from api: ", ns_list)
        if err == nil {
            if len(ns_list) != 1 {
                beego.Critical("more than one tag set")
                return  ""
            }
            ns = ns_list[0]
        }
    }else{
        beego.Debug("query resource from ns with error:", err)
    }
    if len(ns) > 0 {
        for _, tag := range ns {
            tag_str := tag.Key + ":" +  tag.Value
            ns_string = append(ns_string, tag_str) 
        }
        //specifiy the resource type
        ns_string = append(ns_string,"resource:machine")
    }
    return strings.Join(ns_string, ",")
} 

//small the host only the ns 
func replace_host(metric string, tagkvs []Tagkvs) []Tagkvs {
    var ns_tagkvs []Tagkvs
    var tmp_tagkvs []Tagkvs
    var machines []models.Machine

    for _, tagkv := range tagkvs {
        if tagkv.Name != "host" {
            tmp_tagkvs = append(tmp_tagkvs, tagkv)
            continue
        }
        //用户选择对应的机器, 不需要过滤
        if tagkv.Checked {
            tmp_tagkvs = append(tmp_tagkvs, tagkv)
        }else{
            //用户默认不选, 需要过滤机器到对应的 ns 下面
            httpport := beego.AppConfig.String("httpport")
            if httpport == "" {
              //try 8080 as default if not config
              httpport = "8080"
            }
            var ns string
            metric_part := strings.Split(metric, "@")
            ns = getMetricNs(metric_part[0])
            beego.Trace("get ns from api is : ", ns , "with resource: ", metric)
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

