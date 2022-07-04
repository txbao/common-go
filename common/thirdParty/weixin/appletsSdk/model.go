package appletsSdk

//获取openid
type Jscode2sesssionStruct struct {
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

//微信小程序 获取手机号
type GetPhoneNumberStruct struct {
	CountryCode     string `json:"countryCode"`
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	Watermark       struct {
		Appid     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}
