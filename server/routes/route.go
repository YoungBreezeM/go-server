package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "server/docs"
)

func AuthMiddleware(c *gin.Context) {
	// 处理身份验证逻辑
	// 检查是否存在有效的身份验证令牌，如果不存在或无效，返回未经授权的错误响应
	// 如果身份验证通过，可以将用户信息存储到上下文中以供后续处理程序使用
	// 示例：
	// user, err := services.ValidateToken(c.GetHeader("Authorization"))
	// if err != nil {
	//     c.JSON(401, gin.H{
	//         "error": "Unauthorized",
	//     })
	//     c.Abort()
	//     return
	// }
	// c.Set("user", user)

	c.Next()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 添加路由处理程序
	router.GET("/users/:id", controllers.GetUserInfo)
	router.PUT("/users/:id", controllers.UpdateUserInfo)

	return router
}
