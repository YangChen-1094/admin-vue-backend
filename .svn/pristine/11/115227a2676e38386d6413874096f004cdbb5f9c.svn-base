package models

import "time"

//标签包含的内容
type Tag struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  int64  `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn int64  `json:"modified_on"`
	State      int    `json:"state"` //状态 0为禁用、1为启用
}

//文章表字段信息
type Article struct {
	Id    int `json:"id"`
	TagId int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  string `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn string `json:"modified_on"`
	State      int    `json:"state"`
}

type Auth struct {
	ID       int           `gorm:"primary_key" json:"id"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Token    string        `json:"token"` //表中要有的字段
	Time     time.Duration `json:"time"`
}
