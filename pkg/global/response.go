package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonRet(ctx *gin.Context, code int, msg string, data interface{}){
	ret := gin.H{}
	ret["code"] = code
	if msg != ""{
		ret["msg"] =  msg
	}else{
		ret["msg"] =  GetMsg(code)
	}

	if data != nil {
		ret["data"] = data
	}

	ctx.JSON(http.StatusOK, ret)
	return
}
