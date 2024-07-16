package article_routers

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	// 添加文章路由
	articleGroup := e.Group("/article")

	articleGroup.POST("/add", addArticle)
}
