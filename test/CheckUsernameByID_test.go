package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestCheckUsernameByID(t *testing.T) {
	uid := 1
	username := mysql.CheckUsernameByID(uid)
	if username != "" {
		t.Log(username)
	} else {
		t.Error(username)
	}
}
