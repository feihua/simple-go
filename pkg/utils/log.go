package utils

import (
	"github.com/natefinch/lumberjack" // Lumberjack进行日志切割归档
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

// Logger 第一步：创建一个Sugared Logger
var Logger *zap.SugaredLogger

// 第二步：将编码器从JSON Encoder更改为普通Encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 为此，我们需要将NewJSONEncoder()更改为NewConsoleEncoder()
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 第三步：使用Lumberjack进行日志切割归档
func getLogWriter() zapcore.WriteSyncer {
	logFileName := "./logs/" + time.Now().Format("2006-01-02") + ".log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName, // 指定日志将写到哪里去
		MaxSize:    1,           // 每次满1M进行切割
		MaxBackups: 5,           // 最多报错5个文件
		MaxAge:     30,          // 文件最多保存30天
		Compress:   false,       // 是否压缩/归档旧文件
	}
	//同时输出文件和控制台
	fileWriter := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(fileWriter)
}

// InitLogger 第四步：重写InitLogger()方法
func InitLogger() {
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	//   zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	/*
	   Encoder:编码器(如何写入日志),我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的
	   WriterSyncer ：指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去
	   Log Level：哪种级别的日志将被写入。
	*/
	// 我们将使用zap.New(…)方法来手动传递所有配置，而不是使用像zap.NewProduction()这样的预置方法来创建logger。
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar() // 实例化全局变量sugarLogger
}
