package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var defaultLogLevel = zapcore.ErrorLevel

var Logger *zap.Logger

func init(){
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	path,err := GetGlobalConfig().GetConfigByPath("logger.serverlog")
	if err != nil{
		panic("config logger.serverlog get error: "+err.Error())
	}

	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)


	fileWriteSyncer := zapcore.AddSync(file)
	core := zapcore.NewTee(
		//zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, GetLoggerLevel()),
	)
	Logger = zap.New(core)
}

func getLoggerLevel(level string) zapcore.Level {
	switch level{
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	case "warn":
		return zapcore.WarnLevel
	}
	return defaultLogLevel
}

func GetLoggerLevel() zapcore.Level {
	level,err := GetGlobalConfig().GetConfigByPath("logger.level")
	if err != nil{
		panic("logger.level get error : "+err.Error())
	}
	return getLoggerLevel(level)
}