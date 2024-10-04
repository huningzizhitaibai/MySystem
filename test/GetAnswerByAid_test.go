package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestGetAnswerByQid(t *testing.T) {
	aid := 1
	waid := 77
	rret := mysql.GetAnswerByAid(aid)
	wret := mysql.GetAnswerByAid(waid)
	t.Log(rret)
	t.Log(wret)
}
