package dao

import (
	"fmt"
	"kwok-comment/helper"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// Register mysql server
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is gorm instance Lib
var DB *gorm.DB

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      dao.Model
//    }
type Model struct {
	ID        uint      `gorm:"column:id;primary_key;unique;not null;AUTO_INCREMENT" json:"id" form:"id"`
	CreatedAt time.Time `gorm:"column:created_at;ASSOCIATION_AUTOCREATE;type:timestamp" json:"createdAt" time_utc:"1"`
	DeletedAt *string   `gorm:"column:deleted_at;type:timestamp" json:"-" sql:"index"`
}

func init() {
	var err error
	config := helper.Config
	var uri string = fmt.Sprintf("%s:%s@(%s)/%s?loc=Local&parseTime=True", config.MysqlUsr, config.MysqlPwd, config.MysqlHost, config.MysqlDB)
	DB, err = gorm.Open("mysql", uri)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}

func destory() {
	defer DB.Close()
}
