package db

import (
	"fmt"

	"api.service.com/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlClient *gorm.DB

func InitMysql() {
	conf := config.GetConfig()
	conn, err := gorm.Open("mysql", conf.GetString("mysql.dsn"))
	if err != nil {
		panic("mysql error:" + fmt.Sprintf("%v", err))
	}
	mysqlClient = conn
}

func GetDB() *gorm.DB {
	return mysqlClient
}
