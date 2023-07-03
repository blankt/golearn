package database

import (
	"database/sql"
	"fmt"
	"time"
)

func Open() {
	//open并不会建立一个连接 只是返回一个db对象 在ping或者第一次操作数据库时才会建立建立数据库 所以这里dsn错误也不会报错
	db, err := sql.Open("mysql", "xxx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	time.Sleep(time.Second * 10)
}

func Connection() {
	db, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	//设置数据库最大连接数 包括闲置
	db.SetMaxOpenConns(6)
	// 设置最大闲置连接数
	db.SetMaxIdleConns(5)
	fmt.Println("please exec show processlist")
	time.Sleep(10 * time.Second)
	fmt.Println("please exec show processlist again")
	db.Ping()
	time.Sleep(10 * time.Second)
}
