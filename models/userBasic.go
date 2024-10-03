package models

import "github.com/jinzhu/gorm"

type uerBasic struct {
	//为用户结构添加了基础的几个字段
	gorm.Model
	Name         string
	password     string
	phoneNumber  string
	emailAddress string
	clientIP     string
	ID           string
}

func (table *uerBasic) TableName() string {
	return "uer_Basics"
}
