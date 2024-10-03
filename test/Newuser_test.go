package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestNewuser(t *testing.T) {
	username := "Lee"
	password := "177533"
	result := mysql.AddNewUser(username, password)
	if result {
		t.Log("success")
	} else {
		t.Log("fail")
	}
}
