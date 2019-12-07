package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type appConf struct {
	PageSize  int    `yaml:"page_size"`
	JwtSecret string `yaml:"jwt_secret"`
}

type serverConf struct {
	RunMode      string `yaml:"run_mode"`
	HTTPPort     string `yaml:"http_port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	SaveInternal uint   `yaml:"save_internal"`
}

type databaseConf struct {
	HOST string `yaml:"host"`
	DB   string `yaml:"db"`
}

// Conf Conf
type Conf struct {
	APP appConf `yaml:"app"`

	Server serverConf `yaml:"server"`

	Database databaseConf `yaml:"database"`
}

// Config Config
var Config = &Conf{}

func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
