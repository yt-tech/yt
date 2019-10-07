package client

import (
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func initLogger(logpath string, loglevel string) *zap.Logger {
	LogPath, LogLevel := logConfig()
	mlog.Println(LogPath, LogLevel)
	logger = initLogger(LogPath, LogLevel)

	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 30,      // 最多保留300个备份
		MaxAge:     7,       // days
		Compress:   false,   // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = timeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core, zap.AddCaller())
	return logger
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// func newEncoderConfig() zapcore.EncoderConfig {
// 	return zapcore.EncoderConfig{
// 		// Keys can be anything except the empty string.
// 		TimeKey:        "T",
// 		LevelKey:       "L",
// 		NameKey:        "N",
// 		CallerKey:      "C",
// 		MessageKey:     "M",
// 		StacktraceKey:  "S",
// 		LineEnding:     zapcore.DefaultLineEnding,
// 		EncodeLevel:    zapcore.CapitalLevelEncoder,
// 		EncodeTime:     timeEncoder,
// 		EncodeDuration: zapcore.StringDurationEncoder,
// 		EncodeCaller:   zapcore.ShortCallerEncoder,
// 	}
// }
