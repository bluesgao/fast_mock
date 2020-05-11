package dao

import (
	"fast_mock/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"log"
)

var database *Database

type Database struct {
	dbCli *sqlx.DB
}

//连接数据库
func (db *Database) Init(conf *conf.Conf) {
	log.Printf("database init conf:%+v \n", conf)
	//DbCli, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名?charset=编码")
	//连接字符串中新增parseTime=true，防止时间类型转换错误
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", conf.Db.Username, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.Database, conf.Db.Charset)
	//root:Gx-123456@tcp(47.97.205.190:3306)/test?charset=utf8mb4&parseTime=true
	log.Printf("database init connStr:%+v \n", connStr)

	client, err := sqlx.Connect(conf.Db.Driver, connStr)
	log.Printf("database init client err:%+v \n", err)
	if err != nil {
		panic(err)
	}
	log.Printf("database init client:%+v  \n", client)
	db.dbCli = client
	database = db
}

func (db *Database) GetDbCli() *sqlx.DB {
	return database.dbCli
}
