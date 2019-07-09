package lib_test

import (
	"api.service.com/config"
	"api.service.com/db"
	"flag"
	"testing"
	"fmt"
	"api.service.com/lib"
)

func init() {
	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		// t.Errorf("Usage: server -e {mode} #mode should be: dev|test|pro")
	}
	flag.Parse()
	config.Init(*env)

	// logger
	lib.InitLogger()
	defer lib.GetLogger().Sync()

	// mysql
	db.InitMysql()
	// defer db.GetDB().Close()

	// redis
	// db.InitRedis()
	// defer db.GetRedis().Close()
	db.GetDB().LogMode(true)
}

func TestUsersController_GetUserInfo(t *testing.T) {
	userInfo, err := lib.GetUserInfo(1)
	if err != nil {
		t.Error("desk not found")
	}

	fmt.Println(userInfo)
}
