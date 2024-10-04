package mysql

import "fmt"

func CheckUsernameByID(uid int) (username string) {
	db := ConnectMyDatabase()

	if db == nil {
		fmt.Println("db is nil")
		return
	}

	res, err := db.Query("select username from UserInfoBasic where uid = ?", uid)
	if err != nil {
		fmt.Println("unfind user")
		return
	}

	for res.Next() {
		var r_username string

		err = res.Scan(&r_username)

		//测试获得的密码是什么
		//fmt.Println(r_password)

		if err != nil {
			return
		}

		defer db.Close()
		return r_username
	}
	return
}
