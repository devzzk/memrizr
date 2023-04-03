package core

import (
	"fmt"
	"github.com/devzzk/memrizr/global"
	"gopkg.in/yaml.v3"
	"log"
	"os"

	"github.com/devzzk/memrizr/config"
)

// InitConf 读取配置文件的数据
func InitConf() {
	ConfigFile := "settings.yaml"
	c := &config.Config{}

	yamlConf, err := os.ReadFile(ConfigFile)

	if err != nil {
		panic(fmt.Errorf("getYamlConf error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)

	if err != nil {
		log.Fatalf("config Init Unmarshal:", err)
	}

	log.Println("config yaml File load Init success.")
	global.Config = c
}
