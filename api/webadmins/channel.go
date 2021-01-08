package webadmins

import (
	"bufio"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"io"
	"my_gin/models"
	"my_gin/pkg/file"
	"my_gin/pkg/global"
	"my_gin/pkg/util"
	"strconv"
)

type channel struct {
}
func NewChannel() *channel{
	return &channel{}
}

func (this *channel) List(ctx *gin.Context){
	modelChannel := models.NewModelChannel()
	size := com.StrTo(ctx.DefaultPostForm("size", "20")).MustInt()
	page := com.StrTo(ctx.DefaultPostForm("page", "1")).MustInt()
	list := modelChannel.GetChannelList(page, size)
	data := make(map[string]interface{})
	data["list"] = list
	data["count"] = modelChannel.GetChannelCount()
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}

func (this *channel) Modify(ctx *gin.Context){
	data := make(map[string]interface{})
	param := make(map[string]interface{})
	modelChannel := models.NewModelChannel()
	id := com.StrTo(ctx.PostForm("id")).MustInt()
	param["name"] = com.StrTo(ctx.PostForm("name")).String()
	param["channelId"] = com.StrTo(ctx.PostForm("channelId")).String()
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
	param["channelId"] = com.StrTo(ctx.PostForm("channelId")).String()

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
		one["channelId"] = aCh[0]
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

//导出所有的渠道信息
func (this *channel) Export(ctx *gin.Context){
	data := make(map[string]interface{})
	modelChannel := models.NewModelChannel()

	channels := modelChannel.GetAllChannel()

	var csvData [][]string
	//csvData := make([][]string, len(channels))//先确定多少数量
	i := 0
	for _, items := range channels{
		tmp := make([]string, 4)
		tmp[0] = strconv.Itoa(items.ID)
		tmp[1] = items.Name
		tmp[2] = items.ChannelId
		tmp[3] = items.Datetime
		//csvData[i] = tmp
		csvData = append(csvData, tmp)//最佳数组
		i++
	}
	data["allList"] = csvData
	err := file.ExportToCsv("channel.csv", csvData)
	if err != nil {
		global.JsonRet(ctx, global.ERROR, "", data)
		return
	}
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}
