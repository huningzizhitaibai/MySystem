package test

import (
	"fmt"
	"myTestProject/mysql"
	"reflect"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	username := "huning"
	password := "Beingalone.1216"
	ip := "47.98.147.86"

	fmt.Scanf("%s %s %s", &username, &password, &ip)
	got := mysql.ConnectMyDatabase(username, password, ip)

	if reflect.DeepEqual(got, nil) {
		t.Errorf("未连接到数据库")
	} else {
		t.Log("连接正常")
	}
}
