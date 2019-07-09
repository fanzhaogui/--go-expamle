package controller

import (
	"api.service.com/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct{}

func (u *UsersController) GetUserInfo(c *gin.Context) {
	res := lib.InitApiRes()

	err, code := func() (error, string){
		ret := make(map[string]interface{})
		userId := 1
		userInfo, err := lib.GetUserInfo(userId)
		if nil != err {
			return err, lib.SystemBusy
		}

		ret["user_info"] = userInfo
		res.Data = ret
		c.JSON(http.StatusOK, res)
		return nil, ""
	}()

	if err != nil {
		res.Status = code
		res.Msg = err.Error()
		c.AbortWithStatusJSON(http.StatusOK, res)
	}
	return
}


func (u *UsersController) UpdateInfo(c *gin.Context) {
	res := lib.InitApiRes()

	c.JSON(http.StatusOK, res)
}

