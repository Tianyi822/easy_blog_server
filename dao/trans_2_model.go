package dao

import (
	"github.com/Tianyi822/easy_blog_server/dao/article"
	"github.com/Tianyi822/easy_blog_server/model/dto"
)

func TransArticleDao2Dto(ad *article.ArticleDao) *dto.ArticleDto {
	return &dto.ArticleDto{
		Aid:        ad.Aid,
		Title:      ad.Title,
		BriefIntro: ad.BriefIntro,
		Url:        ad.Url,
	}
}

func TransArticleDto2Dao(atDto *dto.ArticleDto) *article.ArticleDao {
	return &article.ArticleDao{
		Aid:        atDto.Aid,
		Title:      atDto.Title,
		BriefIntro: atDto.BriefIntro,
		Url:        atDto.Url,
	}
}
