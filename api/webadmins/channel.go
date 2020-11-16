package webadmins

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"my_gin/models"
	"my_gin/pkg/global"
)

type channel struct {
}
func NewChannel() *channel{
	return &channel{}
}

func (this *channel) List(ctx *gin.Context){
	modelChannel := models.NewModelChannel()
	list := modelChannel.GetChannelList(1, 30)
	data := make(map[string]interface{})
	data["list"] = list
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}

func (this *channel) Modify(ctx *gin.Context){
	data := make(map[string]interface{})
	param := make(map[string]interface{})
	modelChannel := models.NewModelChannel()
	id := com.StrTo(ctx.PostForm("id")).MustInt()
	param["name"] = com.StrTo(ctx.PostForm("name")).String()
	param["channel_id"] = com.StrTo(ctx.PostForm("channel_id")).String()
	update := modelChannel.Modify(id, param)
	if update {
		global.JsonRet(ctx, global.SUCCESS, "", data)
	}else{
		global.JsonRet(ctx, global.ERROR, "", data)
	}

	return
}

func (this *channel) Delete(ctx *gin.Context){
	data := make(map[string]interface{})
	modelChannel := models.NewModelChannel()
	id := com.StrTo(ctx.PostForm("id")).MustInt()
	del := modelChannel.Del(id)
	if del {
		global.JsonRet(ctx, global.SUCCESS, "", data)
	}else{
		global.JsonRet(ctx, global.ERROR, "", data)
	}
	return
}

func (this *channel) Add(ctx *gin.Context){
	data := make(map[string]interface{})
	param := make(map[string]interface{})
	modelChannel := models.NewModelChannel()
	param["name"] = com.StrTo(ctx.PostForm("name")).String()
	param["channel_id"] = com.StrTo(ctx.PostForm("channel_id")).String()

	add := modelChannel.Add(param)
	if add {
		global.JsonRet(ctx, global.SUCCESS, "", data)
	}else{
		global.JsonRet(ctx, global.ERROR, "", data)
	}
	return
}
