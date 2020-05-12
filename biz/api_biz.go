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
	mgo      *dao.MongoCli
	database string
	coll     string
}

func NewApiBiz() ApiBiz {
	biz := ApiBiz{}
	biz.mgo = dao.GetMongoCli()
	biz.database = "mock"
	biz.coll = "mock_api"
	return biz
}

func (biz ApiBiz) Create(ctx *gin.Context) {
	log.Printf("ApiBiz.Create: %+v \n", ctx.Request)
	var input dto.ProjectDto

	if err := ctx.Bind(&input); err != nil {
		log.Println(err.Error())
		util.FAIL(ctx, "参数错误", err.Error())
		return
	}
	log.Printf("ApiBiz.Create input: %+v \n", input)

	ret, err := biz.mgo.InsertOne(nil, biz.database, biz.coll, input)
	if err != nil {
		log.Printf("ApiBiz.Create err: %+v \n", err)
		util.FAIL(ctx, "创建错误", err.Error())
		return
	} else {
		log.Printf("ApiBiz.Create ret: %+v \n", ret)
		util.SUCCESS(ctx, ret)

	}
}

func (biz ApiBiz) List(ctx *gin.Context) {
	log.Printf("ApiBiz.List: %+v \n", ctx.Request)
	pn := ctx.Query("projectName") //查询请求URL后面的参数

	log.Printf("ApiBiz.List input: %+v \n", pn)

	query := bson.M{
		"projectname": pn,
	}
	ret, err := biz.mgo.FindMany(nil, biz.database, biz.coll, query)
	if err != nil {
		log.Printf("ApiBiz.List err: %+v \n", err)
		util.FAIL(ctx, "创建错误", err.Error())
		return
	} else {
		log.Printf("ApiBiz.List ret: %+v \n", ret)
		util.SUCCESS(ctx, ret)

	}
}
