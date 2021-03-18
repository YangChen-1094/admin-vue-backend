package webadmins

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"my_gin/models"
	"my_gin/pkg/global"
	"my_gin/pkg/util"
	"strconv"
	"strings"
)

type item struct {
}
func NewItem() *item{
	return &item{}
}

func (this *item) List(ctx *gin.Context){
	modelItem := models.NewModelItem()
	size := com.StrTo(ctx.DefaultPostForm("size", "20")).MustInt()
	page := com.StrTo(ctx.DefaultPostForm("page", "1")).MustInt()
	itemList := modelItem.GetItemList(page, size)
	data := make(map[string]interface{})
	aType := modelItem.GetItemType()

	data["page"] = page
	data["size"] = size
	data["total"] = modelItem.GetItemCount()
	data["list"] = itemList
	data["aType"] = aType
	data["usetype"] = models.UseType
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}

func (this *item) Modify(ctx *gin.Context){
	data := make(map[string]interface{})
	param := make(map[string]interface{})
	modelItem := models.NewModelItem()
	id := com.StrTo(ctx.DefaultPostForm("id", "0")).MustInt()
	itemid := com.StrTo(ctx.PostForm("itemid")).MustInt()
	itemname := com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("itemname", ""))).String()
	if itemid == 0 || itemname == ""{
		global.JsonRet(ctx, global.ERROR, "参数为空", data)
		return
	}

	checkUseType := ctx.PostFormMap("usetype")
	fmt.Println(checkUseType)
	var aType []int
	for _, typeId := range checkUseType {
		id, _ := strconv.Atoi(typeId)
		_, exists := models.UseType[id]
		if exists {
			aType = append(aType, id)
		}
	}
	var sType string
	if len(aType) > 0 {
		sType = util.Implode(aType, "#")
	}else{
		sType = ""
	}

	fmt.Println("aType:", aType, "; sType:",sType)
	param["productID"] = ""
	param["usetype"] = sType
	param["itemid"] = itemid
	param["itemname"] = itemname
	param["type"] = com.StrTo(ctx.DefaultPostForm("type", "0")).MustInt()
	param["descr"] = com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("descr", ""))).String()
	param["buttonUrl"] = com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("buttonUrl", ""))).String()
	param["bannerUrl"] = com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("bannerUrl", ""))).String()
	param["buttonDesc"] = com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("buttonDesc", ""))).String()

	imgUrl := com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("imgUrl", ""))).String()
	imgUrlMedium := com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("imgUrlMedium", ""))).String()
	imgUrlLow := com.StrTo(strings.TrimSpace(ctx.DefaultPostForm("imgUrlLow", ""))).String()
	arr := []string{imgUrl, imgUrlMedium, imgUrlLow}
	imgUrl = util.Implode(arr, "#")
	param["imgUrl"] = imgUrl
	param["price"] = com.StrTo(ctx.DefaultPostForm("price", "0")).String()

	var err error
	if id == 0 {
		err = modelItem.Insert(param)
	}else{
		err = modelItem.Modify(id, param)
	}
	if err != nil {
		global.JsonRet(ctx, global.ERROR, err.Error(), data)
	}else{
		global.JsonRet(ctx, global.SUCCESS, "", data)
	}
	return
}

func (this *item) Delete(ctx *gin.Context){
	data := make(map[string]interface{})
	modelItem := models.NewModelItem()
	id := com.StrTo(ctx.PostForm("id")).MustInt()
	del := modelItem.Del(id)
	if del {
		global.JsonRet(ctx, global.SUCCESS, "", data)
	}else{
		global.JsonRet(ctx, global.ERROR, "", data)
	}
	return
}

func (this *item) SetRedis(ctx *gin.Context){
	data := make(map[string]interface{})
	param := make(map[string]interface{})
	modelItem := models.NewModelItem()
	itemList := modelItem.GetAllItem()
	param["name"] = com.StrTo(ctx.PostForm("name")).String()
	param["channelId"] = com.StrTo(ctx.PostForm("channelId")).String()

	var itemJsonString = make(map[string]interface{})
	i := 0
	for _, oneItem := range itemList{
		jsonVal, _ := json.Marshal(oneItem)
		stringNum := strconv.Itoa(i)
		itemJsonString[stringNum] = string(jsonVal)
		i++
	}
	models.RedisMgr.RdsClient.HMSet("item-test", itemJsonString)
	if true {
		global.JsonRet(ctx, global.SUCCESS, "", data)
	}else{
		global.JsonRet(ctx, global.ERROR, "", data)
	}
	return
}
