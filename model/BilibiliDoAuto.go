package model

import (
	"bilibili/drivers/mysql"
	"errors"
	"gorm.io/gorm"
)

type BilibiliDoAuto struct {
	Id int `gorm:"primaryKey"`
	Url string
	JsonData string `gorm:"mediumtext"`
	IsModify int
	Name string
	Mid int
	Sid string
	Num int
}

func (bda *BilibiliDoAuto)BilibiliDoAutoList()[]BilibiliDoAuto {
	var BilibiliDoAutoModel []BilibiliDoAuto
	mysql.Db.Find(&BilibiliDoAutoModel)
	return BilibiliDoAutoModel
}

func (bda *BilibiliDoAuto) BilibiliDoAutoEdit(params BilibiliDoAuto)error  {
	var result *gorm.DB
	err := mysql.Db.Where("url = ?", params.Url).First(bda).Error
	if errors.Is(err,gorm.ErrRecordNotFound) {
		result = mysql.Db.Create(&params)
	} else {
		params.Id = bda.Id
		params.Mid = bda.Mid
		result = mysql.Db.Save(&params)
	}
	return result.Error
}

func (bda *BilibiliDoAuto)BilibiliDoDel(id int)error {
	result := mysql.Db.Where("id = ?", id).Delete(bda)
	return result.Error
}