package routers

import (
	"github.com/gin-gonic/gin"
	v1 "my_gin/api/v1"
)


/**
 * 此处都是业务逻辑的接口路由
 */

func initApiRouters(router *gin.Engine){
	initArticleRouter(router)
	initTagRouter(router)
}

//文章路由
func initArticleRouter(router *gin.Engine){
	v1Art := router.Group("/v1/article")
	{
		apiArticle := v1.NewApiArticle()
		v1Art.POST("/info", apiArticle.GetArticleInfo)
		v1Art.POST("/getList", apiArticle.GetArticleList)
		v1Art.POST("/add", apiArticle.AddArticle)
	}
}

//标签路由
func initTagRouter(router *gin.Engine){
	v1Tags := router.Group("/v1/tags")
	{
		apiTag := v1.NewApiTag()
		v1Tags.POST("/get", apiTag.GetTags)
		v1Tags.POST("/add", apiTag.AddTags)
		v1Tags.POST("/update", apiTag.EditTags)
		v1Tags.POST("/del", apiTag.DelTags)
	}
}