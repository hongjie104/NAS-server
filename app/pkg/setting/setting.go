package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg a
	Cfg *ini.File
	// RunMode s
	RunMode string
	// HTTPPort a
	HTTPPort string
	// ReadTimeout a
	ReadTimeout time.Duration
	// WriteTimeout a
	WriteTimeout time.Duration

	// DbHost a
	DbHost string
	// DbName a
	DbName string

	// PageSize a
	PageSize int
	// JwtSecret a
	JwtSecret string
)

func init() {
	var err error
	// Cfg, err = ini.Load("../../config/app.ini")
	Cfg, err = ini.Load("app/config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}

	loadBase()
	loadServer()
	loadDB()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustString(":8080")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)@#!$@#@#!")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func loadDB() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	DbHost = sec.Key("HOST").MustString("")
	DbName = sec.Key("DB").MustString("go-test1")
}
