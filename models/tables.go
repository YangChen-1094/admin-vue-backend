package models

import (
	"my_gin/pkg/setting"
	"time"
)

//表明组成前缀+结构名（小写）

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

type Channel struct {
	ID        int    `json:"id"          gorm:"primary_key;column:id"`
	Name      string `json:"name"        gorm:"column:name"`
	Datetime  string `json:"datetime"    gorm:"column:datetime" `
	ChannelId string `json:"channelId"   gorm:"column:channelId" ` //表中要有的字段
}

type ItemConfig struct {
	Type       int `json:"type"             gorm:"column:type" `
	ItemId     int    `json:"itemid"        gorm:"primary_key;column:itemid"`
	Itemname   string `json:"itemname"      gorm:"column:itemname" `
	ProductID  string `json:"productID"     gorm:"column:productID"`
	Descr      string `json:"descr"         gorm:"column:descr" `
	ImgUrl     string `json:"imgUrl"        gorm:"column:imgUrl" ` //表中要有的字段
	ButtonUrl  string `json:"buttonUrl"     gorm:"column:buttonUrl" `
	BannerUrl  string `json:"bannerUrl"     gorm:"column:bannerUrl" `
	ButtonDesc string `json:"buttonDesc"    gorm:"column:buttonDesc" `
	Usetype    string `json:"usetype"       gorm:"column:usetype" `
	Price      string `json:"price"         gorm:"column:price" `
}

type ItemType struct {
	Id       int    `json:"id"              gorm:"primary_key;column:id"`
	Itemtype string `json:"itemtype"        gorm:"column:itemtype"`
	Typename string `json:"typename"        gorm:"column:typename" `
}

//定义表名，不然会视为加"s"的表： bh_item_configs表
func (ItemConfig) TableName() string {
	return setting.DeployConfig.Database.TablePrefix + "item_config"
}

func (ItemType) TableName() string {
	return setting.DeployConfig.Database.TablePrefix + "item_type"
}
