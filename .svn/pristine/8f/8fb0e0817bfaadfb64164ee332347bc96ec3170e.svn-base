package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"my_gin/pkg/setting"
)

func GetPage(c *gin.Context) int{
	ret :=0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		ret = (page - 1) * setting.AppSetting.PageSize
	}
	return ret
}
