package models

type Tagging struct {
	TagsetId   int64     `orm:"column(tagset_id)"`
	SysId      *Subsys   `orm:"column(sys_id);rel(fk)"`
	ResourceId int64     `orm:"column(resource_id)"`
	TagName    *TagMeta  `orm:"column(tag_name);rel(fk)"`
	TagValue   *TagValue `orm:"column(tag_value);rel(fk)"`
}
