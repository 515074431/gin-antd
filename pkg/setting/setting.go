package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	WebDavDir string
	WebDavPrefix string

	PageSize int
	JwtSecret string
	TokenExpireHour int
)

func init()  {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'conf/app.ini': %v",err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

//加载基本
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//加载服务配置
func LoadServer() {
	sec, err :=Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	WebDavDir = sec.Key("WEBDAV_DIR").MustString(".")
	WebDavPrefix = sec.Key("WEBDAV_PREFIX").MustString("/api/v1/dav/")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)

}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil{
		log.Fatal("Fail to get section 'app': %v", err)
	}
	TokenExpireHour = sec.Key("TOKEN_EXPIRE_HOUR").MustInt(24)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)123*#)!456@U#@789*!@!")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
