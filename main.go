package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"api.service.com/config"
	"api.service.com/db"
	"api.service.com/lib"
	"api.service.com/router"
	"github.com/gin-gonic/gin"
)

func main() {
	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode} #mode should be: dev|test|pro")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*env)

	//logger
	lib.InitLogger()
	defer lib.GetLogger().Sync()

	//mysql
	db.InitMysql()
	defer db.GetDB().Close()

	//redis 要先安装-本地开启reids-service
	// db.InitRedis()
	// defer db.GetRedis().Close()

	if *env == "pro" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		db.GetDB().LogMode(true)
	}

	conf := config.GetConfig()
	r := router.NewRouter()

	srv := &http.Server{
		Addr:    ":" + conf.GetString("server.port"),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	ioutil.WriteFile(conf.GetString("server.pid-file"), []byte(strconv.Itoa(syscall.Getpid())), 0644)
	log.Printf("Actual pid is %d", syscall.Getpid())

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	timeout := time.Duration(conf.GetInt("server.shutdown-timeout"))
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
