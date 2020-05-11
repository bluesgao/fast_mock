package entity

import "time"

type Project struct {
	Id        int64 `db:"id"`
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
