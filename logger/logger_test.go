package logger

import (
	"github.com/Tianyi822/easy_blog_server/config"
	"testing"
)

func init() {
	// 加载配置文件
	config.LoadConfig("../resources/config/config-test.yaml")
	// 初始化 Logger 组件
	InitLogger()
}

func TestLogger(t *testing.T) {
	Info("info -- test logger component")
	Debug("debug -- test logger component")
	Warn("warn -- test logger component")
}

func TestBigLog(t *testing.T) {
	for {
		Info("info -- test logger component")
		Debug("debug -- test logger component")
		Warn("warn -- test logger component")
	}
}
