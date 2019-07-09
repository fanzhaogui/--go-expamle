package model

type Config struct {
	Id          int    `gorm:"primary_key" json:"id"` //ID
	OptionKey   string `json:"option_key"`
	OptionValue string `json:"option_value"`
	OptionText  string `json:"option_text"`
}

func (Config) TableName() string {
	return "act_config"
}
