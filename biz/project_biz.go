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
	biz := ProjectBiz{
		dao: dao.NewProjectDao(),
	}
	return biz

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

	id, err := biz.dao.Create(project)
	log.Printf("id:%d, err:%s", id, err)
	if err != nil {
		util.ResponseByErr(ctx, "创建错误", err.Error())
		return
	}

	project.Id = id
	util.ResponseByOk(ctx, "创建成功", &project)
	return
}

func (biz ProjectBiz) ListProject(ctx *gin.Context) {
	log.Printf("ProjectBiz.ListProject: %+v \n", ctx.Request)

	list, err := biz.dao.List()
	log.Printf("ProjectBiz.ListProject list:%+v, err:%+v", list, err)
	if err != nil {
		util.ResponseByErr(ctx, "查询错误", err.Error())
		return
	}
	util.ResponseByOk(ctx, "查询成功", &list)
	return
}

func (biz ProjectBiz) updateProject(ctx *gin.Context) {
	log.Printf("ProjectBiz.updateProject: %+v \n", ctx.Request)
	var project model.Project
	if err := ctx.Bind(&project); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	id, err := biz.dao.UpdateById(project)
	log.Printf("ProjectBiz.updateProject id,:%+v, err:%+v", id, err)
	if err != nil {
		util.ResponseByErr(ctx, "更新错误", err.Error())
		return
	}
	util.ResponseByOk(ctx, "更新成功", id)
	return
}
