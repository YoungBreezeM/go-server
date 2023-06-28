package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host   string `yaml:"host"`
	Port   uint16 `yaml:"port"`
	Log    Log    `yaml:"log"`
	MySql  MySql  `yaml:"mysql"`
	WeChat WeChat `yaml:"wechat"`
}

var Cfg Config

func init() {
	Cfg = Config{
		Host:  "127.0.0.1",
		Port:  8080,
		MySql: MySql{},
	}
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	if err = yaml.Unmarshal(yamlFile, &Cfg); err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}
}
