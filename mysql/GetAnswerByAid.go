package mysql

import (
	"fmt"
	"myTestProject/models"
)

func GetAnswerByAid(aid int) models.Answer {
	db := ConnectMyDatabase()

	if db == nil {
		fmt.Println("db is nil")
		return models.Answer{}
	}

	res, err := db.Query("select Content,Provider from Answer where Aid = ?", aid)
	if err != nil {
		fmt.Println("unfind Answer")
		return models.Answer{}
	}

	for res.Next() {
		var r_anwser models.Answer

		err = res.Scan(&r_anwser.Content, &r_anwser.Provider)

		//测试获得的密码是什么
		//fmt.Println(r_password)

		if err != nil {
			return models.Answer{}
		}
		return r_anwser
	}
	defer db.Close()
	return models.Answer{}
}
