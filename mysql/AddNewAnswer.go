package mysql

import (
	"fmt"
	"myTestProject/models"
)

func AddNewAnswer(answer models.Answer) bool {
	db := ConnectMyDatabase()

	stmt, err := db.Prepare("insert into Answer (Content, Provider,ConnectQuestion) values(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return false
	}

	//在？处插入相关的值,并执行相关的sql语句
	stmt.Exec(answer.Content, answer.Provider, answer.ConectQuestion)
	defer db.Close()
	return true
}
