package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"log"
)

var database *Database

type Database struct {
	Db *sqlx.DB
}

//连接数据库
func (db *Database) Init() {
	log.Printf("database init")
	//Db, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名?charset=编码")
	client, err := sqlx.Connect("mysql", "root:Gx-123456@tcp(47.97.205.190:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	log.Printf("database client:%+v", client)
	db.Db = client
	database = db

}

func (db *Database) GetDb() *sqlx.DB {
	return database.Db
}
