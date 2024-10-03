package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectMyDatabase(user string, password string, ip string) (db *sql.DB) {
	//成功连接位于服务器上的mysql数据库

	dsn := user + ":" + password + "@tcp(" + ip + ")/gogin"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
		log.Fatal(err)

	}

	err = db.Ping()
	if err != nil {
		return nil
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
