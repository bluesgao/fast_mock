package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//连接数据库
func InitDb() *sqlx.DB {
	var err error
	//Db, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名?charset=编码")
	Db, err = sqlx.Open("mysql", "root:Gx-123456@tcp(47.97.205.190:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	return Db
}

func GetDb() *sqlx.DB {
	return Db
}
