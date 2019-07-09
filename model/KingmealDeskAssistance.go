package model

import (
	"api.service.com/db"
	"fmt"
	"time"
)

type KingmealDeskAssistance struct {
	AssId      int `gorm:"primary_key" json:"ass_id"`
	DeskId     int `json:"desk_id"`
	AssUserId  int `json:"ass_user_id"`
	AddTime    int `json:"add_time"`
	UpdateTime int `json:"update_time"`
}

func (KingmealDeskAssistance) TableName() string {
	return "act_kingmeal_desk_assistance"
}

func GetTodayAssedTimes(userId int) int {
	st := time.Now().Unix()
	et := time.Now().AddDate(0,0, 1).Unix()
	start :=getDayStartUnix(st)
	end := getDayStartUnix(et)
	fmt.Printf("%s~~%s",start, end)
	sql := "SELECT count(*) FROM act_kingmeal_desk_assistance WHERE add_time >= ? AND add_time < ? AND ass_user_id = ?";
	var total int
	db.GetDB().Raw(sql, start, end, userId).Count(&total)
	return total
}

func getDayStartUnix(sec int64) int64{
	t := time.Unix(sec, 0).Format("2006-01-02")
	sts, _ := time.Parse("2006-01-02", t)
	return sts.Unix()
}