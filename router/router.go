package router

import (
	"api.service.com/controller"
	"api.service.com/middleware"
	"github.com/gin-gonic/gin"
)

//NewRouter --
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogMiddleware())
	r.Use(gin.Recovery())
	// 加载view目录
	r.LoadHTMLGlob("templates/*")

	// 测试接口
	health := new(controller.HealthController)
	r.GET("/health", health.Status)
	r.POST("/health", health.Status)

	users := new(controller.UsersController)
	r.GET("/user/getUserInfo", users.GetUserInfo)

	//begin---需要验证签名
	r.Use(middleware.AuthMiddleware())

	//begin---需要登录
	r.Use(middleware.MustLogin())


	return r
}
