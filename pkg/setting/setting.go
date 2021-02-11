package setting

import (
	"github.com/go-ini/ini"
	"github.com/casbin/casbin/v2"
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
var CasbinEnforcer *casbin.SyncedEnforcer
var AppSetting = &App{}
var RedisSetting = &Redis{}
var MongoDBSetting = &MongoDB{}
func Setup()  {
	globalSetup()
	casbinSetup()
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

func casbinSetup()  {
	ef, err := casbin.NewSyncedEnforcer("../conf/acl_simple_model.conf", "../conf/acl_simple_policy.csv")
	if err!=nil {
		println(err)
	}
	CasbinEnforcer = ef
}