package helper

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	ServerPort string `yaml:"server_port"`

	LogFilePath string `yaml:"log_file_path"`

	MysqlHost string `yaml:"mysql_host"`
	MysqlUsr  string `yaml:"mysql_usr"`
	MysqlPwd  string `yaml:"mysql_pwd"`
	MysqlDB   string `yaml:"mysql_db"`

	IPBlockPath string `yaml:"ip_block_path"`
}

// Config 配置内容
var Config conf

var configPath string = "./config/config.yaml"

// Setup 装载配置
func init() {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		Config, err = initConfigFile()
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		err = yaml.Unmarshal(yamlFile, &Config)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func initConfigFile() (config conf, err error) {
	config.ServerPort = "80"

	config.LogFilePath = "./logs/"

	config.MysqlHost = "url"
	config.MysqlUsr = "username"
	config.MysqlPwd = "password"
	config.MysqlDB = "comment"

	config.IPBlockPath = "./config/block_ip"
	yamlFile, err := yaml.Marshal(config)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(configPath, yamlFile, 0666)
	if err != nil {
		return
	}
	return
}
