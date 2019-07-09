package model

//Business --
type Business struct {
	ID       int    `gorm:"primary_key" json:"id"` //账号ID
	ShopName string `json:"shop_name"`
}

//TableName --
func (Business) TableName() string {
	return "vs_business"
}
