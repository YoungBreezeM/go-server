package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/log"
	"server/models"
)

func GetWXAccessToken(appId string, appsecret string, code string) models.WXAuthToken {
	r, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appId, appsecret, code))
	if err != nil {
		fmt.Println(err)
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	//
	var wxToken models.WXAuthToken
	err = json.Unmarshal(b, &wxToken)
	if err != nil {
		fmt.Println(err)
	}

	return wxToken
}

func GetWXUserInfo(accessToken string, openId string) models.WXUserInfo {
	r, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken, openId))
	if err != nil {
		log.Log.Error(err)
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Log.Error(err)
	}
	//
	var wxUser models.WXUserInfo
	err = json.Unmarshal(b, &wxUser)
	if err != nil {
		log.Log.Error(err)
	}

	return wxUser
}
