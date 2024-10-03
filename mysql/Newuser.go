package mysql

import (
	"fmt"
)

func AddNewUser(username, password string) {
	db := ConnectMyDatabase("huning", "Beingalone.1216", "47.98.147.86")

	//说明连接数据库失败了
	if db == nil {
		fmt.Println("db is nil")
	}

	//预定义了一个sql语句
	stmt, err := db.Prepare("insert into users_for_login (username,password) values(?,?)")
	if err != nil {
		fmt.Println(err)
	}

	//在？处插入相关的值,并执行相关的sql语句
	res, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println("执行错误")
	}

	//打印操作的结果和操作的编号
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("查询操作错误")
	}

	aff, err := res.RowsAffected()
	if err != nil {
		fmt.Println("查找操作行错误")
	}
	fmt.Println("id: %d  affected: %d", id, aff)
}
