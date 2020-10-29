package webadmins

import (
	"github.com/gin-gonic/gin"
	"my_gin/pkg/global"
	"my_gin/pkg/logger"
)

type system struct {
}


func NewSystem() *system{
	return &system{}
}

func (this *system) ApiTest(ctx *gin.Context){
	data := make(map[string]interface{})
	logger.Info("GetAuth", data)
	global.JsonRet(ctx, global.SUCCESS, "", data)
	return
}