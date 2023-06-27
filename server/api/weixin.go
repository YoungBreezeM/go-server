package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/log"
	"server/models"
	"strings"
)

func GetWXAccessTokenByCode(appId string, appsecret string, code string) models.WXAuthToken {
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

func GetWXAccessTokenByClient(appId string, appsecret string) models.WXAuthToken {
	r, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?appid=%s&secret=%s&grant_type=client_credential", appId, appsecret))
	if err != nil {
		fmt.Println(err)
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
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

func GetWXQrCode(accessToken string, qr *models.WXQrCodeReq) *models.WXQrCodeRes {
	json_qr, err := json.Marshal(qr)
	if err != nil {
		log.Log.Error(err)
	}
	//
	r, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", accessToken), "application/json", strings.NewReader(string(json_qr)))
	if err != nil {
		log.Log.Error(err)
	}
	//
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Log.Error(err)
	}
	//
	if r.StatusCode != 200 {
		log.Log.Info(string(b))
		return nil
	}
	//
	var wxQrCodeRes models.WXQrCodeRes
	err = json.Unmarshal(b, &wxQrCodeRes)
	if err != nil {
		log.Log.Error(err)
	}

	return &wxQrCodeRes
}

func GetWXQrCodeImg(ticket string) []byte {
	//
	r, err := http.Get(fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", ticket))
	if err != nil {
		log.Log.Error(err)
	}
	//
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Log.Error(err)
	}

	return b
}
