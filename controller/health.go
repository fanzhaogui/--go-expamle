package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"api.service.com/db"
	"api.service.com/lib"
	"api.service.com/middleware"
	"api.service.com/model"
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h *HealthController) Status(c *gin.Context) {
	var user model.User
	db.GetDB().Where("user_id=?", 2).First(&user)

	// err := db.GetRedis().Set("ok", "good", time.Second*1000)
	// fmt.Printf("%v", err)

	// res := db.GetRedis().Get("ok")

	ret := make(map[string]interface{})
	ret["user"] = user
	ret["res"] = "OK"

	sugar := lib.GetLogger().Sugar()
	for i := 0; i < 1; i++ {
		sugar.Infof("Info : %s", "info error"+strconv.Itoa(i))
	}
	
	uid := middleware.GetWeixinUserId(c)
	fmt.Println(uid)

	loc, _ := time.LoadLocation("Local")
	timeStr := time.Now().Format("2006-01-02 00:00:00")
	localTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	ret["now_date"] = int(localTime.Unix())

	c.JSON(http.StatusOK, ret)
}
