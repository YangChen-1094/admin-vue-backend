package models

import (
	"my_gin/pkg/logger"
	"my_gin/pkg/setting"
)

type ModelItem struct {
	TblItemConfig string
	TblItemType string
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

func (this *ModelItem) GetItemList(page int, size int) (itemList []ItemConfig) {
	if size <= 0{
		size = 20
	}
	if page <= 0{
		page = 1
	}
	itemList = []ItemConfig{ItemConfig{}}
	offset := (page - 1) * size
	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "") //mhjy_admins数据库
	Mysql.Table(this.TblItemConfig).Limit(size).Offset(offset).Find(&itemList)
	//Db.Table(this.TblItemConfig).Limit(size).Offset(offset).Find(&itemList) //my_gin数据库
	return
}
func (this *ModelItem) GetAllItem() (chs []ItemConfig) {
	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "")
	Mysql.Table(this.TblItemConfig).Find(&chs)
	//Db.Table(this.TblItemConfig).Find(&chs)
	return
}

func (this *ModelItem) GetItemCount() (count int)  {
	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "")
	Mysql.Table(this.TblItemConfig).Count(&count)
	//Db.Table(this.TblItemConfig).Count(&count)
	return
}
func (this *ModelItem) GetItemType() (list []ItemType) {
	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "")
	Mysql.Table(this.TblItemType).Find(&list)
	//Db.Table(this.TblItemType).Find(&list)
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
	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "")
	err = Mysql.Table(this.TblItemConfig).Create(&item).Error
	//err = Db.Table(this.TblItemConfig).Create(&item).Error
	return
}

func (this *ModelItem) Modify(id int, param map[string]interface{}) (err error) {
	wh := make(map[string]interface{})
	wh["itemid"] = id
	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "")
	err = Mysql.Table(this.TblItemConfig).Where(wh).Update(param).Error
	//err = Db.Table(this.TblItemConfig).Where(wh).Update(param).Error
	return
}

func (this *ModelItem) Del(id int) bool {//删除道具配置
	if id <= 0 {
		return false
	}

	Mysql := setting.Mysql.GetMysqlClient("wbAdminDatas", "")
	err := Mysql.Table(this.TblItemConfig).Where(&ItemConfig{ItemId: id}).Delete(ItemConfig{}).Error
	//err := Db.Table(this.TblItemConfig).Where(&ItemConfig{ItemId: id}).Delete(ItemConfig{}).Error
	if err!= nil {
		logger.Info("mysql", "channel [Del]", " Del err: ", err)
		return false
	}
	return true
}