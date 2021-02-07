package setting

import "github.com/go-ini/ini"

type App struct {
	Host string
}

type MongoDB struct {
	Host string
	UserName string
	Password string
	DbName string
}

var cfg *ini.File
var MongoDBSetting = &MongoDB{}
var AppSetting = &App{}

func Setup()  {
	var err error
	cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		println("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}


	mapTo("app", AppSetting)
	mapTo("mongo", MongoDBSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		println("Cfg.MapTo %s err: %v", section, err)
	}
}