package oss

import (
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/logger"
	"testing"
)

func init() {
	config.LoadConfig("../resources/config/config-test.yaml")
	logger.InitLogger()
	ConnectOss()
}

func TestConnectOss(t *testing.T) {
	println("测试连接 OSS 成功")
}
