package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type CorsConfig struct {
	Origins []string `yaml:"origins"`
	Headers []string `yaml:"headers"`
	Methods []string `yaml:"methods"`
}

type ServerConfig struct {
	Port                uint16     `yaml:"port"`
	RecommendKey        string     `yaml:"recommend_key"`
	TokenKey            string     `yaml:"token_key"`
	TokenExpireDuration int        `yaml:"token_expire_duration"` // token 过期时间，单位：天
	Cors                CorsConfig `yaml:"cors"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	DB       string `yaml:"database"`
	MaxOpen  int    `yaml:"max_open"`
	MaxIdle  int    `yaml:"max_idle"`
}

type RedisConfig struct {
	Host      string `yaml:"host"`
	Port      uint16 `yaml:"port"`
	Password  string `yaml:"password"`
	DbNum     int    `yaml:"db_num"`
	PoolSize  int    `yaml:"pool_size"`
	KeyPrefix string `yaml:"key_prefix"`
}

type LoggerConfig struct {
	Level      string `yaml:"level"`
	Path       string `yaml:"path"`
	MaxAge     int    `yaml:"max_age"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}

type ProjectConfig struct {
	Server ServerConfig `yaml:"server"`
	MySQL  MySQLConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	Logger LoggerConfig `yaml:"logger"`
}

var ServerConf = new(ServerConfig)
var MySQLConf = new(MySQLConfig)
var RedisConf = new(RedisConfig)
var LoggerConf = new(LoggerConfig)

func LoadConfig(configFilePath string) {

	// 读取配置文件内容
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		msg := fmt.Sprintf("Load config file error: %v", err)
		panic(msg)
	}

	// 解析配置文件内容
	conf := &ProjectConfig{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic("Reflect config to Struct error")
	}

	ServerConf = &conf.Server
	MySQLConf = &conf.MySQL
	RedisConf = &conf.Redis
	LoggerConf = &conf.Logger
}
