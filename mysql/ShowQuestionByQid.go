package mysql

import "myTestProject/models"

func ShowQuestionByQid(qid uint) models.Question {
	db := ConnectMyDatabase()

	query := "select title, tag, content from Question where qid=?"
	var question = models.Question{}

	db.QueryRow(query, qid).Scan(
		&question.Title,
		&question.Tag,
		&question.Content,
	)

	return question
}
