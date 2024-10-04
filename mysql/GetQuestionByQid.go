package mysql

import (
	"fmt"
	"myTestProject/models"
)

func GetQuestionByQid(qid int) models.Question {
	db := ConnectMyDatabase()

	if db == nil {
		fmt.Println("db is nil")
		return models.Question{}
	}

	res, err := db.Query("select title, tag, content from Question where qid = ?", qid)
	if err != nil {
		fmt.Println("unfind question")
		return models.Question{}
	}

	for res.Next() {
		var r_question models.Question

		err = res.Scan(&r_question.Title, &r_question.Tag, &r_question.Content)

		//测试获得的密码是什么
		//fmt.Println(r_password)

		if err != nil {
			return models.Question{}
		}
		return r_question
	}
	defer db.Close()
	return models.Question{}
}
