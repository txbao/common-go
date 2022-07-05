package sms

import (
	"common-go/common/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"unicode/utf8"
)

const (
	SmsSignInSecond = 300
)

type CacheCode struct {
	Code   string `json:"code"`
	Mobile string `json:"mobile"`
}

type SmsData struct {
	Mobile string `json:"mobile"`
	SmsKey string `json:"sms_key"`
}

type SmsSdk struct {
	rds          *redis.Redis
	client       *dysmsapi.Client
	signName     string
	templateCode string
}

func NewSmsSdk(rds *redis.Redis, regionId, accessKyeId, accessSecret string, signName string, templateCode string) (*SmsSdk, error) {
	client, err := dysmsapi.NewClientWithAccessKey(regionId, accessKyeId, accessSecret)
	if err != nil {
		return nil, err
	}
	return &SmsSdk{
		rds:          rds,
		client:       client,
		signName:     signName,
		templateCode: templateCode,
	}, nil
}

//发送验证码
func (o *SmsSdk) SendSms(data *SmsData) error {
	code := utils.RandomString(6, 1)

	mobile := data.Mobile
	smsKey := data.SmsKey

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = o.signName
	request.TemplateCode = o.templateCode
	request.TemplateParam = fmt.Sprintf("{\"code\":%s}", code)
	response, err := o.client.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Println("SendSmsCode"+mobile+":", "code:", code, "mobile:", mobile)
	//  &dysmsapi.SendSmsResponse{BaseResponse:(*responses.BaseResponse)(0xc00004ecc0), RequestId:"B221D9E7-3DD3-40B9-A7EA-8DD1621C815C", BizId:"876614788237204929^0", Code:"OK", Message:"OK"}
	fmt.Printf("response is %#v\n", response)
	if response.Code == "OK" {
		// 发送成功后缓存验证码到redis
		cacheVal := &CacheCode{Code: code, Mobile: mobile}
		jsonStr, err := json.Marshal(cacheVal)
		if err != nil {
			return err
		}
		err = o.rds.Setex(smsKey+mobile, string(jsonStr), SmsSignInSecond)
		if err != nil {
			return err
		}
		err = o.rds.Setex("bank_sms_"+mobile, smsKey, SmsSignInSecond)
		if err != nil {
			return err
		}
	}
	return nil
}

// VerifySmsCode 校验手机验证码
func (o *SmsSdk) VerifySmsCode(mobile string, key string, code string) error {
	if key == "" || utf8.RuneCountInString(code) != 6 {
		return errors.New("无效的验证码")
	}
	key = key + mobile
	cacheCode, err := o.rds.Get(key)
	if cacheCode == "" {
		return errors.New("验证码已失效")
	}

	var codeRes CacheCode
	err = json.Unmarshal([]byte(cacheCode), &codeRes)
	if err != nil {
		return errors.New("转码失败")
	}
	if codeRes.Code != code {
		return errors.New("验证码错误")
	}
	_, _ = o.rds.Del(key)
	_, _ = o.rds.Del("bank_sms_" + mobile)
	return nil
}
