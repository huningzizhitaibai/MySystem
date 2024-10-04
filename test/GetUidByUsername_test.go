package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestGetUidByUsername(t *testing.T) {
	username := "huning"
	uid := mysql.GetUidByUsername(username)
	if uid == 0 {
		t.Log("错误")
	} else {
		t.Log(uid)
	}
}
