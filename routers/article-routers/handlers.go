package article_routers

import (
	"github.com/Tianyi822/easy_blog_server/routers"
	"github.com/Tianyi822/easy_blog_server/services/article-service"
	"github.com/gin-gonic/gin"
)

func addArticle(c *gin.Context) {
	// 获取并解析 RawData 数据
	articleDto, err := routers.GetArticleDataFromRawData(c)
	if err != nil {
		return
	}

	// 保存文章副本到本地，并上传到 OSS
	article_service.AddArticle(articleDto)

	routers.Ok(c, "添加文章成功", nil)
}
