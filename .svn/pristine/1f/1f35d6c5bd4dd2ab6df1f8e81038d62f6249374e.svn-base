package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
)

//文章接口

type ApiArticle struct {
}
func NewApiArticle() *ApiArticle{
	return &ApiArticle{}
}

// @Summary 获取单个文章信息
// @Description  根据文章id获取文章信息
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   id path    int     true "文章id"
// @Success 200 {string} string	"ok"
// @Router /v1/article/info [post]
func (this *ApiArticle)GetArticleInfo(ctx *gin.Context){
	id := com.StrTo(ctx.PostForm("id")).MustInt()
	ModelArticle := models.NewModelArticle()
	if !ModelArticle.CheckArticleExists(id) {
		global.JsonRet(ctx, 404, "该文章不存在", nil)
		return
	}
	info := ModelArticle.GetArticleInfo(id)
	global.JsonRet(ctx, 200, "", info)
}


// @Summary 获取文章列表
// @Description  用户id获取文章列表
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param   user_id header    int   true "用户id"
// @Param   state header    int   true "禁用状态0禁用，1正常"
// @Param   page header    int   true "页数"
// @Param   size header    int   true "每页多少文章"
// @Param   auth_id header    int   true "验签id"
// @Param   token header    string   true "token"
// @Success 200 {string} string	"ok"
// @Router /v1/article/getList [post]
func (this *ApiArticle)GetArticleList(ctx *gin.Context){
	state := com.StrTo(ctx.DefaultPostForm("state", "0")).MustInt()
	page := com.StrTo(ctx.DefaultPostForm("page", "1")).MustInt()//查询第几页
	size := com.StrTo(ctx.DefaultPostForm("size", "2")).MustInt()//每页多少条数据
	user_id := com.StrTo(ctx.DefaultPostForm("user_id", "0")).MustInt()
	where := make(map[string]interface{})
	data := make(map[string]interface{})
	test, bool := ctx.Get("test")
	if !bool {
		global.JsonRet(ctx, 200, "", data)
		return
	}

	offset := (page - 1) * size //查询偏移量
	where["state"] = state
	if user_id != 0 {
		where["created_by"] = user_id
	}
	ModelArticle := models.NewModelArticle()
	data["test"] = test
	data["list"] = ModelArticle.GetArticleList(offset, size, where)
	data["total"] = ModelArticle.GetArticleCount(where)
	logger.Info("articleList", data)
	global.JsonRet(ctx, 200, "", data)
}

func  (this *ApiArticle)AddArticle(ctx *gin.Context){
	title := ctx.PostForm("title")
	tagId := com.StrTo(ctx.PostForm("tag_id")).MustInt()
	desc := ctx.PostForm("desc")
	content := ctx.PostForm("content")
	createdBy := ctx.PostForm("created_by")
	state := com.StrTo(ctx.DefaultPostForm("state", "0")).MustInt()

	v := validation.Validation{}
	v.Min(tagId, 1, "tag_id").Message("标签id必须大于0")
	v.Required(title, "title").Message("标题不能为空")
	v.Required(desc, "desc").Message("简述不能为空")
	v.Required(content, "content").Message("内容不能为空")
	v.Required(createdBy, "created_by").Message("创建人不能为空")
	v.MaxSize(title, 20, "title").Message("标题长度限制10字符")
	v.MaxSize(desc, 50, "desc").Message("描述长度限制20字符")
	v.MaxSize(content, 200, "content").Message("内容长度限制200字符")
	v.Range(state,0,1, "state").Message("状态只允许0或1")
	log.Println(tagId, title, desc, state)
	if v.HasErrors() {
		var msg string
		for _, err := range v.Errors {
			msg += err.Key + ":" + err.Message + ";"
		}

		global.JsonRet(ctx, global.INVALID_PARAMS, msg, nil)
		return
	}

	data := make(map[string]interface{})
	data["title"] = title
	data["tag_id"] = tagId
	data["desc"] = desc
	data["state"] = state
	data["content"] = content
	data["created_by"] = createdBy

	ModelArticle := models.NewModelArticle()
	ret := ModelArticle.AddArticle(data)
	if ret == false {
		global.JsonRet(ctx, global.ERROR, "添加文章失败", nil)
		return
	}

	global.JsonRet(ctx, global.SUCCESS, "添加文章成功", nil)
	return
}