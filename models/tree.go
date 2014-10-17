package models


type NodeResource struct {
    Sys_id      int                                                                                                                                                          
    Resource_id int
    Tag_string  string
}

type TreeNode struct {
    Name string `json:"name"`
    Meta string `json:"meta"`
    IsParent bool `json:"isParent"`
    IconSkin string `json:"iconSkin"`
    Children []*TreeNode `json:"children"`
}


type Tree struct {
    Root TreeNode
}
