package biz

import (
	"fast_mock/dao"
	"fast_mock/model"
	"fast_mock/util"
	"github.com/gin-gonic/gin"
	"log"
)

type ProjectBiz struct {
	dao dao.ProjectDao
}

func NewProjectBiz() ProjectBiz {
	return ProjectBiz{
		dao: dao.NewProjectDao(),
	}
}

func (biz ProjectBiz) CreateProject(ctx *gin.Context) {
	log.Printf("ProjectBiz.CreateProject: %+v \n", ctx.Request)
	var project model.Project

	if err := ctx.Bind(&project); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	log.Printf("ProjectBiz.CreateProject project: %+v \n", project)

	if id, err := biz.dao.Create(project); err != nil {
		util.ResponseByErr(ctx, "创建错误", err.Error())
	} else {
		log.Printf("id:%d, err:%s", id, err)
		project.Id = id
		util.ResponseByOk(ctx, "创建成功", &project)
	}
}

func (biz ProjectBiz) ListProject(ctx *gin.Context) {
	log.Printf("ProjectBiz.ListProject: %+v \n", ctx.Request)

	if list, err := biz.dao.List(); err != nil {
		util.ResponseByErr(ctx, "查询错误", err.Error())
	} else {
		log.Printf("ProjectBiz.ListProject list:%+v, err:%+v", list, err)
		util.ResponseByOk(ctx, "查询成功", &list)
	}
}

func (biz ProjectBiz) updateProject(ctx *gin.Context) {
	log.Printf("ProjectBiz.updateProject: %+v \n", ctx.Request)
	var project model.Project
	if err := ctx.Bind(&project); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	if id, err := biz.dao.UpdateById(project); err != nil {
		util.ResponseByErr(ctx, "更新错误", err.Error())
	} else {
		log.Printf("ProjectBiz.updateProject id,:%+v, err:%+v", id, err)
		util.ResponseByOk(ctx, "更新成功", id)
	}
}
