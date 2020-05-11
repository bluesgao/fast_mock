package main

import (
	"fast_mock/conf"
	"fast_mock/dao"
	"fmt"
)

func main() {
	c := conf.Conf{}
	c.Load("conf/conf.yaml")
	fmt.Println("confg:%+v", c)
	dao.InitDb()
	dao.NewProjectDao()
}
