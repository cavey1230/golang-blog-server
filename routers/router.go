package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/api/v1"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	engine := gin.Default()

	routerV1 := engine.Group("/api/v1")
	{
		// 用户模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		// 文章模块的路由接口

		// 分类模块的路由接口
	}

	err := engine.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
