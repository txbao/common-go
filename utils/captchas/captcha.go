package captchas

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"math/rand"
	"time"

	"github.com/mojocn/base64Captcha"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

type captcha struct {
	store RedisStore
}

func NewCaptcha(rds *redis.Redis) captcha {
	return captcha{
		store: RedisStore{rds: rds},
	}
}

//var store = RedisStore{}

//var store = base64Captcha.DefaultMemStore

// 使用 digit
// key 为用户标识，比如手机号，token, openid等
func (obj *captcha) Generate(key string) (id string, b64s string, err error) {
	//parse request parameters
	param := configJsonBody{
		Id:          "",
		CaptchaType: "digit",
		VerifyValue: "",
		DriverDigit: &base64Captcha.DriverDigit{
			Height:   40,
			Width:    120,
			Length:   4,
			MaxSkew:  0.6,
			DotCount: 10,
		},
	}
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, obj.store)

	// 默认的方式下，id是随机生产的，现在改成定制的id
	//id, b64s, err := c.Generate()

	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}

	// 使用key代替id来存储验证码
	//c.Store.Set(key, answer)
	obj.store.Set(id, answer)

	b64s = item.EncodeB64string()

	return id, b64s, err
}

func (obj *captcha) Verify(id string, VerifyValue string) (res bool) {
	//parse request json body
	param := configJsonBody{
		Id:          id,
		CaptchaType: "digit",
		VerifyValue: VerifyValue,
		DriverDigit: &base64Captcha.DriverDigit{
			Height:   40,
			Width:    120,
			Length:   4,
			MaxSkew:  0.7,
			DotCount: 80,
		},
	}
	//verify the captcha
	return obj.store.Verify(param.Id, param.VerifyValue, true)
}
