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

type ServerConfig struct {
    Port                uint16     `yaml:"port"`
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
    Logger LoggerConfig `yaml:"logger"`
}

var loadConfigLock sync.Once

var ServerConf = new(ServerConfig)
var MySQLConf = new(MySQLConfig)
var LoggerConf = new(LoggerConfig)

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
            conf := &ProjectConfig{}
            err = yaml.Unmarshal(data, &conf)
            if err != nil {
                panic("Reflect config to Struct error")
            }

            ServerConf = &conf.Server
            MySQLConf = &conf.MySQL
            LoggerConf = &conf.Logger
        },
    )
}
