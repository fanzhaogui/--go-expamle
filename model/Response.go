package model

type ResponDrawPrize struct {
	PrizeId     int    `json:"prize_id"`
	PrizeName   string `json:"prize_name"`
	Description string `json:"description"`
	AddTime     string `json:"add_time"`
	IssueName   string `json:"issue_name"`
	ExpireTime  string `json:"expire_time"`
}

type ResponListPrize struct {
	PrizeId    int    `json:"prize_id"`
	PrizeName  string `json:"prize_name"`
	PrizeImg   string `json:"prize_img"`
	ShopName   string `json:"shop_name"`
	IssueName  string `json:"issue_name"`
	ExpireTime string `json:"expire_time"`
	AddTime    string `json:"add_time"`
}
