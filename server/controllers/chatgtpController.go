package controllers

import (
	"fmt"
	"server/api"
	"server/constant"
	"server/log"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

var ChatGPTResChan map[string]chan []byte

func init() {
	ChatGPTResChan = make(map[string]chan []byte, 10)
}

// @Summary chatgtp send msg
// @Description by openId and content
// @Tags Auth
// @Accept json
// @Produce json
// @Param code path uint true "chat gtp"
// @Success 200 {object} .
// @Router /chatgpt/chat [post]
// func Chat(c *gin.Context) {

// 	var cq api.APIRequest
// 	if err := c.BindJSON(&cq); err != nil {
// 		log.Log.Error(err)
// 	}
// 	//

// 	token := c.GetHeader("Authorization")
// 	tokenClaims, err := jwt.ParseWithClaims(token, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return config.JWTSECRET, nil
// 	})
// 	//
// 	if err != nil {
// 		log.Log.Errorln(err)
// 	}
// 	//
// 	if tokenClaims != nil {
// 		if claims, ok := tokenClaims.Claims.(*models.JWTClaims); ok && tokenClaims.Valid {
// 			user, err := db.GetUserInfoByOpenId(claims.OpenId)
// 			if err != nil {
// 				log.Log.Errorln(err)
// 				return
// 			}
// 			//send msg to chatgtp by api
// 			log.Log.Info(user)
// 			if user.Integral > 0 {
// 				go cq.SendMsg(user, ChatGPTResChan)

// 				c.JSON(200, models.R[models.User]{
// 					Status:  0,
// 					Data:    *user,
// 					Message: constant.SUCCESS,
// 				})

// 			} else {
// 				c.JSON(402, models.R[string]{
// 					Status:  2,
// 					Message: constant.NOT_ENOUGH_POINTS,
// 				})
// 			}
// 		}

// 	}

// }

func Gtp4Chat(c *gin.Context) {
	var cq api.APIRequest
	if err := c.BindJSON(&cq); err != nil {
		log.Log.Error(err)
	}
	//
	msg := api.NewGTP4Request()
	msg.ClientId = utils.GenerateRandomString(12)
	msg.ContextId = cq.ChatId
	msg.NewMessage = cq.Messages[0].Content
	//
	go msg.SendMsg(ChatGPTResChan)

	c.JSON(200, models.R[string]{
		Status:  0,
		Message: constant.SUCCESS,
	})

}

func ChatGTPCallback(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	// 设置 SSE 缓冲区大小为 16KB
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")
	c.Header("X-Accel-Buffering", "no") // 禁用 Nginx 缓冲
	c.Header("Buffer-Size", "16384")    // 设置缓冲区大小为 16KB
	//
	chatId := c.Param("chatId")
	ChatGPTResChan[chatId] = make(chan []byte)
	for {
		select {
		case event := <-ChatGPTResChan[chatId]:
			fmt.Fprintf(c.Writer, "%s\n\n", event)
			c.Writer.Flush()
		case <-c.Writer.CloseNotify():
			log.Log.Info("chatId close")
		}

	}

}
