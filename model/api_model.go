package model

import "time"

type Api struct {
	Id        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	//项目id
	ProjectId int64 `db:"project_id"`
	//项目名称
	ProjectName string `db:"project_name"`
	//模块id
	ModuleId int64 `db:"module_id"`
	//模块名称
	ModuleName string `db:"module_name"`
	ApiDesc    string `db:"api_desc"`
	Request    string
	Response   string
}
