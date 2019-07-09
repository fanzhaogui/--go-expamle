package model

type Users struct {
	UserId     int    `gorm:"primary_key" json:"user_id"` // 自增ID
	UserName   string `json:"user_name"`
	Tel        string `json:"tel"`
	Addr       string `json:"addr"`
	UpdateTime int    `json:"update_time"`
	AddTime    int    `json:"add_time"`
	Gender     int    `json:"gender"`
	Age        int    `json:"age"`
}

func (Users) TableName() string {
	return "users"
}
