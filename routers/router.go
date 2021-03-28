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
		V1NeedAuth.GET("user/find", v1.FindUser)
		// 分类模块的路由接口
		V1NeedAuth.POST("category/add", v1.AddCategory)
		V1NeedAuth.PUT("category/:id", v1.EditCategory)
		V1NeedAuth.DELETE("category/:id", v1.DeleteCategory)
		V1NeedAuth.GET("category/find", v1.FindCategory)
		// 文章模块的路由接口
		V1NeedAuth.POST("article/add", v1.AddArticle)
		V1NeedAuth.PUT("article/:id", v1.EditArticle)
		V1NeedAuth.DELETE("article/:id", v1.DeleteArticle)
		V1NeedAuth.GET("article/find", v1.FindAllArticles)
		// 个人信息模块的路由接口
		V1NeedAuth.POST("info/add", v1.AddInfo)
		V1NeedAuth.PUT("info/:id", v1.EditInfo)
		V1NeedAuth.DELETE("info/:id", v1.DeleteInfo)
		V1NeedAuth.GET("info/find", v1.FindInfo)
		// 工具链接地址模块的路由接口
		V1NeedAuth.POST("tools_link/add", v1.AddToolsLink)
		V1NeedAuth.PUT("tools_link/:id", v1.EditToolsLink)
		V1NeedAuth.DELETE("tools_link/:id", v1.DeleteToolsLink)
		V1NeedAuth.GET("tools_link/find", v1.FindToolsLink)
		// 友情链接地址模块的路由接口
		V1NeedAuth.POST("blogroll/add", v1.AddBlogroll)
		V1NeedAuth.PUT("blogroll/:id", v1.EditBlogroll)
		V1NeedAuth.DELETE("blogroll/:id", v1.DeleteBlogroll)
		// 版权信息模块的路由接口
		V1NeedAuth.POST("copyright/add", v1.AddCopyright)
		V1NeedAuth.PUT("copyright/:id", v1.EditCopyright)
		V1NeedAuth.DELETE("copyright/:id", v1.DeleteCopyright)
		// 上传接口
		V1NeedAuth.POST("upload", v1.UpLoad)
	}
	V1Public := engine.Group("/api/v1/public")
	{
		// 获取用户信息
		V1Public.GET("users", v1.GetUsers)
		// 获取类型
		V1Public.GET("category", v1.GetCategory)
		// 获取所有类型 不分页
		V1Public.GET("all_category", v1.GetAllCategory)
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
		// 获取友情链接地址
		V1Public.GET("blogroll", v1.GetAllBlogroll)
		V1Public.GET("blogroll/:id", v1.GetOneBlogroll)
		// 获取版权信息
		V1Public.GET("copyright", v1.GetAllCopyright)
		V1Public.GET("copyright/:id", v1.GetOneCopyright)
		// 登录注册
		V1Public.POST("login", v1.Login)
		V1Public.POST("admin_login", v1.AdminLogin)
		V1Public.POST("user/add", v1.AddUser)
	}

	err := engine.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
