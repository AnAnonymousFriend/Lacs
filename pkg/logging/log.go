package logging

import (
	"Lacs/pkg/setting"
	"fmt"
)

func Info(v ...interface{}) {
	 setting.Logger.Info(v)
}

func Debug(v ...interface{}) {
	setting.Logger.Debug(v)
}

func Error(v ...interface{}) {
	fmt.Println("执行Log.Error 方法")
	setting.Logger.Error(v)
}
