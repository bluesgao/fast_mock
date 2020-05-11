package biz

import (
	"fast_mock/dao"
	"fast_mock/model"
	"fast_mock/util"
	"github.com/gin-gonic/gin"
	"log"
)

type ModuleBiz struct {
	dao dao.ModuleDao
}

func NewModuleBiz() ModuleBiz {
	biz := ModuleBiz{
		dao: dao.NewModuleDao(),
	}
	return biz

}

func (biz ModuleBiz) CreateModule(ctx *gin.Context) {
	log.Printf("ModuleBiz.CreateModule: %+v \n", ctx.Request)
	var module model.Module

	if err := ctx.Bind(&module); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	log.Printf("ModuleBiz.CreateModule module: %+v \n", module)

	if id, err := biz.dao.Create(module); err != nil {
		util.ResponseByErr(ctx, "创建错误", err.Error())
	} else {
		log.Printf("id:%d, err:%s \n", id, err)
		module.Id = id
		util.ResponseByOk(ctx, "创建成功", &module)
	}
}

func (biz ModuleBiz) ListModule(ctx *gin.Context) {
	log.Printf("ModuleBiz.ListModule: %+v \n", ctx.Request)
	if list, err := biz.dao.List(); err != nil {
		util.ResponseByErr(ctx, "查询错误", err.Error())
	} else {
		log.Printf("ModuleBiz.ListModule list:%+v, err:%+v \n", list, err)
		util.ResponseByOk(ctx, "查询成功", &list)
	}
}

func (biz ModuleBiz) updateModule(ctx *gin.Context) {
	log.Printf("ModuleBiz.updateModule: %+v \n", ctx.Request)
	var module model.Module
	if err := ctx.Bind(&module); err != nil {
		log.Println(err.Error())
		util.ResponseByErr(ctx, "参数错误", err.Error())
		return
	}
	if id, err := biz.dao.UpdateById(module); err != nil {
		util.ResponseByErr(ctx, "更新错误", err.Error())
	} else {
		log.Printf("ModuleBiz.updateModule id:%+v, err:%+v \n", id, err)
		util.ResponseByOk(ctx, "更新成功", id)
	}
}
