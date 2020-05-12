package biz

import (
	"fast_mock/dao"
	"fast_mock/dto"
	"fast_mock/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type ApiBiz struct {
	mgo      dao.Mgo
	database string
	coll     string
}

func NewApiBiz() ApiBiz {
	biz := ApiBiz{}
	biz.mgo = dao.NewMgo("mongodb://admin:123456@47.97.205.190:27017")
	biz.database = "mock"
	biz.coll = "mock_api"
	return biz
}

func (biz ApiBiz) CreateApi(ctx *gin.Context) {
	log.Printf("ApiBiz.CreateApi: %+v \n", ctx.Request)
	var input dto.ProjectDto

	if err := ctx.Bind(&input); err != nil {
		log.Println(err.Error())
		util.FAIL(ctx, "参数错误", err.Error())
		return
	}
	log.Printf("ApiBiz.CreateApi input: %+v \n", input)

	ret, err := biz.mgo.InsertOne(nil, biz.database, biz.coll, input)
	if err != nil {
		log.Printf("ApiBiz.CreateApi err: %+v \n", err)
		util.FAIL(ctx, "创建错误", err.Error())
		return
	} else {
		log.Printf("ApiBiz.CreateApi ret: %+v \n", ret)
		util.SUCCESS(ctx, ret)

	}
}

func (biz ApiBiz) ListApi(ctx *gin.Context) {
	log.Printf("ApiBiz.ListApi: %+v \n", ctx.Request)
	pn := ctx.Query("projectName") //查询请求URL后面的参数

	log.Printf("ApiBiz.ListApi input: %+v \n", pn)

	query := bson.M{
		"projectname": pn,
	}
	ret, err := biz.mgo.FindMany(nil, biz.database, biz.coll, query)
	if err != nil {
		log.Printf("ApiBiz.ListApi err: %+v \n", err)
		util.FAIL(ctx, "创建错误", err.Error())
		return
	} else {
		log.Printf("ApiBiz.ListApi ret: %+v \n", ret)
		util.SUCCESS(ctx, ret)

	}
}
