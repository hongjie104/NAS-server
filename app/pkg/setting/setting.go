package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type appConf struct {
	PageSize  int8   `yaml:"page_size"`
	JwtSecret string `yaml:"jwt_secret"`
}

type serverConf struct {
	HTTPPort     string `yaml:"http_port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type databaseConf struct {
	HOST string `yaml:"host"`
	DB   string `yaml:"db"`
}

// Conf Conf
type Conf struct {
	RunMode string `yaml:"run_mode"`

	APP appConf `yaml:"app"`

	Server serverConf `yaml:"server"`

	Database databaseConf `yaml:"database"`
}

// Config Config
var Config = &Conf{}

func init() {
	yamlFile, err := ioutil.ReadFile("app/config/app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
