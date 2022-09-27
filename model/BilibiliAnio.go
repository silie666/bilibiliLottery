package model

import (
	"bilibili/drivers/mysql"
	"bilibili/logger"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type BilibiliAnio struct {
	Id                   int `gorm:"primaryKey"`
	OriginDynamicIdStr   string
	OriginRidStr         string
	OriginType           int
	OriginUid            int
	PreviousDynamicIdStr string
	PreviousRidStr       string
	PreviousType         int
	PreviousUid          int
	IsOk                 int
	IsModify             int
	IsComment            int
	IsRepost             int
	IsLike               int
	Bvid                 string
	JsonData             string `gorm:"mediumtext"`
	Str                  string
	ForwardUid           string
	Ctrl                 string
}

func (ba *BilibiliAnio) BilibiliAnioAdd(params BilibiliAnio) error {
	var result *gorm.DB
	var info BilibiliAnio
	err := mysql.Db.Where("origin_dynamic_id_str = ? and previous_dynamic_id_str = ? and forward_Uid = ?", params.OriginDynamicIdStr, params.PreviousDynamicIdStr, params.ForwardUid).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result = mysql.Db.Create(&params)
		fmt.Println("添加成功")
	} else {
		logger.LogToFile("添加失败")
		return nil
	}
	return result.Error
}

func (ba *BilibiliAnio) IsBilibiliAnio(where string) bool {
	var info BilibiliAnio
	err := mysql.Db.Where(where).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else {
		return true
	}
}

func (ba *BilibiliAnio) BilibiliAnioSave(params BilibiliAnio) error {
	var result *gorm.DB
	var info BilibiliAnio
	err := mysql.Db.Where("id = ?", params.Id).First(&info).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		result = mysql.Db.Updates(&params)
	}
	return result.Error
}

func (cm *BilibiliAnio) BilibiliAnioList(where string) []BilibiliAnio {
	var bilibiliAnio []BilibiliAnio
	mysql.Db.Where(where).Find(&bilibiliAnio)
	return bilibiliAnio
}
