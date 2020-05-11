package model

import "time"

type Module struct {
	Id        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	//项目id
	ProjectId int64 `db:"project_id"`
	//项目名称
	ProjectName string `db:"project_name"`
	//模块名称（最多64汉字）
	ModuleName string `db:"module_name"`
	//模块描述（最多128汉字）
	ModuleDesc string `db:"module_desc"`
}
