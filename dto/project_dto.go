package dto

import (
	"fast_mock/model"
	"time"
)

//json，binding-参数绑定验证
type ProjectDto struct {
	Id          int64     `json:"id" form:"id"`
	CreatedAt   time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" form:"updatedAt"`
	ProjectName string    `json:"projectName" form:"projectName" binding:"required,gte=2,lte=60"`
	//项目描述（最多128汉字）
	ProjectDesc string `json:"projectDesc" form:"projectDesc" binding:"required,gte=2,lte=60"`
}

func ToProjectModel(dto ProjectDto) model.Project {
	return model.Project{
		Id:          dto.Id,
		ProjectName: dto.ProjectName,
		ProjectDesc: dto.ProjectDesc,
	}
}

func ToProjectDto(model model.Project) ProjectDto {
	return ProjectDto{
		Id:          model.Id,
		ProjectName: model.ProjectName,
		ProjectDesc: model.ProjectDesc,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
