package model
// d, issue_name, start_time, end_time
type KingmealIssue struct {
	IssueId   		int `gorm:"primary_key" json:"issue_id"`
	Periods 		string `json:"periods"`
	IssueName       string `json:"issue_name"`
	StartTime 		int `json:"start_time"`
	EndTime 		int `json:"end_time"`
	BusinessId      int `json:"business_id"`
	SwitchStatus 	int `json:"switch_status"`
}

func (KingmealIssue) TableName() string {
	return "act_kingmeal_issue"
}
