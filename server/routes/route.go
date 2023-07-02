package routes

import (
	"server/constant"
	"server/controllers"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "server/docs"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) <= 0 {
		c.JSON(200, models.R[string]{
			Status:  401,
			Message: constant.UNAITHORIZED,
		})
		c.Abort()
		return
	}
	//
	if !services.VerifyToken(token) {
		c.JSON(200, models.R[string]{
			Status:  401,
			Message: constant.UNAITHORIZED,
		})
		c.Abort()
		return
	}
	//
	c.Next()

}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// wechat handler
	// router.GET("/wx/auth/:code", controllers.GetWXAccessToken)
	// router.GET("/wx/user/:access_token/:openId", controllers.GetWXUserInfo)
	// router.GET("/wx/qrCode/:sceneId", controllers.GetWeChatQrCode)
	router.POST("/notify", controllers.WatchWechatSubscribe)
	router.GET("/notify", controllers.WechatCheckToken)
	router.GET("/loginByKey", controllers.Login)
	// router.GET("/watchQrCodeScan", controllers.WatchQrcodeIsScan)
	//
	router.GET("/chatgtp/msg/callback/:chatId", controllers.ChatGTPCallback)
	//middleware
	router.Use(AuthMiddleware)
	// router.POST("/v1/chat/completions", controllers.Chat)
	router.POST("/v2/chat/completions", controllers.Gtp4Chat)
	//
	// router.GET("/user", controllers.GetUserInfo)
	// router.PUT("/users/:id", controllers.UpdateUserInfo)

	return router
}
