package article

import (
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/db"
	"github.com/Tianyi822/easy_blog_server/logger"
	"testing"
)

func init() {
	// 加载配置文件
	config.LoadConfig("../../resources/config/dao-test-config.yaml")
	// 初始化 Logger 组件
	logger.InitLogger()
	// 连接数据库
	db.ConnectSqlite()
}

func TestArticleDao_Create(t *testing.T) {
	articleDao := &ArticleDao{
		Aid:        "123456",
		Title:      "12345",
		BriefIntro: "12345",
		Url:        "12345",
		Salt:       "12345",
	}

	err := articleDao.Create()

	if err != nil {
		t.Errorf("创建文章记录失败: %v", err)
	} else {
		t.Logf("创建文章记录成功")
	}
}
