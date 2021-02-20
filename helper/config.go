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

	CROSEnabled bool `yaml:"cros_enabled"`

	ManageRouter string `yaml:"manage_router"`
	AdminRoot    string `yaml:"admin_root"`
	AdminPass    string `yaml:"admin_pass"`

	SMTPEnabled  bool   `yaml:"smtp_enabled"`
	SMTPHost     string `yaml:"smtp_host"`
	SMTPPort     int    `yaml:"smtp_port"`
	SMTPUsername string `yaml:"smtp_username"`
	SMTPPassword string `yaml:"smtp_password"`
	SMTPForm     string `yaml:"smtp_form"`
	SMTPTo       string `yaml:"smtp_to"`

	SensitivePath string `yaml:"sensitive_path"`
	IPBlockPath   string `yaml:"ip_block_path"`
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

	config.MysqlHost = "127.0.0.1:3306"
	config.MysqlUsr = "root"
	config.MysqlPwd = "password"
	config.MysqlDB = "comment"

	config.ManageRouter = "manage_router"
	config.MysqlUsr = "username"
	config.MysqlPwd = "password"

	config.CROSEnabled = false

	config.AdminRoot = "root"
	config.AdminPass = "pass"

	config.SMTPEnabled = false
	config.SMTPHost = "smtp_host"
	config.SMTPPort = 465
	config.SMTPUsername = "smtp_username"
	config.SMTPPassword = "smtp_password"
	config.SMTPForm = "smtp_form"
	config.SMTPTo = "smtp_to"

	config.SensitivePath = "./config/sensitive.txt"
	config.IPBlockPath = "./config/block_ip.txt"
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
