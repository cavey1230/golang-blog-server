package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func Cors() gin.HandlerFunc  {
//	return func(context *gin.Context) {
//		cors.New(cors.Config{
//			//AllowAllOrigins: true,
//			AllowOrigins:     []string{"*"},//允许跨域的域名
//			AllowMethods:     []string{"*"},//允许跨域的方式
//			AllowHeaders:     []string{"Origin"},//允许跨域的请求头
//			ExposeHeaders:    []string{"Content-Length","Authorization"}, //允许跨域的请求字段
//			//AllowCredentials: true, cook跨域
//			//AllowOriginFunc: func(origin string) bool {
//			//	return origin == "https://github.com"
//			//},
//			MaxAge: 12 * time.Hour,
//		})
//	}
//}
func Cors2() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//c.Header("Access-Control-Allow-Credentials", "true") //允许前端携带cook

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
