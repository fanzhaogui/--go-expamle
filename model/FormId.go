package model

type FormId struct {
	Id         int    `gorm:"primary_key" json:"id"`
	FormIdType int    `json:"form_id_type"`
	FormIdFrom string `json:"form_id_from"`
	FormId     string `json:"form_id"`
	UserId     int    `json:"user_id"`
	OpenId     string `json:"open_id"`
	CreateTime int    `json:"create_time"`
	ExpiryTime int    `json:"expiry_time"`
	FormStatus uint8  `json:"form_status"`
}

func (FormId) TableName() string {
	return "act_formid"
}
