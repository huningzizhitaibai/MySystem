package test

import (
	"myTestProject/mysql"
	"testing"
)

func TestSearchAidsByQid(t *testing.T) {
	qid1 := 1
	qid2 := 0
	raids := mysql.SearchAidsByQid(qid1)
	waids := mysql.SearchAidsByQid(qid2)
	t.Log(raids, waids)
}
