#fast mock
接口mock系统
## 技术架构
fasthttp + httprouter + sqlx + mysql + go



var Db *sqlx.DB
//连接数据库
func init(){
	var err error
	//Db, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名?charset=编码")
	Db, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
}

type Member struct {
	Username string `db:"username"`
	Money float64   `db:"money"`
	Birthday sql.NullString `db:"birthday"`
}

func Select(){
	//查一条
	var info Member
	err:=Db.Get(&info,"SELECT username,money,birthday FROM member WHERE id=?",1)
	if err!=nil {
		fmt.Println(err);return
	}
	fmt.Println(info)

	//查询多条
	var list []Member
	err=Db.Select(&list,"SELECT username,money,birthday FROM member")
	if err != nil {
		fmt.Println(err);return
	}
	fmt.Println(list)
}

func Insert(){
	result,err:=Db.Exec("INSERT INTO member(username,money,created_at)VALUES (?,?,?)","test",20,time.Now().Unix())
	if err != nil {
		fmt.Println(err);return
	}
	id,err:=result.LastInsertId()
	if err != nil {
		fmt.Println(err);return
	}
	fmt.Println(id)
}

func Update(){
	result,err:=Db.Exec("UPDATE member SET money=money+3 WHERE id=?",1)
	if err != nil {
		fmt.Println(err);return
	}
	rows,err:=result.RowsAffected()
	if err != nil {
		fmt.Println(err);return
	}
	fmt.Println(rows)
}

func Delete(){
	result,err:=Db.Exec("DELETE FROM member WHERE id=?",7)
	if err != nil {
		fmt.Println(err);return
	}
	rows,err:=result.RowsAffected()
	if err != nil {
		fmt.Println(err);return
	}
	fmt.Println(rows)
}

func Transaction(){
	db,err:=Db.Begin()
	if err!=nil {
		fmt.Println(err);return
	}
	defer clearTransaction(db) //如果出现了异常，导致没有 commit和rollback，可以用来收尾

	//1、
	result,err:=db.Exec("UPDATE member SET money=money+3 WHERE id=?",1)
	if err!=nil {
		fmt.Println(err);return
	}
	num,err:=result.RowsAffected()
	if err != nil {
		db.Rollback()
		fmt.Println(err);return
	}
	fmt.Println(num)

	//2、
	result,err=db.Exec("UPDATE member SET money=money-3 WHERE id=?",2)
	if err!=nil {
		db.Rollback()
		fmt.Println(err);return
	}
	num,err=result.RowsAffected()
	if err != nil {
		db.Rollback()
		fmt.Println(err);return
	}
	db.Commit()
}


//事务回滚
func clearTransaction(tx *sql.Tx){
	fmt.Println("这里负责收尾，rollback")
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil{
		fmt.Println(err)
	}
}
