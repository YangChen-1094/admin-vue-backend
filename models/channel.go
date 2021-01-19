package models

import (
	"errors"
	"fmt"
	"my_gin/pkg/logger"
	"my_gin/pkg/setting"
	"my_gin/pkg/util"
	"strings"
	"time"
)

type ModelChannel struct {
}

func (this *ModelChannel) GetChannelList(page int, size int) (chs []Channel) {//channel表
	if size <= 0{
		size = 20
	}
	if page <= 0{
		page = 1
	}
	offset := (page - 1) * size
	Db.LogMode(true)
	Db.Limit(size).Order("id desc").Offset(offset).Find(&chs)
	return
}

func (this *ModelChannel) GetChannelCount() (count int)  {//channel表
	Db.Model(&Channel{}).Count(&count)
	return
}



func (this *ModelChannel) Modify(id int, param map[string]interface{}) bool {//channel表
	if id <= 0 {
		return false
	}
	name := param["name"]
	channelId := param["channelId"]
	if len(name.(string)) > 0 {
		param["name"] = strings.Trim(name.(string), " ")
	}
	if len(channelId.(string)) > 0 {
		param["channelId"] = strings.Trim(channelId.(string), " ")
	}
	timeUnix := time.Now().Unix()
	param["datetime"] = fmt.Sprintf("%v", timeUnix)
	Db.LogMode(true)
	err := Db.Model(&Channel{}).Where("id=?", id).Updates(param).Error
	if err!= nil {
		logger.Info("mysql", "channel [Modify]", " update err: ", err)
		return false
	}
	return true
}


func (this *ModelChannel) Del(id int) bool {//channel表
	if id <= 0 {
		return false
	}
	Db.LogMode(true)
	//err := Db.Debug().Delete(&Channel{ID: id}).Error
	err := Db.Where(&Channel{ID: id}).Delete(Channel{}).Error
	if err!= nil {
		logger.Info("mysql", "channel [Del]", " Del err: ", err)
		return false
	}
	return true
}

func (this *ModelChannel) Add(param map[string]interface{}) bool {//channel表
	name := strings.Trim(param["name"].(string), " ")
	channelId := strings.Trim(param["channelId"].(string), " ")
	timeUnix := time.Now().Unix()
	datetime := fmt.Sprintf("%v", timeUnix)
	Db.LogMode(true)
	err := Db.Create(&Channel{Name: name, ChannelId: channelId,Datetime: datetime}).Error
	if err!= nil {
		logger.Info("mysql", "channel [Add]", " Add err: ", err)
		return false
	}
	return true
}

//批量插入渠道信息
func (this *ModelChannel) BatchAdd(aFields []map[string]string) (err error) {
	var item map[string]string
	var aInsert []interface{}
	var aUpdate []interface{}
	var keys []interface{}
	list := this.GetAllChannel()
	var channelIds []interface{}//已经存在的渠道id
	for _, oneChannel := range list{
		channelIds = append(channelIds, oneChannel.ChannelId)
	}
	for _, item = range aFields {
		if len(item["name"]) <= 0{
			continue
		}
		if len(item["channelId"]) <= 0{
			continue
		}

		one := Channel{}
		one.Name = item["name"]
		one.ChannelId = item["channelId"]
		one.Datetime = fmt.Sprintf("%v", time.Now().Unix())

		isUpdate := false
		for _, oneChannel := range list {
			if oneChannel.ChannelId == item["channelId"] {
				one.ID = oneChannel.ID
				isUpdate = true
				keys = append(keys, oneChannel.ID)
				break
			}
		}

		if isUpdate {//需要更新的数据
			aUpdate = append(aUpdate, one)
		}else{//需要插入的数据
			aInsert = append(aInsert, one)
		}
	}
	if len(aInsert) <= 0 && len(aUpdate) <= 0 {
		return errors.New("批量插入数据为空")
	}

	sql:=""
	if len(aInsert) > 0 {
		var insert []Channel
		sql := util.GetBranchInsertSql(aInsert, setting.DeployConfig.Database.TablePrefix + "channel")
		err = Db.Raw(sql).Scan(&insert).Error
	}

	if len(aUpdate) > 0 {
		var update []Channel
		notUpdate := []string{
			"id", "channelId",
		}
		sql = util.GetBranchUpdateSql(aUpdate, setting.DeployConfig.Database.TablePrefix + "channel", keys, notUpdate, "")
		err = Db.Raw(sql).Scan(&update).Error
	}
	return err
}

func (this *ModelChannel) GetAllChannel() (chs []Channel) {//channel表
	Db.LogMode(true)
	Db.Find(&chs)
	return
}