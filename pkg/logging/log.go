package logging

import "Lacs/pkg/setting"

func Info(v ...interface{}) {
	 setting.Logger.Info(v)
}

func Debug(v ...interface{}) {
	setting.Logger.Debug(v)
}

func Error(v ...interface{}) {
	setting.Logger.Error(v)
}
