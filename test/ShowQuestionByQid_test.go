package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestShowQuestionByQid(t *testing.T) {
	var qid uint
	qid = 1
	question := mysql.ShowQuestionByQid(qid)
	t.Log(question)
}
