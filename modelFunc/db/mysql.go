package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() (err error) {
	// 配置账户密码，以及ip地址端口号
	dsn := "root:admin@tcp(127.0.0.1:3306)/goodsk?charset=utf8mb4"
	// 也可以使用MustConnect连接不成功就panic
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	return DB.Ping()
}

func Close() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
