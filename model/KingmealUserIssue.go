package model

import "api.service.com/db"

type KingmealUserIssue struct {
	UserIssueId int `gorm:"primary_key" json:"user_issue_id"`
	IssueId     int `json:"issue_id"`
	FreeSeats   int `json:"free_seats"`
	AddTime     int `json:"add_time"`
	UpdateTime  int `json:"update_time"`
	UserId      int `json:"user_id"`
}

func (KingmealUserIssue) TableName() string {
	return "act_kingmeal_user_issue"
}

/**
 * 获取主题用户的参与信息
 */
func GetLastUserIssue(issueId int, userId int) (KingmealUserIssue, error) {
	var desks KingmealUserIssue
	querySql := "SELECT * FROM act_kingmeal_user_issue WHERE issue_id = ? AND user_id = ? ORDER BY add_time DESC" +
		" LIMIT 1 "
	err := db.GetDB().Raw(querySql, issueId, userId).Scan(&desks).Error
	if err != nil {

	}
	return desks, nil
}