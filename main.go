package main

import (
	"goblog/model"
	"goblog/routers"
)

func main() {
	//引用数据库
	model.InitDb()
	//路由
	routers.InitRouter()
}
