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
	if page < 0{
		page = 0
	}
	offset := (page - 1) * size
	Db.LogMode(true)
	Db.Limit(size).Offset(offset).Find(&chs)
	fmt.Println(setting.DatabaseSetting)
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
	channel_id := strings.Trim(param["channel_id"].(string), " ")
	timeUnix := time.Now().Unix()
	datetime := fmt.Sprintf("%v", timeUnix)
	Db.LogMode(true)
	err := Db.Create(&Channel{Name: name, ChannelId: channel_id,Datetime: datetime}).Error
	if err!= nil {
		logger.Info("mysql", "channel [Add]", " Add err: ", err)
		return false
	}
	return true
}

//批量插入渠道信息
func (this *ModelChannel) BatchAdd(aFields []map[string]string) (err error) {
	var item map[string]string
	var aData []interface{}
	for _, item = range aFields {
		if len(item["name"]) <= 0{
			continue
		}
		if len(item["channel_id"]) <= 0{
			continue
		}
		one := Channel{}
		one.Name = item["name"]
		one.ChannelId = item["channel_id"]
		one.Datetime = fmt.Sprintf("%v", time.Now().Unix())
		aData = append(aData, one)
	}
	if len(aData) <= 0 {
		return errors.New("批量插入数据为空")
	}
	var insert []Channel
	sql := util.GetBranchInsertSql(aData, setting.DatabaseSetting.TablePrefix + "channel")
	err = Db.Raw(sql).Scan(&insert).Error
	return err
}
