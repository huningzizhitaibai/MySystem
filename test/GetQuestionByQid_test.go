package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestGetQuestionByQid(t *testing.T) {
	qid := 1
	wqid := 77
	rret := mysql.GetQuestionByQid(qid)
	wret := mysql.GetQuestionByQid(wqid)
	t.Log(rret)
	t.Log(wret)
}
