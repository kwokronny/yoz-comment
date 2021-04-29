package dao

import (
	"YozComment/util"
	"fmt"
	"time"

	"gorm.io/gorm"
	// Register mysql server

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

// DB is gorm instance Lib
var DB *gorm.DB

// Model 通用数据表结构
type Model struct {
	ID        uint      `gorm:"AUTOINCREMENT;primary_key;unique;type:int;size:32;comment:ID" json:"id" form:"id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:timestamp;comment:发表日期" json:"createdAt" time_utc:"1"`
	DeletedAt *time.Time
}

func init() {
	var err error
	if util.Config.Installed == false {
		return
	}
	config := util.Config
	var dbSource gorm.Dialector
	if config.DBApp == "mysql" {
		var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=True", config.DBUsr, config.DBPwd, config.DBHost, config.DBPort, config.DBLib)
		dbSource = mysql.New(mysql.Config{
			DSN: dsn,
		})
	} else if config.DBApp == "postgresql" {
		var dsn = fmt.Sprintf("host=%s user=%s password=%s DB.name=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUsr, config.DBPwd, config.DBLib, config.DBPort)
		dbSource = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	} else if config.DBApp == "sqlite" {
		var dbName = config.DBLib + ".db"
		dbSource = sqlite.Open(dbName)
	}
	DB, err = gorm.Open(dbSource, &gorm.Config{})
	if err != nil {
		log.Panicln("GORM Error:", err.Error())
	}
}
