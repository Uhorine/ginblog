package model

import (
	"fmt"
	"ginblog/utils/errmsg"
)

type Profile struct {
	ID     int    `grom:"primaryKey" json:"id"`
	Name   string `grom:"type:varchar(20)" json:"name"`
	Avatar string `grom:"type:varchar(200)" json:"avatar"`
	Img    string `grom:"type:varchar(200)" json:"img"`
	Desc   string `grom:"type:varchar(200)" json:"desc"`
	Qqchat string `grom:"type:varchar(200)" json:"qqchat"`
	Wechat string `grom:"type:varchar(100)" json:"wechat"`
}

func GetProfile(id int) (Profile, int) {
	var profile Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	fmt.Println(data)
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
