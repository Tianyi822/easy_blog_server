package dao

import (
	"github.com/Tianyi822/easy_blog_server/dao/article"
	"github.com/Tianyi822/easy_blog_server/model/dto"
)

func TransArticleDto2Dao(atDto *dto.ArticleDto) *article.ArticleDao {
	return &article.ArticleDao{
		Title:      atDto.Title,
		BriefIntro: atDto.BriefIntro,
	}
}
