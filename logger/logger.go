package logger

import (
	"github.com/Tianyi822/easy_blog_server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"sync"
)

var Log *zap.SugaredLogger

var loggerLock sync.Once

func Info(format string, args ...any) {
	if Log == nil {
		InitLogger()
	}
	Log.Infof(format, args...)
}

func Debug(format string, args ...any) {
	if Log == nil {
		InitLogger()
	}
	Log.Debugf(format, args...)
}

func Error(format string, args ...any) {
	if Log == nil {
		InitLogger()
	}
	Log.Errorf(format, args...)
}

func Warn(format string, args ...any) {
	if Log == nil {
		InitLogger()
	}
	Log.Warnf(format, args...)
}

func Panic(format string, args ...any) {
	if Log == nil {
		InitLogger()
	}
	Log.Panicf(format, args...)
}

func InitLogger() {
	loggerLock.Do(
		func() {
			loggerConf := config.LoggerConf

			// 日志基本配置
			conf := &lumberjack.Logger{
				Filename:   loggerConf.Path,       // 日志文件路径，默认 os.TempDir()
				MaxSize:    loggerConf.MaxSize,    // 每个日志文件保存10M，默认 100M
				MaxBackups: loggerConf.MaxBackups, // 保留30个备份，默认不限
				MaxAge:     loggerConf.MaxAge,     // 保留7天，默认不限
				Compress:   loggerConf.Compress,   // 是否压缩，默认不压缩
			}

			encoderConfig := zapcore.EncoderConfig{
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "line num",
				MessageKey:     "msg",
				StacktraceKey:  "stack trace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
				EncodeName:     zapcore.FullNameEncoder,
			}

			// 设置日志级别
			var level zapcore.Level
			switch strings.ToLower(loggerConf.Level) {
			case "debug":
				level = zap.DebugLevel
			case "info":
				level = zap.InfoLevel
			case "error":
				level = zap.ErrorLevel
			default:
				level = zap.InfoLevel
			}

			var write = zapcore.NewMultiWriteSyncer(
				// 添加控制台打印
				zapcore.AddSync(os.Stdout),
				// 添加文件输出
				zapcore.AddSync(conf),
			)

			core := zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				write,
				level,
			)

			// 开启文件及行号
			development := zap.Development()

			// 开启开发模式，堆栈跟踪
			caller := zap.AddCaller()
			callerSkip := zap.AddCallerSkip(1)

			// 构造日志
			Log = zap.New(core, caller, callerSkip, development).Sugar()
		},
	)
}
