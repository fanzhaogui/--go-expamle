package model

type User struct {
	UserId   int    `gorm:"primary_key" json:"user_id"` // 主键
	NickName string `gorm:"size:255" json:"nick_name"`  // string默认长度为255。
}

func (User) TableName() string {
	return "vs_user"
}
