package db

import (
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/logger"
	"testing"
)

func init() {
	// 加载配置文件
	config.LoadConfig("../resources/config/config-test.yaml")
	// 初始化 Logger 组件
	logger.InitLogger()
}

func TestConnectSqlite(t *testing.T) {
	ConnectSqlite()
}
