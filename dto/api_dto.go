package dto

import "time"

type Api struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	//项目id
	ProjectId int64
	//项目名称
	ProjectName string
	//模块id
	ModuleId int64
	//模块名称（最多64汉字）
	ModuleName string
	//模块描述（最多128汉字）
	ModuleDesc string
	Request string
	Response string
}
