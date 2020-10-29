package models

type ModelArticle struct {
}

func (this *ModelArticle) GetArticleInfo(id int) (art Article) {
	db.Where("id=?", id).First(&art)
	db.Model(&art).Related(&art.Tag) //把tag表中的信息关联到tag字段
	return
}

func (this *ModelArticle) CheckArticleExists(id int) bool {
	var art Article
	db.Select("id").Where("id=?", id).First(&art)
	if art.Id > 0 {
		return true
	}
	return false
}

func (this *ModelArticle) GetArticleList(offset int, size int, maps interface{})(artList []Article){
	db.Preload("Tag").Where(maps).Offset(offset).Limit(size).Find(&artList)
	return
}

func (this *ModelArticle) GetArticleCount(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func (this *ModelArticle) AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId: data["tag_id"].(int),
		Title: data["title"].(string),
		Desc: data["desc"].(string),
		Content: data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State: data["state"].(int),
	})
	return true
}
