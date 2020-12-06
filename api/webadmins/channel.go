package webadmins

import (
	"bufio"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"io"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/util"
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
	redisPool := models.RedisMgr.RdsClient.PoolStats()
	data["list"] = list
	data["pool"] = redisPool
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

func (this *channel) Import(ctx *gin.Context){
	data := make(map[string]interface{})
	modelChannel := models.NewModelChannel()

	channelFile, err := ctx.FormFile("channelFile")
	if err != nil {
		global.JsonRet(ctx, global.ERROR, "文件格式错误", data)
		return
	}
	if channelFile.Size > global.WEB_ADMINS_MAX_UPLOAD_SIZE {
		global.JsonRet(ctx, global.ERROR, "文件大小超过限制", data)
		return
	}

	file, err := channelFile.Open()
	if err != nil {
		global.JsonRet(ctx, global.ERROR, "文件格式错误", data)
		return
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	var aField []map[string]string
	for{
		aCh, err := reader.Read()
		if err == io.EOF {
			break
		}else if err != nil {
			global.JsonRet(ctx, global.ERROR, "文件内容读取错误", data)
			return
		}
		for i:=0; i< len(aCh); i ++ {
			aCh[i], _ = util.ConvertToString(aCh[i], "GBK", "UTF-8")
		}
		one := make(map[string]string)
		one["channel_id"] = aCh[0]
		one["name"] = aCh[1]
		aField = append(aField, one)
	}
	err = modelChannel.BatchAdd(aField)
	if err == nil {
		global.JsonRet(ctx, global.SUCCESS, "", data)
	}else{
		global.JsonRet(ctx, global.ERROR, "", data)
	}
	return
}


