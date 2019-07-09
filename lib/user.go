package lib

import (
	"api.service.com/model"
	"api.service.com/db"
)

type UserRes struct {
	model.Users
	SessionID string `json:"session_id"`
}

/**
 * 指定用户的信息
 */
func  GetUserInfo(userId int) (UserRes, error) {
	var user UserRes
	sql := "SELECT user_id, user_name, tel, addr, update_time, add_time, gender, age FROM users WHERE user_id = ? Limit 1"
	err := db.GetDB().Raw(sql, userId).Scan(&user).Error
	if nil != err {
		return user, err
	}
	return user, nil
}
