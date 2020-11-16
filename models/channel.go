package models

import (
	"fmt"
	"my_gin/pkg/logger"
	"strings"
	"time"
)

type ModelChannel struct {
}

func (this *ModelChannel) GetChannelList(page int, size int) (chs []Channel) {//channel表
	if size <= 0{
		size = 20
	}
	if page < 0{
		page = 0
	}
	offset := (page - 1) * size
	db.LogMode(true)
	db.Limit(size).Offset(offset).Find(&chs)
	return
}


func (this *ModelChannel) Modify(id int, param map[string]interface{}) bool {//channel表
	if id <= 0 {
		return false
	}
	name := param["name"]
	channel_id := param["channel_id"]
	if len(name.(string)) > 0 {
		param["name"] = strings.Trim(name.(string), " ")
	}
	if len(channel_id.(string)) > 0 {
		param["channel_id"] = strings.Trim(channel_id.(string), " ")
	}
	timeUnix := time.Now().Unix()
	param["datetime"] = fmt.Sprintf("%v", timeUnix)
	db.LogMode(true)
	err := db.Model(&Channel{}).Where("id=?", id).Updates(param).Error
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
	db.LogMode(true)
	//err := db.Debug().Delete(&Channel{ID: id}).Error
	err := db.Where(&Channel{ID: id}).Delete(Channel{}).Error
	if err!= nil {
		logger.Info("mysql", "channel [Del]", " Del err: ", err)
		return false
	}
	return true
}

func (this *ModelChannel) Add(param map[string]interface{}) bool {//channel表
	name := strings.Trim(param["name"].(string), " ")
	channel_id := strings.Trim(param["channel_id"].(string), " ")
	timeUnix := time.Now().Unix()
	datetime := fmt.Sprintf("%v", timeUnix)
	db.LogMode(true)
	err := db.Create(&Channel{Name: name, ChannelId: channel_id,Datetime: datetime}).Error
	if err!= nil {
		logger.Info("mysql", "channel [Add]", " Add err: ", err)
		return false
	}
	return true
}