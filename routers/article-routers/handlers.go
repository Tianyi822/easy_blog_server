package article_routers

import (
	"github.com/Tianyi822/easy_blog_server/logger"
	"github.com/Tianyi822/easy_blog_server/routers"
	"github.com/gin-gonic/gin"
)

func addArticle(c *gin.Context) {
	// TODO
	logger.Info("添加文章")
	routers.Ok(c, "添加文章成功", nil)
}
