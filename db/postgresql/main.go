package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	//导入pq驱动程序，通过包database/sql来注册数据库驱动程序
	_ "github.com/lib/pq"
)

func main() {
	//调用openDB()帮助函数(参见下面)来创建连接池
	db, err := openDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://admin:123456@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	//创建一个具有5秒超时期限的上下文。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//使用PingContext()建立到数据库的新连接，并传入上下文信息，连接超时就返回
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	// 返回sql.DB连接池
	return db, nil
}
