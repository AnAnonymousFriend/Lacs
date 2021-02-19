package setting

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func LogSetup()  *zap.SugaredLogger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()


	core := zapcore.NewCore(encoder,writeSyncer,zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	var sugarLogger = logger.Sugar()
	return sugarLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {

	fileName := getLogFileName()
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    LogSetting.LogFileMaxSize,
		MaxBackups: LogSetting.LogMaxBackups,
		MaxAge:     LogSetting.LogMaxAge,
		Compress:   LogSetting.LogCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		"./logs/",
		time.Now().Format("20060102"),
		"log",
	)
}


//func getLogWriter() zapcore.WriteSyncer  {
//	fileName := getLogFileName()
//	file,_ := os.Create(fileName)
//	return zapcore.AddSync(file)
//}

