package models

type TagSchema struct {
	Key     string
	Value   []string
}

type Tag struct {
    Key     string
    Value   string
}

type Ns []Tag
