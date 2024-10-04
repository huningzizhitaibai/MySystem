package mysql

import "fmt"

func GetUidByUsername(username string) int {
	db := ConnectMyDatabase()

	if db == nil {
		fmt.Println("db is nil")
		return 0
	}

	res, err := db.Query("select uid from UserInfoBasic where username = ?", username)
	if err != nil {
		fmt.Println("unfind user")
		return 0
	}

	for res.Next() {
		var r_uid int

		err = res.Scan(&r_uid)

		//测试获得的密码是什么
		//fmt.Println(r_password)

		if err != nil {
			return 0
		}
		return r_uid
	}
	defer db.Close()
	return 0
}
