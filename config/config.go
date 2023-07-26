package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

const (
	File = "./config.yml"
)

var GlobalConfig *Config

type Config struct {
	Notion       map[string]string `yaml:"notion" json:"notion"`
	System       map[string]string `yaml:"system" json:"system"`
	SqlCondition map[string]string `yaml:"sql_condition" json:"sql_condition"`
}

func GetConfig() *Config {
	yamlFile, err := ioutil.ReadFile(File)
	if err != nil {
		log.Println(err.Error())
	}
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		log.Println(err.Error())
	}
	return GlobalConfig
}

func InitConfig() {
	GetConfig()
}
