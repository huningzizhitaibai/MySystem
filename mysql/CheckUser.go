package mysql

import "fmt"

func CheckUser(username string, password string) bool {
	db := ConnectMyDatabase()

	if db == nil {
		fmt.Println("db is nil")
		return false
	}

	res, err := db.Query("select password from UserInfoBasic where username = ?", username)
	if err != nil {
		fmt.Println("unfind password")
		return false
	}

	for res.Next() {
		var r_password string

		err = res.Scan(&r_password)

		//测试获得的密码是什么
		fmt.Println(r_password)

		if err != nil {
			return false
		}
		if r_password == password {
			return true
		}
	}
	defer db.Close()
	return false
}
