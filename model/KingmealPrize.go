package model

type KingmealPrize struct {
	PrizeId     int    `gorm:"primary_key" json:"prize_id"`
	IssueId     int    `json:"issue_id"`
	PrizeName   string `json:"prize_name"`
	PrizeImg    string `json:"prize_img"`
	Rate        int    `json:"rate"`
	Description string `json:"description"`
	Rank        int    `json:"rank"`
	StockNum    int    `json:"stock_num"`
	SurplusNum  int    `json:"surplus_num"`
	GrantNum    int    `json:"grant_num"`
	AddTime     int    `json:"add_time"`
	UpdateTime  int    `json:"update_time"`
	ExpireType  int    `json:"expire_type"`
	ExpireTime  int    `json:"expire_time"`
	ExpireDays  int    `json:"expire_days"`
}

func (KingmealPrize) TableName() string {
	return "act_kingmeal_prize_setting"
}
