package models

type Permission struct {
	TasksetId int64 `orm:"column(taskset_id)"`
	GroupId   int64 `orm:"column(group_id)"`
	SysId     int64 `orm:"column(sys_id)"`
	Perm      int   `orm:"column(perm)"`
}
