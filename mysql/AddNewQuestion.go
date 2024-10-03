package mysql

import (
	"fmt"
	"myTestProject/models"
)

func AddNewQuestion(question models.Question) bool {
	//连接上我的数据库
	db := ConnectMyDatabase()

	//说明连接数据库失败了
	if db == nil {
		fmt.Println("db is nil")
		return false
	}

	//预定义了一个sql语句
	stmt, err := db.Prepare("insert into Question (title,tag,content) values(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return false
	}

	//在？处插入相关的值,并执行相关的sql语句
	stmt.Exec(question.Title, question.Tag, question.Content)

	//执行完增加操作后就将数据库进行关闭
	defer db.Close()
	return true
}
