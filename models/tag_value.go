package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
    "strconv"

	"github.com/astaxie/beego/orm"
)

type TagValue struct {
	Id    int    `orm:"column(id);pk"`
	KeyId int `orm:"column(key_id)"`
	Value string `orm:"column(value);size(255)"`
}

func init() {
	orm.RegisterModel(new(TagValue))
}

// AddTagValue insert a new TagValue into database and returns
// last inserted Id on success.
func AddTagValue(m *TagValue) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTagValueById retrieves TagValue by Id. Returns error if
// Id doesn't exist
func GetTagValueById(id int) (v *TagValue, err error) {
	o := orm.NewOrm()
	v = &TagValue{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetTagValueByKey(tag_key string) []string {
    o := orm.NewOrm()
    var tag_list orm.ParamsList        
    var tag_values []string

    o.Raw("select value from tag_value ,  tag_meta where tag_meta.tag_key = ? and tag_meta.id = tag_value.key_id", tag_key).ValuesFlat(&tag_list)
    for _, value := range  tag_list {
        if v, ok := value.(string) ; ok {
            tag_values = append(tag_values, v)
        } else {
            fmt.Println("warning:  convert to string failed.", ok)
        }
    }
    return tag_values
}

// GetTagAllValueByKeyId 根据 tag key id 来获取对应的 tag value 列表
// 如果不存在返回空字符
func GetTagAllValueByKeyId(key_id int) string {
    var values string
    var fields []string
    sortby := []string{"Value"}
    order :=  []string{"asc"}
    var offset  int64 = 0
    var limit   int64 = 10
    var query map[string]string = make(map[string]string)

    query["KeyId"] = strconv.Itoa(key_id)
    tag_list, err := GetAllTagValue(query, fields, sortby, order, offset, limit)
    if  err == nil {
        for _, value_model := range tag_list {
            if tag , ok := value_model.(TagValue) ; ok {
                values += tag.Value + ","
            }
        }
        values = strings.TrimRight(values,",")
    }
    return values
}

// GetAllTagValue retrieves all TagValue matches certain condition. Returns empty list if
// no records exist
func GetAllTagValue(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TagValue))
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

	var l []TagValue
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateTagValue updates TagValue by Id and returns error if
// the record to be updated doesn't exist
func UpdateTagValueById(m *TagValue) (err error) {
	o := orm.NewOrm()
	v := TagValue{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTagValue deletes TagValue by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTagValue(id int) (err error) {
	o := orm.NewOrm()
	v := TagValue{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TagValue{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func DeleteTagValueByKeyId(key_id int) (err error) {
    o := orm.NewOrm() 
    res, err := o.Raw("DELETE FROM tag_value WHERE key_id = ?", key_id).Exec()
    if err == nil {
        num, _ := res.RowsAffected()
        fmt.Println("delete tag value affect nums : ", num)
    }
    return 
}
