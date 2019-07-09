package model

type VendorActionLog struct {
	Id      int    `gorm:"primary_key" json:"id"` // 主键
	From    int    `json:"from"`                  //
	UserId  int    `json:"user_id"`               //
	Sid     int    `json:"sid"`                   //
	Ip      string `json:"ip"`                    //
	Content string `json:"content"`               //
	AddTime int    `json:"add_time"`              //创建时间
}

func (VendorActionLog) TableName() string {
	return "vs_vendor_action_log"
}
