package models

/*
定义了一个关于用户相关信息的模版
对于每个用户基础信息进行收集和保存
*/

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

// TableName 这是对这个表进行初始化的的一个操作
// 虽然在go语言中没有类，但是这应该可以相当于是一个类的方法
func (table *uerBasic) TableName() string {
	return "uer_Basics"
}
