package article_service

import (
	"github.com/Tianyi822/easy_blog_server/dao"
	"github.com/Tianyi822/easy_blog_server/logger"
	"github.com/Tianyi822/easy_blog_server/model/dto"
)

// AddArticle 添加文章
func AddArticle(atDto *dto.ArticleDto) {
	err := dao.TransArticleDto2Dao(atDto).Create()
	if err != nil {
		logger.Warn("添加文章失败: %v", err)
		return
	}
}
