package dao

import (
	"YozComment/util"
	"fmt"
	"time"

	"gorm.io/gorm"
	// Register mysql server

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
)

// DB is gorm instance Lib
var DB *gorm.DB

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      dao.Model
//    }
type Model struct {
	ID        uint      `gorm:"AUTOINCREMENT;primary_key;unique;type:int;size:32;comment:ID" json:"id" form:"id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:timestamp;comment:发表日期" json:"createdAt" time_utc:"1"`
	DeletedAt *time.Time
}

func init() {
	var err error
	config := util.Config
	var uri string = fmt.Sprintf("%s:%s@tcp(%s)/%s?loc=Local&parseTime=True", config.MysqlUsr, config.MysqlPwd, config.MysqlHost, config.MysqlDB)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: uri,
	}), &gorm.Config{})
	if err != nil {
		log.Panicln("GORM Error:", err.Error())
	}
}
