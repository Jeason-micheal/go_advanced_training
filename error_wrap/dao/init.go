package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 初始化, 提供Open操作, sqlx有做close 所以不需要做
var (
	_db *sqlx.DB
	dsn = "root:123456@tcp(127.0.0.1:3306)/user"
	// "root:123456@tcp(127.0.0.1:3306)/user"
)

//type DB struct {
//	DB    *sqlx.DB
//	table string
//	// ....
//}

func init() {
	_db = initDB()
}

func initDB() (db *sqlx.DB) {
	// open db...
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return
}

// 返回包内的成员, 不能定义返回值的名称 , 不然是一个中间变量
func DB() *sqlx.DB {
	return _db
}
