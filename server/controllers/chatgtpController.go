package controllers

import (
	"fmt"
	"server/api"
	"server/constant"
	"server/log"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
)

var ChatGPTResChan map[string]chan []byte

func init() {
	ChatGPTResChan = make(map[string]chan []byte, 10)
}

func Gtp4Chat(c *gin.Context) {
	var cq api.APIRequest
	if err := c.BindJSON(&cq); err != nil {
		log.Log.Error(err)
	}
	//
	msg := api.NewGTPRequest()
	msg.Messages = cq.Messages
	msg.ChatId = cq.ChatId
	//
	go msg.SendMsg(ChatGPTResChan)

	services.RecordChat(c, msg.ChatId, cq.Messages)

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
