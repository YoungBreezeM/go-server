package config

var (
	WX_APPID              = "wxc56c8ba374294c4d"
	WX_APPSECERT          = "74680cd7ed6f52c119351f88f9612582"
	TOKEN                 = "40ecd7c78e67cdf292b6213ac3a3d8ac"
	QRCODE_EXpPIRE_SECOND = 60
)

type WeChat struct {
	AppId     string `yaml:"appId"`
	AppSecert string `yaml:"appSecert"`
	Token     string `yaml:"token"`
}
