package model

import (
	"bilibili/drivers/mysql"
	"gorm.io/gorm"
)

type BilibiliDoMsg struct {
	Id int `gorm:"primaryKey"`
	Name string
	Sid string
	Msg string
}

func (bda *BilibiliDoMsg) BilibiliDoMsgAdd(params BilibiliDoMsg)error  {
	var result *gorm.DB
	result = mysql.Db.Create(&params)
	return result.Error
}

