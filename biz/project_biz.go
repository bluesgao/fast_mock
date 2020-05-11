package biz

import (
	"fast_mock/dao"
	"fast_mock/dto"
	"fast_mock/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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
	var input dto.ProjectDto

	if err := ctx.Bind(&input); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	log.Printf("ProjectBiz.CreateProject input: %+v \n", input)

	if id, err := biz.dao.Create(dto.ToProjectModel(input)); err != nil {
		util.ResponseByErr(ctx, "创建错误", err.Error())
	} else {
		log.Printf("id:%d, err:%s", id, err)
		input.Id = id
		util.ResponseByOk(ctx, "创建成功", &input)
	}
}

func (biz ProjectBiz) ListProject(ctx *gin.Context) {
	log.Printf("ProjectBiz.ListProject: %+v \n", ctx.Request)

	if list, err := biz.dao.List(); err != nil {
		util.ResponseByErr(ctx, "查询错误", err.Error())
	} else {
		log.Printf("ProjectBiz.ListProject list:%+v, err:%+v \n", list, err)
		util.ResponseByOk(ctx, "查询成功", &list)
	}
}

func (biz ProjectBiz) UpdateProject(ctx *gin.Context) {
	log.Printf("ProjectBiz.UpdateProject: %+v \n", ctx.Request)
	var input dto.ProjectDto
	if err := ctx.Bind(&input); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	if id, err := biz.dao.UpdateById(dto.ToProjectModel(input)); err != nil {
		util.ResponseByErr(ctx, "更新错误", err.Error())
	} else {
		log.Printf("ProjectBiz.UpdateProject id:%+v, err:%+v \n", id, err)
		util.ResponseByOk(ctx, "更新成功", id)
	}
}

func (biz ProjectBiz) GetProjectById(ctx *gin.Context) {
	log.Printf("ProjectBiz.GetProjectById: %+v \n", ctx.Request)
	if id, ok := ctx.Params.Get("id"); !ok {
		util.ResponseByErr(ctx, "参数错误"+"id不能为空", nil)
		return
	} else {
		if projectId, err := strconv.ParseInt(id, 10, 64); err != nil {
			util.ResponseByErr(ctx, "参数错误"+"id必须是数字", nil)
			return
		} else {
			log.Printf("ProjectBiz.GetProjectById projectId: %+v \n", projectId)
			if ret, err := biz.dao.GetOne(projectId); err != nil {
				util.ResponseByErr(ctx, "查询错误", err.Error())
			} else {
				log.Printf("ProjectBiz.GetProjectById ret:%+v, err:%+v \n", ret, err)
				util.ResponseByOk(ctx, "查询成功", dto.ToProjectDto(ret))
			}
		}
	}
}
