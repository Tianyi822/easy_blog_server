package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type CorsConfig struct {
	Origins []string `yaml:"origins"`
	Headers []string `yaml:"headers"`
	Methods []string `yaml:"methods"`
}

type serverConfig struct {
	Port                uint16     `yaml:"port"`
	TokenKey            string     `yaml:"token_key"`
	TokenExpireDuration int        `yaml:"token_expire_duration"` // token 过期时间，单位：天
	Cors                CorsConfig `yaml:"cors"`
}

type sqliteConfig struct {
	DataPath string `yaml:"path"`
	DB       string `yaml:"database"`
	MaxOpen  int    `yaml:"max_open"`
	MaxIdle  int    `yaml:"max_idle"`
}

type loggerConfig struct {
	Level      string `yaml:"level"`
	Path       string `yaml:"path"`
	MaxAge     int    `yaml:"max_age"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}

type projectConfig struct {
	Server serverConfig `yaml:"server"`
	Sqlite sqliteConfig `yaml:"sqlite"`
	Logger loggerConfig `yaml:"logger"`
}

var loadConfigLock sync.Once

var ServerConf = new(serverConfig)
var SqliteConf = new(sqliteConfig)
var LoggerConf = new(loggerConfig)

func LoadConfig(configFilePath string) {
	loadConfigLock.Do(
		func() {
			// 读取配置文件内容
			data, err := os.ReadFile(configFilePath)
			if err != nil {
				msg := fmt.Sprintf("Load config file error: %v", err)
				panic(msg)
			}

			// 解析配置文件内容
			conf := &projectConfig{}
			err = yaml.Unmarshal(data, &conf)
			if err != nil {
				panic("Reflect config to Struct error")
			}

			ServerConf = &conf.Server
			SqliteConf = &conf.Sqlite
			LoggerConf = &conf.Logger
		},
	)
}
