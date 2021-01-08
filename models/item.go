package models

import (
	"fmt"
)

type ModelItem struct {
}

var UseType = make(map[int]string)
var UseTypeNameId = make(map[string]int)

func init(){
	UseType[1001] = "弹窗"
	UseType[1002] = "计费点"
	UseTypeNameId = useTypeNameId()
}

func useTypeNameId() map[string]int {
	var nameId = make(map[string]int)
	for id, name := range UseType {
		nameId[name] = id
	}
	return nameId
}

func (this *ModelItem) GetItemList(page int, size int) (chs []ItemConfig) {
	if size <= 0{
		size = 20
	}
	if page <= 0{
		page = 1
	}
	offset := (page - 1) * size
	Db.Limit(size).Offset(offset).Find(&chs)
	return
}
func (this *ModelItem) GetAllItem() (chs []ItemConfig) {
	Db.Find(&chs)
	return
}

func (this *ModelItem) GetItemCount() (count int)  {
	Db.Model(&ItemConfig{}).Count(&count)
	return
}
func (this *ModelItem) GetItemType() (list []ItemType) {
	Db.Find(&list)
	return
}

func (this *ModelItem) Insert(param map[string]interface{}) (err error) {
	item := ItemConfig{
		ItemId:param["itemid"].(int),
		Itemname: param["itemname"].(string),
		Type:param["type"].(int),
		ProductID: param["productID"].(string),
		Descr: param["descr"].(string),
		ImgUrl: param["imgUrl"].(string),
		ButtonUrl: param["buttonUrl"].(string),
		BannerUrl: param["bannerUrl"].(string),
		ButtonDesc: param["buttonDesc"].(string),
		Usetype: param["usetype"].(string),
		Price: param["price"].(string),
	}
	fmt.Println("item", item)
	err = Db.Create(&item).Error
	return
}

func (this *ModelItem) Modify(id int, param map[string]interface{}) (err error) {
	wh := make(map[string]interface{})
	wh["itemid"] = id
	err = Db.Model(&ItemConfig{}).Where(wh).Update(param).Error
	return
}