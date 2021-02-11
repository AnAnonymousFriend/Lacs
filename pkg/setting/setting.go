package setting

import (
	"github.com/go-ini/ini"
	"time"
)

type App struct {
	Host string
}

type MongoDB struct {
	Host string
	UserName string
	Password string
	DbName string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}



var cfg *ini.File

var AppSetting = &App{}
var RedisSetting = &Redis{}
var MongoDBSetting = &MongoDB{}
func Setup()  {
	globalSetup()
	CasbinSetting()
}

func globalSetup()  {
	var err error
	cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		println("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("mongo", MongoDBSetting)
	mapTo("redis", RedisSetting)
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}


func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		println("Cfg.MapTo %s err: %v", section, err)
	}
}