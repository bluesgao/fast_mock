package model

import "time"

// db-数据库字段 json，binding-参数绑定验证
type Project struct {
	Id        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	//作者id
	//OwnerId int64
	//作者名称
	//OwnerName string
	//项目名称（最多64汉字）
	ProjectName string `db:"project_name"`
	//项目描述（最多128汉字）
	ProjectDesc string `db:"project_desc"`
}
