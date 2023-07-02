package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"server/config"
	"server/constant"
	"server/db"
	"server/log"
	"server/models"
	"server/services"
	"server/utils"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

var (
	QrCodeWaitForScan  map[string]models.QrCodeStatus
	LoginChan          chan models.QrCodeStatus
	WechatOpenIdAndKey map[string]string
)

func init() {
	QrCodeWaitForScan = make(map[string]models.QrCodeStatus, 10)
	LoginChan = make(chan models.QrCodeStatus)
	WechatOpenIdAndKey = make(map[string]string, 10)
}

// func GetWXAccessToken(c *gin.Context) {
// 	code := c.Param("code")
// 	wxUser := api.GetWXAccessTokenByCode(config.WX_APPID, config.WX_APPSECERT, code)

// 	c.JSON(200, models.R[models.WXAuthToken]{
// 		Status:  0,
// 		Data:    wxUser,
// 		Message: constant.SUCCESS,
// 	})
// }

// func GetWXUserInfo(c *gin.Context) {
// 	// 处理获取用户信息的逻辑
// 	accessToken := c.Param("access_token")
// 	openId := c.Param("openId")
// 	wxUser := api.GetWXUserInfo(accessToken, openId)
// 	c.JSON(200, models.R[models.WXUserInfo]{
// 		Status:  0,
// 		Data:    wxUser,
// 		Message: constant.SUCCESS,
// 	})
// }

// func GetWeChatQrCode(c *gin.Context) {
// 	sceneId := c.Param("sceneId")
// 	num, err := strconv.ParseUint(sceneId, 10, 32)
// 	if err != nil {
// 		c.JSON(200, models.R[string]{
// 			Status:  0,
// 			Message: "SceneId Incorrect Data Type",
// 		})
// 		return
// 	}

// 	wt := api.GetWXAccessTokenByClient(config.WX_APPID, config.WX_APPSECERT)
// 	log.Log.Debugln(wt)
// 	//
// 	qrReq := &models.WXQrCodeReq{
// 		ExpireSeconds: config.QRCODE_EXpPIRE_SECOND,
// 		ActionName:    "QR_SCENE",
// 		ActionInfo: models.ActionInfo{
// 			Scene: models.Scene{
// 				SceneID: uint32(num),
// 			},
// 		},
// 	}
// 	wcr := api.GetWXQrCode(wt.Access_token, qrReq)
// 	log.Log.Info(wcr)
// 	QrCodeWaitForScan[sceneId] = models.QrCodeStatus{
// 		Ticket: wcr.Ticket,
// 		Status: false,
// 	}
// 	//
// 	imgBytes := api.GetWXQrCodeImg(wcr.Ticket)

// 	c.Header("Content-Type", "image/png")

// 	c.Writer.Write(imgBytes)
// }

// func WechatMessageCallback(c *gin.Context) {
// 	callbackMsg := models.CallbackMsg{}
// 	if err := c.ShouldBindXML(&callbackMsg); err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	//
// 	log.Log.Info("Wechat-Notify:", callbackMsg)

// 	//subscribed and scan qr code
// 	if callbackMsg.Event == "SCAN" && callbackMsg.MsgType == "event" {
// 		//Write off QR code
// 		if value, ok := QrCodeWaitForScan[callbackMsg.EventKey]; ok && value.Ticket == callbackMsg.Ticket {
// 			//Watch QR code is scan
// 			if v, o := SubscribeQrCodeScan[callbackMsg.EventKey]; o {
// 				//gen token
// 				nowTime := time.Now()
// 				token := utils.GeneratorToken(&models.JWTClaims{
// 					OpenId: callbackMsg.FromUserName,
// 					StandardClaims: jwt.StandardClaims{
// 						ExpiresAt: nowTime.Add(time.Minute * 60).Unix(),
// 						Issuer:    "ai box",
// 						IssuedAt:  nowTime.Unix(),
// 					},
// 				})
// 				log.Log.Debugln(token)
// 				//send token to client
// 				v <- token
// 			}
// 		} else {
// 			fmt.Println("Not Found Action")
// 		}
// 	}

// 	//subscribe
// 	if callbackMsg.Event == "subscribe" && callbackMsg.MsgType == "event" {
// 		services.Subscribe(callbackMsg.FromUserName)
// 		//
// 		key := callbackMsg.EventKey[8:]
// 		if value, ok := QrCodeWaitForScan[key]; ok && value.Ticket == callbackMsg.Ticket {
// 			if value, ok := SubscribeQrCodeScan[key]; ok {
// 				//gen token
// 				nowTime := time.Now()
// 				token := utils.GeneratorToken(&models.JWTClaims{
// 					OpenId: callbackMsg.FromUserName,
// 					StandardClaims: jwt.StandardClaims{
// 						ExpiresAt: nowTime.Add(time.Minute * 60).Unix(),
// 						Issuer:    "ai box",
// 						IssuedAt:  nowTime.Unix(),
// 					},
// 				})
// 				log.Log.Debugln(token)
// 				//send token to client
// 				value <- token
// 			}
// 		} else {
// 			fmt.Println("Not Found Qr Code")
// 		}

// 	}

// 	//unsubscribe
// 	if callbackMsg.Event == "unsubscribe" && callbackMsg.MsgType == "event" {
// 		//Delete user info
// 		services.Unsubscribe(callbackMsg.FromUserName)
// 	}

// }
func Login(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(401, models.R[string]{
			Status:  401,
			Message: "not found key",
		})
		return
	}
	log.Log.Debug(key)
	//
	val, err := db.Redis.Get(c, key).Result()
	if err != nil {
		if err == redis.Nil {
			c.JSON(401, models.R[string]{
				Status:  401,
				Message: constant.UNAITHORIZED,
			})
		}
		c.JSON(401, models.R[string]{
			Status:  401,
			Message: err.Error(),
		})
		return
	}

	nowTime := time.Now()
	token := utils.GeneratorToken(&models.JWTClaims{
		OpenId: val,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: nowTime.Add(time.Minute * 60).Unix(),
			Issuer:    "FreeAiBox",
			IssuedAt:  nowTime.Unix(),
		},
	})
	db.Redis.Del(c, key)
	//
	c.JSON(200, models.R[string]{
		Status:  0,
		Data:    token,
		Message: constant.SUCCESS,
	})

}

// Watch user subscribe public number
func WatchWechatSubscribe(c *gin.Context) {
	c.Header("Content-Type", "application/xml")
	callbackMsg := models.CallbackMsg{}
	if err := c.ShouldBindXML(&callbackMsg); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	//
	log.Log.Info("Wechat-Notify:", callbackMsg)
	//subscribed and scan qr code
	if callbackMsg.Event == "subscribe" && callbackMsg.MsgType == "event" {
		services.Subscribe(callbackMsg.FromUserName)
	}

	//unsubscribe
	if callbackMsg.Event == "unsubscribe" && callbackMsg.MsgType == "event" {
		services.Unsubscribe(callbackMsg.FromUserName)
	}

	if callbackMsg.MsgType == "text" {
		if !config.SysStatus {
			replyMsg := models.TextReply{
				ToUserName:   callbackMsg.FromUserName,
				FromUserName: callbackMsg.ToUserName,
				CreateTime:   time.Now().Unix(),
				MsgType:      "text",
				Content:      "系统正在维护中，请稍后再试，谢谢理解！",
			}
			c.XML(http.StatusOK, replyMsg)
			return
		}
		key := utils.GenerateRandomString(32)
		var replyMsg models.TextReply
		//
		if callbackMsg.Content == "链接" || callbackMsg.Content == "密钥" {
			replyMsg = models.TextReply{
				ToUserName:   callbackMsg.FromUserName,
				FromUserName: callbackMsg.ToUserName,
				CreateTime:   time.Now().Unix(),
				MsgType:      "text",
				Content:      fmt.Sprintf("%s\n点开下面这个链接进入网页。\nhttp://localhost:3000/home?key=%s", config.DESC, key),
			}
			//
			db.Redis.Set(c, key, callbackMsg.FromUserName, time.Minute*60)
			c.XML(http.StatusOK, replyMsg)
		} else {
			replyMsg = models.TextReply{
				ToUserName:   callbackMsg.FromUserName,
				FromUserName: callbackMsg.ToUserName,
				CreateTime:   time.Now().Unix(),
				MsgType:      "text",
				Content:      "您输入的关键词不对哦！目前可用关键词为[链接]",
			}
			c.XML(http.StatusOK, replyMsg)
		}

		// 设置回复消息的Content-Type为XML

	}
}

// func WatchQrcodeIsScan(c *gin.Context) {
// 	//
// 	c.Header("Content-Type", "text/event-stream")
// 	if HasUserWait {
// 		r := models.R[string]{
// 			Status:  503,
// 			Message: "当前登录用户过多，请稍后再试！",
// 		}
// 		b, err := json.Marshal(r)
// 		if err != nil {
// 			log.Log.Errorln(err)
// 		}
// 		//
// 		fmt.Fprintf(c.Writer, "data: %s\n\n", b)
// 		c.Writer.Flush()
// 		return
// 	} else {

// 		HasUserWait = true
// 	}
// 	//
// 	SubscribePublucCodeCh = make(chan string)
// 	defer close(SubscribePublucCodeCh)
// 	//clear user with occupying for long time
// 	timer := time.NewTimer(10 * time.Second)
// 	go func() {
// 		<-timer.C
// 		r := models.R[string]{
// 			Status:  502,
// 			Message: "二维码已过期,请刷新页面",
// 		}
// 		b, err := json.Marshal(r)
// 		if err != nil {
// 			log.Log.Errorln(err)
// 		}
// 		fmt.Fprintf(c.Writer, "data: %s\n\n", b)
// 		c.Writer.Flush()
// 		//
// 		HasUserWait = false
// 	}()
// 	//
// 	for {
// 		select {
// 		case token := <-SubscribePublucCodeCh:
// 			log.Log.Info("send:", token)
// 			r := models.R[string]{
// 				Status:  0,
// 				Data:    token,
// 				Message: constant.SUCCESS,
// 			}
// 			b, err := json.Marshal(r)
// 			if err != nil {
// 				log.Log.Errorln(err)
// 			}
// 			fmt.Fprintf(c.Writer, "data: %s\n\n", b)
// 			c.Writer.Flush()
// 			//
// 			HasUserWait = false
// 		case <-c.Writer.CloseNotify():
// 			log.Log.Info("login out")
// 			HasUserWait = false
// 		}
// 	}

// }

func WechatCheckToken(c *gin.Context) {
	queryValues := c.Request.URL.Query()

	signature := queryValues.Get("signature")
	timestamp := queryValues.Get("timestamp")
	nonce := queryValues.Get("nonce")
	echoStr := queryValues.Get("echostr")

	// 将 token、timestamp 和 nonce 放入一个切片中，并按照字典顺序排序
	params := []string{config.Cfg.WeChat.Token, timestamp, nonce}
	sort.Strings(params)

	// 将排序后的参数拼接成一个字符串
	str := strings.Join(params, "")

	// 对拼接后的字符串进行 SHA1 计算
	h := sha1.New()
	io.WriteString(h, str)
	sha1Sum := hex.EncodeToString(h.Sum(nil))

	// 将计算得到的 SHA1 值与 signature 进行比较
	if sha1Sum == signature {
		// Token 验证通过，返回 echostr 参数
		c.String(http.StatusOK, echoStr)
	} else {
		// Token 验证失败
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
