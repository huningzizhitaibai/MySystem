package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestNewuser(t *testing.T) {
	username := "huning"
	password := "123456"
	mysql.AddNewUser(username, password)
}
