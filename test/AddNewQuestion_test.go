package test

import (
	"myTestProject/models"
	"myTestProject/mysql"
	"testing"
)

func TestAddNewQuestion(t *testing.T) {
	testQuestion := models.Question{
		"如何优雅编程？",
		"other",
		"如何提升自己的编程技巧，同时是自己写出来的代码显得优雅简洁？",
	}

	if mysql.AddNewQuestion(testQuestion) {
		t.Log("添加成功")
	} else {
		t.Log("添加失败")
	}
}
