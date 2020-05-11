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
	app := application{}
	app.init()
	app.start()
	defer app.shutdown()
}

type application struct {
	config   *conf.Conf
	server   *gin.Engine
	database *dao.Database
}

func (app *application) init() {
	log.Println(">>>> app init start <<<<")
	//配置文件
	app.config = new(conf.Conf)
	app.config.Load("conf.yaml")

	//数据库
	app.database = new(dao.Database)
	app.database.Init()

	// gin engin
	app.server = gin.Default()
	// router
	setupRouter(app.server)
}

func (app *application) start() {
	log.Println(">>>> app start start <<<<")
	app.server.Run("localhost:8080")
}

func (app *application) shutdown() {
	log.Println(">>>> app shutdown start <<<<")

	defer app.database.Db.Close()
}

func setupRouter(g *gin.Engine) {
	log.Println(">>>> gin SetupRouter <<<<")

	//if _,ok:=binding.Validator.Engine().(*validator.Validate); !ok{
	//	log.Println("绑定验证器失败")
	//}

	g.GET("/hello", Hello) //联通性接口

	group := g.Group("/project") //项目接口
	pb := biz.NewProjectBiz()
	group.POST("/create", pb.CreateProject) //新增接口
}

func Hello(context *gin.Context) {
	log.Println(">>>> hello gin start <<<<")
	context.JSON(200, gin.H{
		"code":    200,
		"success": true,
	})
}