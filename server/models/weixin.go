package models

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
