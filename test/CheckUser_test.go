package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestCheckUser(t *testing.T) {
	_username := "huning"
	_rightpassword := "123456"
	_wrongpassword := "455667"

	result1 := mysql.CheckUser(_username, _rightpassword)
	result2 := mysql.CheckUser(_username, _wrongpassword)

	if result1 == true {
		t.Log("测试失败，用户名不正确")
	} else {
		t.Log("测试成功，函数错误")
	}

	if result2 == true {
		t.Log("测试错误，函数错误")
	} else {
		t.Log("测试成功，密码是错误的")
	}

}
