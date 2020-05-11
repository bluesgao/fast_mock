package biz

import (
	"fast_mock/dao"
	"fast_mock/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		ctx.JSON(http.StatusNotFound, gin.H{"code": "9999", "msg": "参数错误", "data": err.Error()})
		return
	}
	log.Printf("CreateProject project: %+v \n", project)

	ret, err := biz.dao.Create(project)
	log.Printf("ret:%d, err:%s", ret, err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "9999", "msg": "创建错误", "data": &project})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"code": "0000", "msg": "创建成功", "data": &project})
	return
}
