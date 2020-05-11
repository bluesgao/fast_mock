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
	log.Printf("CreateProject: %+v \n", ctx.Request)
	var project model.Project

	if err := ctx.Bind(&project); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	log.Printf("CreateProject project: %+v \n", project)

	ret, err := biz.dao.Create(project)
	log.Printf("ret:%d, err:%s", ret, err)
	if err != nil {
		util.ResponseByErr(ctx, "创建错误", err.Error())
		return
	}
	util.ResponseByOk(ctx, "创建成功", &project)
	return
}
