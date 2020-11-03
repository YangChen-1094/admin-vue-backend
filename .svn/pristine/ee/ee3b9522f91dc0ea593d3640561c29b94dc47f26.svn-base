package models

import "time"

//类
type ModelTag struct {
}

//获取文章列表
func (this *ModelTag) GetTagsList(num int, size int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(num).Limit(size).Find(&tags)
	return
}

func (this *ModelTag) GetTagsCount(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func (this *ModelTag) AddTags(name string, status int, createBy string) bool {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createBy,
		CreatedOn: time.Now().Unix(),
		State:     status,
	})
	return true
}

func (this *ModelTag) EditTags(id int, data map[string]interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Updates(data)
	return true
}

func (this *ModelTag) CheckTagExistsById(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	if tag.Id > 0 {
		return true
	}
	return false
}

func (this *ModelTag) DelTags(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}
