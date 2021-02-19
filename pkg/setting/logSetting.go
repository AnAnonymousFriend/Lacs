package setting

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
	"fmt"
)

func LogSetup()  *zap.SugaredLogger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder,writeSyncer,zapcore.DebugLevel)

	logger :=zap.New(core)
	var sugarLogger = logger.Sugar()
	return sugarLogger
}

func getEncoder() zapcore.Encoder  {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer  {
	fileName := getLogFileName()
	file,_ := os.Create(fileName)
	return zapcore.AddSync(file)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", "runtime/", "logs/")
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		"../logs/",
		time.Now().Format("20060102"),
		"log",
	)
}
