package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	engine := gin.Default()

	V1NeedAuth := engine.Group("/api/v1")
	V1NeedAuth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		V1NeedAuth.PUT("user/:id", v1.EditUser)
		V1NeedAuth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		V1NeedAuth.POST("category/add", v1.AddCategory)
		V1NeedAuth.PUT("category/:id", v1.EditCategory)
		V1NeedAuth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		V1NeedAuth.POST("article/add", v1.AddArticle)
		V1NeedAuth.PUT("article/:id", v1.EditArticle)
		V1NeedAuth.DELETE("article/:id", v1.DeleteArticle)
	}
	V1Public := engine.Group("/api/v1/public")
	{
		V1Public.GET("users", v1.GetUsers)
		V1Public.GET("category", v1.GetCategory)
		V1Public.GET("article", v1.GetAllArticles)
		V1Public.GET("article/:id", v1.GetOneArticle)
		V1Public.POST("login", v1.Login)
		V1Public.POST("user/add", v1.AddUser)
	}

	err := engine.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
