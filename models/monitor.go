package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var type2name = map[int]string{0:"机器监控", 1:"业务监控"}

type Monitor struct {
	Id      int    `orm:"column(id);auto"`
	Name  string `orm:"column(name);size(255)"`
	Type    int    `orm:"column(type)"`
    TypeName string `orm:"-"`
    Class string `orm:"column(class);size(255)"`
	Comment string `orm:"column(comment);null"`
}

func init() {
	orm.RegisterModel(new(Monitor))
}


// AddMonitor insert a new Monitor into database and returns
// last inserted Id on success.
func AddMonitor(m *Monitor) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMonitorById retrieves Monitor by Id. Returns error if
// Id doesn't exist
func GetMonitorById(id int) (v *Monitor, err error) {
	o := orm.NewOrm()
	v = &Monitor{Id: id}
	if err = o.Read(v); err == nil {
        v.TypeName = type2name[v.Type]
		return v, nil
	}
	return nil, err
}


//GetAllMachineMonitor 获取所有机器监控项
func GetAllMachineMonitor() (ml []interface{}, err error) {
    var querys map[string]string = make(map[string]string)
    var fields []string
    var sortby []string
    var order  []string
    var offset int64
    var limit  int64

    querys["type"] =  "0"

    return GetAllMonitor(querys, fields, sortby, order, offset, limit)
}

// GetAllMonitor retrieves all Monitor matches certain condition. Returns empty list if
// no records exist
func GetAllMonitor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Monitor))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Monitor
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
        
        for index, monitor := range l {
            beego.Warn("type to name with monitor: ", monitor)
            //直接修改 l 内的TypeName
            l[index].TypeName = type2name[monitor.Type]        
            beego.Warn("type to name with monitor: ", monitor)
        }
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMonitor updates Monitor by Id and returns error if
// the record to be updated doesn't exist
func UpdateMonitorById(m *Monitor) (err error) {
	o := orm.NewOrm()
	v := Monitor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMonitor deletes Monitor by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMonitor(id int) (err error) {
	o := orm.NewOrm()
	v := Monitor{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Monitor{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
