package util

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var resp = Response{}

type configStruct struct {
	Installed  bool   `yaml:"installed" json:"installed"`
	SiteName   string `yaml:"site_name" json:"site_name"`
	SiteUrl    string `yaml:"site_url" json:"site_url"`
	ServerPort int    `yaml:"server_port" json:"server_port" `

	DBApp  string `yaml:"db_app" json:"db_app"`
	DBHost string `yaml:"db_host" json:"db_host"`
	DBPort int    `yaml:"db_port" json:"db_port"`
	DBUsr  string `yaml:"db_usr" json:"db_usr"`
	DBPwd  string `yaml:"db_pwd" json:"db_pwd"`
	DBLib  string `yaml:"db_lib" json:"db_lib"`

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

	SendCloudEnabled      bool   `yaml:"send_cloud_enabled" json:"send_cloud_enabled"`
	SendCloudAPIUser      string `yaml:"send_cloud_api_user" json:"send_cloud_api_user"`
	SendCloudAPIKey       string `yaml:"send_cloud_api_key" json:"send_cloud_api_key"`
	SendCloudFrom         string `yaml:"send_cloud_from" json:"send_cloud_from"`
	SendCloudTemplateName string `yaml:"send_cloud_template_name" json:"send_cloud_template_name"`
}

// Config 配置内容
var Config configStruct

var configPath string = "./config.yaml"

// Setup 装载配置
func init() {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		Config.Installed = false
		return
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		Config.Installed = false
		log.Error(err.Error())
		return
	}
	Config.Installed = true
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*")

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// SaveConfigFile 保存配置文件
func SaveConfigFile(c *gin.Context) {
	var conf configStruct
	err := c.BindJSON(&conf)
	if err != nil {
		log.Errorf("配置生成入参错误：%s", err.Error())
		resp.Error(c, ResponseParamError, "入参错误")
		return
	}
	if conf.SensitiveEnabled == true {
		conf.SensitivePath = "./sensitive.txt"
		err = downloadFile(conf.SensitivePath, "https://cdn.jsdelivr.net/gh/importcjj/sensitive@latest/dict/dict.txt")
		if err != nil {
			log.Errorf("下载敏感字典错误：%s", err.Error())
			resp.Error(c, ResponseServerError, "下载敏感字典错误")
			return
		}
	}
	conf.JWTEncrypt = randStringRunes(26)
	conf.Installed = true
	yamlFile, err := yaml.Marshal(conf)
	if err != nil {
		log.Errorf("配置生成错误：%s", err.Error())
		resp.Error(c, ResponseServerError, "生成错误")
		return
	}
	err = ioutil.WriteFile(configPath, yamlFile, 0666)
	if err != nil {
		log.Errorf("配置保存错误：%s", err.Error())
		resp.Error(c, ResponseServerError, "保存错误")
		return
	}
	resp.Success(c, true)
}

// downloadFile 下载文件
func downloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
