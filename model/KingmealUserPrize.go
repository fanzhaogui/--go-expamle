package model

type KingmealUserPrize struct {
	Id         int `gorm:"primary_key" json:"id"`
	PrizeId    int `json:"prize_id"`
	DeskId     int `json:"desk_id"`
	IsUsed     int `json:"is_used"`
	AddTime    int `json:"add_time"`
	UpdateTime int `json:"update_time"`
	UserId     int `json:"user_id"`
	ExpireTime int `json:"expire_time"`
}

func (KingmealUserPrize) TableName() string {
	return "act_kingmeal_user_prize"
}
