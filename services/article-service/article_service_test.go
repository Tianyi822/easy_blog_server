package article_service

import (
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/db"
	"github.com/Tianyi822/easy_blog_server/logger"
	"github.com/Tianyi822/easy_blog_server/model/dto"
	"testing"
)

func init() {
	// 加载配置文件
	config.LoadConfig("../../resources/config/service-test-config.yaml")
	// 初始化 Logger 组件
	logger.InitLogger()
	// 连接数据库
	db.ConnectSqlite()
}

func TestAddArticle(t *testing.T) {
	atDto := &dto.ArticleDto{
		Aid:        "1",
		Title:      "12345",
		BriefIntro: "12345",
		Url:        "12345",
	}

	AddArticle(atDto)
}
