package v1

//标签页逻辑

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/setting"
	"my_gin/pkg/util"
	"strconv"
	"time"
)

type ApiTag struct {
}
func NewApiTag() *ApiTag{
	return &ApiTag{}
}

func (this *ApiTag)GetTags(ctx *gin.Context) {

	data := make(map[string]interface{})
	where := make(map[string]interface{})
	where["created_by"] = ctx.PostForm("user_id")

	state := 1
	if tmp := ctx.PostForm("state"); tmp != "" {
		state, _ = strconv.Atoi(tmp)
	}
	tagModel := models.NewModelTag()
	where["state"] = state

	data["list"] = tagModel.GetTagsList(util.GetPage(ctx, setting.AppSetting.PageSize), setting.AppSetting.PageSize, where)
	data["total"] = tagModel.GetTagsCount(where)

	global.JsonRet(ctx, 200, "", data)
}

func (this *ApiTag)AddTags(ctx *gin.Context) {
	name := ctx.PostForm("name")
	userId := ctx.PostForm("user_id")
	state := com.StrTo(ctx.DefaultPostForm("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 10, "name").Message("名称最长字符为10")
	valid.Required(userId, "user_id").Message("作者id不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	tagModel := models.NewModelTag()
	var tmp bool
	code := global.INVALID_PARAMS
	if !valid.HasErrors() {
		tmp = tagModel.AddTags(name, state, userId)
		log.Printf("name:%v, state:%v,user_id:%v, tmp:%v", name, state, userId, tmp)
	}

	if tmp {
		code = global.SUCCESS
	} else {
		code = global.ERROR
	}
	global.JsonRet(ctx, code, "", nil)
}

func (this *ApiTag)EditTags(ctx *gin.Context) {
	id := com.StrTo(ctx.PostForm("id")).MustInt()
	name := ctx.PostForm("name")
	userId := ctx.PostForm("user_id")
	state := com.StrTo(ctx.DefaultPostForm("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 10, "name").Message("名称最长字符为10")
	valid.Required(userId, "user_id").Message("作者id不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	var tmp bool
	tagModel := models.NewModelTag()
	code := global.INVALID_PARAMS
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["modified_by"] = userId
		data["state"] = state
		data["name"] = name
		data["modified_on"] = time.Now().Unix()
		log.Println("update data:", data)
		if tagModel.CheckTagExistsById(id) {
			tmp = tagModel.EditTags(id, data)
		}

	}

	if tmp {
		code = global.SUCCESS
	} else {
		code = global.ERROR
	}
	global.JsonRet(ctx, code, "", nil)
}

func (this *ApiTag)DelTags(ctx *gin.Context) {
	id := com.StrTo(ctx.PostForm("id")).MustInt()
	var ret bool
	var code int
	tagModel := models.NewModelTag()
	if tagModel.CheckTagExistsById(id) {
		ret = tagModel.DelTags(id)
	}
	if ret {
		code = global.SUCCESS
	} else {
		code = global.ERROR
	}
	global.JsonRet(ctx, code, "", nil)
}
