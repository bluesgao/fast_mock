package main

import (
	"fast_mock/biz"
	"fast_mock/conf"
	"fast_mock/dao"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	app := Application{}
	app.init()
	app.start()
	defer app.shutdown()
}

func setupRouter(g *gin.Engine) {
	log.Println(">>>> gin SetupRouter <<<<")

	//if _,ok:=binding.Validator.Engine().(*validator.Validate); !ok{
	//	log.Println("绑定验证器失败")
	//}

	g.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code":    200,
			"success": true,
		})
	}) //联通性接口

	//项目接口组
	pg := g.Group("/project")
	projectBiz := biz.NewProjectBiz()
	//新增接口
	pg.POST("/create", projectBiz.Create)
	//列表接口
	pg.GET("/list", projectBiz.List)
	//单个接口
	pg.GET("/detail/:id", projectBiz.GetById)

	//api接口组
	ag := g.Group("/api")
	apiBiz := biz.NewApiBiz()
	//新增接口
	ag.POST("/create", apiBiz.Create)
	//列表接口
	ag.GET("/list", apiBiz.List)
}

type Application struct {
	config *conf.Conf
	server *gin.Engine
}

func (app *Application) init() {
	log.Println(">>>> app init start <<<<")
	//配置文件
	app.config = new(conf.Conf)
	app.config.Load("conf.yaml")

	//数据库
	dao.MongoInit()

	// gin engine
	app.server = gin.Default()
	// router
	setupRouter(app.server)
}

func (app *Application) start() {
	log.Println(">>>> app start start <<<<")
	app.server.Run("localhost:" + app.config.Server.Port)
}

func (app *Application) shutdown() {
	log.Println(">>>> app shutdown start <<<<")
	defer dao.MongoClose()
}
