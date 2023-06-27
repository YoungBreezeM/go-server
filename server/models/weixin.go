package models

import "encoding/xml"

type WXUserInfo struct {
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Language   string   `json:"language"`
	Nickname   string   `json:"nickname"`
	Privilege  []string `json:"privilege"`
	Province   string   `json:"province"`
	Sex        int8     `json:"sex"`
}

type WXAuthToken struct {
	Access_token string `json:"access_token"`
	Expires_in   int16  `json:"expires_in"`
	Openid       string `json:"openid"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type WXQrCodeReq struct {
	ExpireSeconds int        `json:"expire_seconds"`
	ActionName    string     `json:"action_name"`
	ActionInfo    ActionInfo `json:"action_info"`
}

type ActionInfo struct {
	Scene Scene `json:"scene"`
}

type Scene struct {
	SceneID uint32 `json:"scene_id"`
}

type WXQrCodeRes struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	URL           string `json:"url"`
}

type QrCodeStatus struct {
	Ticket       string `json:"ticket"`
	Access_token string `json:"access_token"`
	Status       bool   `json:"status"`
	SceneID      string `json:"scene_id"`
	OpenId       string `json:"openId"`
}

type CallbackMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Event        string   `xml:"Event"`
	EventKey     string   `xml:"EventKey"`
	Ticket       string   `xml:"Ticket"`
	Content      string   `xml:"Content"`
}

type TextReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}
