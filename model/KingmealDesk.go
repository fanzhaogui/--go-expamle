package model

import "api.service.com/db"

type KingmealDesk struct {
	DeskId          int `gorm:"primary_key" json:"desk_id"`
	UserIssueId     int `json:"user_issue_id"`
	UserId          int `json:"user_id"`
	AddTime         int `json:"add_time"`
	UpdateTime      int `json:"update_time"`
	NotifiedStatus  int `json:"notified_status"`
	IsDraw          int `json:"is_draw"`
	AssistanceCount int `json:"assistance_count"`
}

func (KingmealDesk) TableName() string {
	return "act_kingmeal_desk"
}

/**
 * 获取主题用户的最近一桌信息
 */
func GetLastDeskInfo(userIssueId int, userId int) (KingmealDesk, error) {
	var desk KingmealDesk
	querySql := "SELECT * FROM act_kingmeal_desk WHERE user_issue_id = ? AND user_id = ? ORDER BY add_time DESC LIMIT 1 "
	err := db.GetDB().Raw(querySql, userIssueId, userId).Scan(&desk).Error
	if err != nil {
		return desk, err
	}
	return desk, nil
}
