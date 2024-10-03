package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestSearchInDatabase(t *testing.T) {
	keyword1 := "优雅"
	keyword2 := "同时"

	qids1 := mysql.SearchInDatabase(keyword1)
	qids2 := mysql.SearchInDatabase(keyword2)
	t.Log(qids1, qids2)
}
