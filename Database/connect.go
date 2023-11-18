package Database

import (
	"database/sql"
	"fmt"
)

func ConnectToDatabase(username string, password string, host string, port int, dbname string) (*sql.DB, error) {
	dsnamestring := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
	fmt.Println(dsnamestring)
	// 创建一个数据库连接池，设置最大连接数和最大空闲连接数
	db, err := sql.Open("mysql", dsnamestring)
	if err != nil {
		return nil, err
	}

	// 设置连接池的参数
	db.SetMaxOpenConns(10)   // 最大打开连接数
	db.SetMaxIdleConns(5)    // 最大空闲连接数
	db.SetConnMaxLifetime(0) // 连接的最大生命周期（0 表示不限制）

	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
