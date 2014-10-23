package models

import (
    "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Tagging  struct {
    SysId      int64 `json:SysId`
    ResourceId int64 `json:ResourceId`
    Tags       []Tag `json:"Tags"`
}

type TaggingRecord struct {
    TagsetId   int64  `json:"tagset_id"`
    SysId      int64  `json:"sys_id"`
    ResourceId int64  `json:"resource_id"`
    TagName    string  `json:"tag_name"`
    TagValue   string  `json:"tag_value"` 
}

// 对对应的资源新增 tag
func AddTagging(taggings []Tagging) int64{
    id := getTaggingId()
    if id != -1 {
        o := orm.NewOrm() 
        for _, tagging := range  taggings {
            fmt.Printf("tags: %+v\n", tagging.Tags)
            for _, tag := range tagging.Tags {
                beego.Trace("start insert tag:", tag)
                sql := fmt.Sprintf("INSERT INTO tagging values(%d,%d,%d,'%s','%s')", id, tagging.SysId, tagging.ResourceId, tag.Key, tag.Value)
                beego.Trace("add tagging with sql:", sql)
                res, err := o.Raw(sql).Exec()
                if err == nil {
                    num, _ := res.RowsAffected()
                    if num != 1 {
                        beego.Warn("insert tag record seems failed. ")
                        return -1
                    }
                }else{
                    beego.Critical("tagging failed with err: ", err)
                    return -1
                }
            }    
        }
        return 0
    }
    return id
}

//获取对应的 业务 id
func getTaggingId()  int64{
    var idalloc Idalloc
    idalloc.Api = "tagging"
    if id, err := AddIdalloc(&idalloc); err == nil {
        fmt.Printf("id alloc for api:%s , value is: [%d] ", idalloc.Api, id)
        return  id
    }else{
        fmt.Println("id alloc failed", err)
    }
    return -1
}
