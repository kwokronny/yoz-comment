package util

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var resp = Response{}

type configStruct struct {
	SiteName   string `yaml:"site_name" json:"site_name"`
	SiteUrl    string `yaml:"site_url" json:"site_url"`
	ServerPort int    `yaml:"server_port" json:"server_port" `

	MysqlHost string `yaml:"mysql_host" json:"mysql_host"`
	MysqlUsr  string `yaml:"mysql_usr" json:"mysql_usr"`
	MysqlPwd  string `yaml:"mysql_pwd" json:"mysql_pwd"`
	MysqlDB   string `yaml:"mysql_db" json:"mysql_db"`

	CROSEnabled bool `yaml:"cros_enabled" json:"cros_enabled"`

	ManageRouter string `yaml:"manage_router" json:"manage_router"`
	JWTEncrypt   string `yaml:"jwt_encrypt" json:"jwt_encrypt"`
	AdminRoot    string `yaml:"admin_root" json:"admin_root"`
	AdminPass    string `yaml:"admin_pass" json:"admin_pass"`

	SMTPEnabled  bool   `yaml:"smtp_enabled" json:"smtp_enabled"`
	SMTPHost     string `yaml:"smtp_host" json:"smtp_host"`
	SMTPPort     int    `yaml:"smtp_port" json:"smtp_port"`
	SMTPUsername string `yaml:"smtp_username" json:"smtp_username"`
	SMTPPassword string `yaml:"smtp_password" json:"smtp_password"`
	SMTPTo       string `yaml:"smtp_to" json:"smtp_to"`

	SensitiveEnabled bool   `yaml:"sensitive_enabled" json:"sensitive_enabled"`
	SensitivePath    string `yaml:"sensitive_path"`
}

// Config 配置内容
var Config configStruct

var configPath string = "./config/config.yaml"

// Setup 装载配置
func init() {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, &Config)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// SaveConfigFile is save config
func SaveConfigFile(c *gin.Context) {
	var conf configStruct
	err := c.BindJSON(&conf)
	if err != nil {
		fmt.Println(err.Error())
		resp.Error(c, ResponseParamError, "入参错误")
		return
	}
	conf.SensitivePath = "./config/sensitive.txt"
	conf.JWTEncrypt = randStringRunes(26)
	yamlFile, err := yaml.Marshal(conf)
	if err != nil {
		resp.Error(c, ResponseServerError, "生成错误")
		return
	}
	err = ioutil.WriteFile(configPath, yamlFile, 0666)
	if err != nil {
		resp.Error(c, ResponseServerError, "保存错误")
		return
	}
	resp.Success(c, true)
	exitInstall()
}

func exitInstall() {
	cmd := exec.Command("sh", "serve.sh")
	if err := cmd.Run(); err != nil {
		log.Error(err)
	}
}
