package model

type CouponAd struct {
	Id      		int    `gorm:"primary_key" json:"id"` //自增ID
	CityId     		int    `json:"city_id"`    //创建时间
	PositionId      int    `json:"position_id"`    //创建时间
	CouponTid     	int    `json:"coupon_tid"`    //创建时间
	AdName      	string `json:"ad_name"`
	ImgUrl    		string `json:"img_url"`
	Type     		int    `json:"type"`    //创建时间
	RefererUrl   	string `json:"referer_url"`
	Desc     		string `json:"desc"`
	IsShow       	int    `json:"is_show"`
	CreateTime  	int    `json:"create_time"` //修改时间
	LastUpdateTime  int    `json:"last_update_time"` //修改时间
}

func (CouponAd) TableName() string {
	return "vs_coupon_ad"
}
