package services

import (
	"encoding/json"
	"fmt"
	"server/api"
	"server/db"
	"server/log"
	"time"

	"github.com/gin-gonic/gin"
)

func RecordChat(c *gin.Context, chatId string, msg []api.ApiMessage) {
	token := c.GetHeader("Authorization")
	openId := token[32:]
	key := fmt.Sprintf("%s:%s", openId, chatId)
	b, err := json.Marshal(msg)
	if err != nil {
		log.Log.Errorln(err)
	}
	//
	s, err := db.Redis.Set(c, key, string(b), time.Hour*24*30).Result()
	if err != nil {
		log.Log.Errorln(err)
	}
	fmt.Println(s)
}
