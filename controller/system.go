package controller

type SystemController struct{}

// func (s *SystemController) CollectFormId(c *gin.Context) {
// 	formId := c.Query("form_id")
//
// 	if c.Request.Method == "POST" {
// 		formId = c.PostForm("form_id")
// 	}
//
// 	res := lib.InitApiRes()
// 	if formId == "" {
// 		res.Status = lib.MissParam
// 		res.Msg = lib.Msg(res.Status)
// 		c.AbortWithStatusJSON(http.StatusOK, res)
// 		return
// 	}
//
// 	wxUid := middleware.GetWeixinUserId(c)
//
// 	err := lib.InsertFormId(wxUid, formId)
// 	if err != nil {
// 		lib.GetLogger().Sugar().Errorf("CollectFormId error:%+v", err)
// 		res.Status = lib.BusinessError
// 		res.Msg = err.Error()
// 		c.AbortWithStatusJSON(http.StatusOK, res)
// 		return
// 	}
// 	c.JSON(http.StatusOK, res)
// }

// func (s *SystemController) Static(c *gin.Context) {
// 	res := lib.InitApiRes()
// 	userId := middleware.GetWeixinUserId(c)
// 	userInfo := lib.GetUserRes(userId)
// 	data := make(map[string]interface{})
// 	data["user_info"] = userInfo
// 	res.Data = data
// 	c.JSON(http.StatusOK, res)
// }
