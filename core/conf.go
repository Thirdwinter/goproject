package core

import (
	"fmt"
	"log"
	"os"

	"goproject/config"
	"goproject/global"
	"gopkg.in/yaml.v2"
)

// 读取配置文件
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("confing error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init Unmarshal error %s", err)
	}
	//log.Println("log config init success.")
	global.Config = c
	//log.Println(c)
}
