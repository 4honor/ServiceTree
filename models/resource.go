package models

import (
        "fmt"
        "strings"
        "github.com/astaxie/beego/orm"
       )

type Resource struct {
    SysId int `json:"SysId"`
    ResourceId int `json:"ResourceId"`
    Name string `json:"Name"`
    Tags string `json:"-"`
}


//Filter resource to check if the resource is record
//为了提高查询速度, 一次查询只能允许相同类型进行查询
func FilterResources(resources []Resource) []orm.Params{
    var sys_name string 
    var sys_id  int 
    var valid_resource []orm.Params
    var query_resources  string

    for index, resource := range resources {
        if index == 0 {
            sys_id = resource.SysId
        }
       query_resources +=  "'" + resource.Name + "',"
    }
    query_resources =  strings.TrimRight(query_resources, ",")

    fmt.Printf("query resources name: >>>>  %s and sys id is: [%d]\n", sys_name, sys_id)
    if sys_id != 0 {
        subsys, _ := GetSubsysById(sys_id)
        if subsys == nil {
            fmt.Println("[Fatal] unable to locate subsys") 
        }
        sys_name = subsys.Name
    }
    
    if len(resources) > 0 {
        var maps []orm.Params
        o := orm.NewOrm()
        query_sql := fmt.Sprintf("SELECT subsys.id as SysId, %s.name as Name, %s.id as ResourceId from %s, subsys where subsys.name = '%s' and %s.name in (%s)", sys_name, sys_name, sys_name, sys_name, sys_name, query_resources) 
        fmt.Println("query sql: " , query_sql)
        num, err := o.Raw(query_sql).Values(&maps)
        if err == nil && num > 0 {
            return maps
        }else{
            fmt.Printf("[Fatal] query failed." ,err)
        }
    }
    return valid_resource
}
