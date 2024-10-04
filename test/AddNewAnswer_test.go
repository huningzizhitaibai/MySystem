package test

import (
	"myTestProject/models"
	"myTestProject/mysql"
	"testing"
)

func TestAddNewAnswer_test(t *testing.T) {
	answer := models.Answer{
		"我不知道a",
		"huning",
		1,
	}
	ret := mysql.AddNewAnswer(answer)
	if ret != true {
		t.Log("add new answer failed")
	} else {
		t.Log("add new answer success")
	}
}
