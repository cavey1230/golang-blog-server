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
	engine := gin.New()
	engine.Use(middleware.WriteLog())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Cors2())

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
		// 个人信息模块的路由接口
		V1NeedAuth.POST("info/add", v1.AddInfo)
		V1NeedAuth.PUT("info/:id", v1.EditInfo)
		V1NeedAuth.DELETE("info/:id", v1.DeleteInfo)
		// 工具链接地址模块的路由接口
		V1NeedAuth.POST("tools_link/add", v1.AddToolsLink)
		V1NeedAuth.PUT("tools_link/:id", v1.EditToolsLink)
		V1NeedAuth.DELETE("tools_link/:id", v1.DeleteToolsLink)
		// 上传接口
		V1NeedAuth.POST("upload", v1.UpLoad)
	}
	V1Public := engine.Group("/api/v1/public")
	{
		// 获取用户信息
		V1Public.GET("users", v1.GetUsers)
		// 获取类型
		V1Public.GET("category", v1.GetCategory)
		// 获取文章信息
		V1Public.GET("article", v1.GetAllArticles)
		V1Public.GET("article/:id", v1.GetOneArticle)
		// 获取精品文章信息
		V1Public.GET("boutique_article", v1.GetAllBoutiqueArticles)
		// 获取个人信息
		V1Public.GET("info", v1.GetAllInfo)
		V1Public.GET("info/:id", v1.GetOneInfo)
		// 获取工具链接地址
		V1Public.GET("tools_link", v1.GetAllToolsLink)
		V1Public.GET("tools_link/:id", v1.GetOneToolsLink)
		// 登录注册
		V1Public.POST("login", v1.Login)
		V1Public.POST("user/add", v1.AddUser)
	}

	err := engine.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
