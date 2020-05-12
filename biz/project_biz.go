package biz

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ProjectBiz struct {
}

func NewProjectBiz() ProjectBiz {
	return ProjectBiz{}
}

func (biz ProjectBiz) Create(ctx *gin.Context) {
	log.Printf("ProjectBiz.Create: %+v \n", ctx.Request)
}

func (biz ProjectBiz) List(ctx *gin.Context) {
	log.Printf("ProjectBiz.List: %+v \n", ctx.Request)
}

func (biz ProjectBiz) Update(ctx *gin.Context) {
	log.Printf("ProjectBiz.Update: %+v \n", ctx.Request)

}

func (biz ProjectBiz) GetById(ctx *gin.Context) {
	log.Printf("ProjectBiz.GetById: %+v \n", ctx.Request)
}
