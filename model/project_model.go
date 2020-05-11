package model

import "time"

type Project struct {
	Id        int64 `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	//作者id
	//OwnerId int64
	//作者名称
	//OwnerName string
	//项目名称（最多64汉字）
	ProjectName string `json:"projectName" binding:"required,gte=2,lte=60"`
	//项目描述（最多128汉字）
	ProjectDesc string `json:"projectDesc" binding:"required,gte=2,lte=60"`
}
