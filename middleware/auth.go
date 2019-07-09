package middleware

import (
	// "fmt"
	"net/http"
	"strings"
	"api.service.com/config"
	"api.service.com/lib"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 鉴权
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := lib.InitApiRes()

		pGet := c.Request.URL.Query()
		params := make(map[string]string)
		for k, p := range pGet {
			params[k] = strings.Join(p, "")
		}
		// sugar := lib.GetLogger().Sugar()
		// sugar.Infof("get param: %v", params)

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
			pPost := c.Request.PostForm
			// fmt.Println("here")
			for k, p := range pPost {
				params[k] = strings.Join(p, "")
			}
		}

		// sugar.Infof("get&post param: %v", params)
		// 获取api_sign字段的值
		apiSign := params["api_sign"]
		if apiSign == "" {
			res.Status = lib.MissParam
			res.Msg = "缺少api_sign参数"
			c.AbortWithStatusJSON(http.StatusOK, res)
			return
		}

		timestamp := params["timestamp"]
		if timestamp == "" {
			res.Status = lib.MissParam
			res.Msg = "缺少timestamp参数"
			c.AbortWithStatusJSON(http.StatusOK, res)
			return
		}

		conf := config.GetConfig()
		apikey := conf.GetString("server.api-key")
		// fmt.Printf("param:%v", params)
		delete(params, "api_sign")
		if lib.MakeSign(params, apikey) != apiSign {
			res.Status = lib.SignError
			res.Msg = lib.Msg(res.Status)
			c.AbortWithStatusJSON(http.StatusOK, res)
			return
		}

		c.Next()
	}
}

// 检查是否登录
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证登入

		c.Next()
	}
}

func GetWeixinUserId(c *gin.Context) int {

	return 1
}
