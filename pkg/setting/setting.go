package setting

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-ini/ini"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type App struct {
	Host string
}

type Log struct {
	Filename string
	LogSaveName string
	LogFileMaxSize int
	LogMaxBackups int
	LogMaxAge	int
	LogCompress	bool
}

type MongoDB struct {
	Host string
	UserName string
	Password string
	DbName string
	MaxConn uint64
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var cfg *ini.File
var CabinEnforcer *casbin.SyncedEnforcer
var Logger *zap.SugaredLogger
var AppSetting = &App{}
var LogSetting = &Log{}
var RedisSetting = &Redis{}
var MongoDBSetting = &MongoDB{}
var MongoDataBase *mongo.Database

func Setup()  {
	globalSetup()


	CasbinSetup()

	Logger = LogSetup()
	if Logger == nil {
		println("Logger 对象为空")
	}

	MongoDataBase = MongoDBSetup()
}

func globalSetup()  {
	var err error
	cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		println("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("mongodb", MongoDBSetting)
	mapTo("redis", RedisSetting)
	mapTo("log", LogSetting)
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		println("Cfg.MapTo %s err: %v", section, err)
	}
}
