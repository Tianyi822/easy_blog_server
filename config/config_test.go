package config

import (
	"testing"
)

func init() {
	// 加载配置文件
	LoadConfig("../resources/config/config-test.yaml")
}

func TestLoadConfig(t *testing.T) {
	println(OssConf.Endpoint)
	println(OssConf.AccessKeyId)
	println(OssConf.AccessKeySecret)
	println(OssConf.Bucket)
}
