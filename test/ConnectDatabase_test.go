package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestConnectDatabase(t *testing.T) {

	got := mysql.ConnectMyDatabase()

	if got == nil {
		t.Errorf("未连接到数据库")
	} else {
		t.Log("连接正常")
	}
}
