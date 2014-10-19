package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

/*
select t1.sys_id, t1.resource_id, concat(t1.tag_value,",",t2.tag_value) as tags from
(select * from tagging where tag_name = 'corp') t1 ,
(select * from tagging where tag_name = 'depart') t2
where t1.sys_id = t2.sys_id and t1.resource_id = t2.resource_id;
*/

type TreeNode struct {
	Meta   string `json:"meta"`
	Name string    `json:"name"`
    IsParent bool `json:"isParent"`
    IconSkin string `json:"iconSkin"`
	//leafs 和 children至少有一个不为空，leafs代表下面是resource 节点，children代表子节点（非叶子节点）
	Resources    []*Resource  `json:"resources"`
	Children []*TreeNode `json:"children"`
}

type QueryError struct {
	Code int
	Msg  string
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("[err code:%d]%s", e.Code, e.Msg)
}

func BuildTreeByTagString(dbname, tag_str string, s_id int) (root *TreeNode, err error) {
	resources, err := queryResourcesByTagString(dbname, tag_str, s_id)
	if err != nil {
		fmt.Println("queryResourceByTagString failed")
		return nil, &QueryError{-1, "queryResourceByTagString failed"}
	}

	root = &TreeNode{Meta:"Root", Name:"Root",  Resources:nil, Children:[]*TreeNode{}, IsParent:true,IconSkin:"pIcon01"}
	for _, res := range resources {
		buildTree(root, tag_str, res)
	}

	return root, err
}

func buildTree(root *TreeNode, tags string, res Resource) {
	keys := strings.Split(tags, ",")
	cur_node := root
	for i, val := range strings.Split(res.Tags, ",") {
		created := false
		for _, child := range cur_node.Children {
			if child.Meta == keys[i] && child.Name == val {
				created = true
				cur_node = child
			}
		}
		if !created {
            new_node := &TreeNode{Meta:keys[i], Name:val, Resources:nil, Children:[]*TreeNode{}, IsParent:true, IconSkin:"pIcon01"}
			cur_node.Children = append(cur_node.Children, new_node)
			cur_node = new_node
		}
	}

	if cur_node.Resources == nil {
		cur_node.Resources = []*Resource{}
        cur_node.IsParent = false
        cur_node.IconSkin = "icon01"
	}

	cur_node.Resources = append(cur_node.Resources, &Resource{SysId:res.SysId, ResourceId:res.ResourceId, Tags:tags})
}

func queryResourcesByTagString(dbname, tag_str string, s_id int) (resources []Resource, err error) {
	o := orm.NewOrm()
	o.Using(dbname)

	q_str, _ := buildQueryString(tag_str, s_id)

	fmt.Println("buildDBQueryString:\n", q_str)
	num, err := o.Raw(q_str).QueryRows(&resources)
	if err == nil {
		fmt.Printf("Query successed and row count is %d\n", num)
		for i, tuple := range resources {
			fmt.Printf("%d  sys_id:%d resource_id:%d tags:%s\n", i+1, tuple.SysId, tuple.ResourceId, tuple.Tags)
		}
	} else {
		fmt.Println("Query Error!")
	}
	return resources, err
}

func buildQueryString(tag_string string, s_id int) (q_string string, err error) {
	tags := strings.Split(tag_string, ",")

	var complete_sql bytes.Buffer
	var cat_sql bytes.Buffer
	var sel_sql bytes.Buffer
	var where_sql bytes.Buffer

	len := len(tags)

	cat_sql.WriteString("concat(")
	for i, tag := range tags {
		j := i + 1
		cat_sql.WriteString(fmt.Sprintf("t%d.tag_value", j))
        if s_id <= 0{
		    sel_sql.WriteString(fmt.Sprintf("(select * from tagging where tag_name = \"%s\") t%d", tag, j))
        }else{
		    sel_sql.WriteString(fmt.Sprintf("(select * from tagging where tag_name = \"%s\" AND sys_id=%d) t%d", tag, s_id, j))
        }
		if j < len {
			cat_sql.WriteString(",\",\",")
			sel_sql.WriteString(" ,")
		}
		sel_sql.WriteString("\n")

		k := (j + 1)
		if k > len {
			k = 1
		}

		if j < k || len > 2 {
            if s_id <= 0{
			    where_sql.WriteString(fmt.Sprintf("t%d.sys_id = t%d.sys_id AND t%d.resource_id = t%d.resource_id\n", j, k, j, k))
            }else{
			    where_sql.WriteString(fmt.Sprintf("t%d.resource_id = t%d.resource_id\n", j, k))
            }
		}

		if j < len && len > 2 {
			where_sql.WriteString("AND\n")
		}

	}
	cat_sql.WriteString(")")
	complete_sql.WriteString(fmt.Sprintf("select t1.sys_id, t1.resource_id, %s as tags from \n", cat_sql.String()))
	complete_sql.WriteString(sel_sql.String())
	if len > 1 {
		complete_sql.WriteString(fmt.Sprintf("where %s;", where_sql.String()))
	}
	return complete_sql.String(), nil
}

//构建服务树
//sys_id:表示资源对应的编号  0: 获取全部资源
func BuildTree(hierarchy string, sys_id int) *TreeNode{

	root, err := BuildTreeByTagString("alfred", hierarchy, sys_id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		depth := 0
		printTree(root, depth)
	}
    return root
}

//get  specify resource within specify ns
// sys_id :0  get all resource within ns
func GetResourcesWithinNs(tags map[string]string, sys_id int) []interface{}{
    var hierarchy []string 
    var names []string
    var resources []interface{}
    var query_resources []*Resource
    var query_all bool 

    if sys_id == 0 {
        query_all = true
    }

    for k, v := range tags {
        hierarchy = append(hierarchy, k)
        names = append(names , v)
    }
    fmt.Printf(">>>>>>>>>start search with: tags:%+v, hierarchy:%+v, names:%+v\n", tags, hierarchy, names)

    root := BuildTree(strings.Join(hierarchy, ","), sys_id) 
    query_resources = searchResources(root,hierarchy,names,0)

    for _, query_resource := range query_resources {
        if query_all {
            sys_id = query_resource.SysId
        }
        resource_id := query_resource.ResourceId
        subsys, err := GetSubsysById(sys_id)
        if err != nil {
            fmt.Println("wfwefwefwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww=> query wrong sys id, err:", err, "with sys_id:", sys_id)
            continue
        }
        fmt.Println("wfwefwefwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww=> start orm")
        if subsys.Name == "machine" {
            r, err := GetMachineById(resource_id)
            if err == nil {
               resources = append(resources, r) 
            }
        }
        if subsys.Name == "monitor" {
            r, err := GetMonitorById(resource_id)
            if err == nil {
               resources = append(resources, r) 
            }
        }
    }
    fmt.Println("============================> get reources: ", resources)
    return resources
}


func searchResources(node *TreeNode, hierarchy []string, names []string, depth int) []*Resource{
    var resources []*Resource
    if depth == 0 {
        for _, child := range node.Children {
            fmt.Printf("root-> children: %+v\n", child)
                resources = searchResources(child, hierarchy, names, depth+1)
                if len(resources) == 0 {
                    fmt.Println("search another node")
                    depth = 0
                    continue
                }else{
                    return resources
                }
        }
    }else{
        fmt.Printf("before search %+v, hierarchy:%+v, names:%+v\n", node, hierarchy, names )
        if (node.IsParent == false) && (node.Meta == hierarchy[0]) && (node.Name == names[0]) {
            fmt.Println("found , return node:", node)
            resources = node.Resources
        }else{
            if (node.Meta == hierarchy[0]) && (node.Name == names[0]) {
                for _, child := range node.Children {
                    fmt.Printf("check child %+v, hierarchy:%+v, names:%+v\n", child, hierarchy, names )
                    if (child.Meta == hierarchy[1]) && (child.Name == names[1]) {
                        fmt.Printf("node:%+v match .start search next\n", child)
                        return searchResources(child, hierarchy[1:], names[1:], depth+1)
                    }
                }
            }
        }
    }
    return resources
}


func printTree(root *TreeNode, dep int) {
	for i := 0; i < dep; i++ {
		fmt.Print("	")
	}
	if dep == 0 {
		fmt.Println("Root")
	} else {
		fmt.Println(root.Meta, root.Name)
	}
	if root.Resources != nil {
		for _, leaf := range root.Resources {
			for i := 0; i < dep+1; i++ {
				fmt.Print("	")
			}
			fmt.Printf("Resource[%d,%d]\n", leaf.SysId, leaf.ResourceId)
		}
	} else if root.Children != nil {
		for _, ch := range root.Children {
			printTree(ch, dep+1)
		}
	} else {
		fmt.Println("something wrong!")
	}

}
