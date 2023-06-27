package models

type ChatRequest struct {
	Content string `json:"content"`
	OpenId  string `json:"openId"`
}
