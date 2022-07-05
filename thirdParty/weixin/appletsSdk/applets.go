package appletsSdk

import (
	"common-go/utils"
	"common-go/utils/logs"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

//微信相关api
//txbao

type _wx struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func NewSdk(appId string, appSecret string) *_wx {
	return &_wx{
		AppId:     appId,
		AppSecret: appSecret,
	}
}

func (o *_wx) Jscode2session(jsCode string) (*Jscode2sesssionStruct, error) {
	var mapData = make(map[string]interface{})
	mapData["appid"] = o.AppId
	mapData["secret"] = o.AppSecret
	mapData["js_code"] = jsCode
	mapData["grant_type"] = "authorization_code"

	weixinUrl := "https://api.weixin.qq.com/sns/jscode2session"

	res, err := utils.HttpGetResty2(weixinUrl, mapData)
	if err != nil {
		logx.Error("远程请求错误：URL："+weixinUrl+"?"+utils.Map2QueryString(mapData), "错误："+utils.GetErrorMsg(err))
		return nil, err
	}
	utils.PrintLnLog("jscode2session返回数据：AppId:"+o.AppId, res, "URL:", weixinUrl)

	var jscode2session Jscode2sesssionStruct
	err = json.Unmarshal([]byte(res), &jscode2session)
	if err != nil {
		return nil, err
	}
	if jscode2session.Errmsg != "" {
		return nil, errors.New(jscode2session.Errmsg)
	}
	return &jscode2session, nil
}

func (o *_wx) AccessToken(jsCode string) map[string]interface{} {
	var mapData = make(map[string]interface{})
	mapData["appid"] = o.AppId
	mapData["secret"] = o.AppSecret
	mapData["grant_type"] = "client_credential"

	weixinUrl := "https://api.weixin.qq.com/cgi-bin/token"

	res, err := utils.HttpGetResty2(weixinUrl, mapData)
	if err != nil {
		logx.Error("远程请求错误：URL："+weixinUrl+"?"+utils.Map2QueryString(mapData), "错误："+utils.GetErrorMsg(err))
	}

	utils.PrintLnLog("token返回数据：AppId:"+o.AppId, res)

	resMapData := utils.Json2map(res)
	return resMapData
}

//获取解密数据unionId
func (o *_wx) DecryptUserInfoData(sessionKey string, encryptedData string, iv string) string {
	getPhoneNumberStruct, err := o.WXBizDataCrypt(encryptedData, iv, sessionKey)
	getPhoneNumberJson, _ := utils.StructToJsonStr(getPhoneNumberStruct)
	logs.Info("授权登录时获取的解密数据:", getPhoneNumberJson)
	if err != nil {
		return ""
	}
	return getPhoneNumberStruct.PurePhoneNumber
}

/**
 * error code 说明.
 * <ul>

 *    <li>-41001: encodingAesKey 非法</li>
 *    <li>-41003: aes 解密失败</li>
 *    <li>-41004: 解密后得到的buffer非法</li>
 *    <li>-41005: base64加密失败</li>
 *    <li>-41016: base64解密失败</li>
 * </ul>
 */
func (o *_wx) WXBizDataCrypt(encryptedData string, iv string, sessionKey string) (*GetPhoneNumberStruct, error) {
	esKey, err1 := utils.Base64Decode(sessionKey)        //密钥
	aesIV, err2 := utils.Base64Decode(iv)                //偏移量
	aesCipher, err3 := utils.Base64Decode(encryptedData) //密文
	if err1 != nil || err2 != nil || err3 != nil {
		return nil, errors.New("base64解密出错！")
	}

	//AES-128-CBC解密
	cipherBlock, err := aes.NewCipher(esKey)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(cipherBlock, aesIV)
	origData := make([]byte, len(aesCipher))
	blockMode.CryptBlocks(origData, aesCipher)
	//填充PKCS#7
	decrypted := o.unPad(origData)
	var getPhoneNumberStruct GetPhoneNumberStruct
	err = json.Unmarshal(decrypted, &getPhoneNumberStruct)
	if err != nil {
		logs.Error("GetPhoneNumberStruct解析错误：decrypted:", string(decrypted), string(origData))
		return nil, err
	}
	if getPhoneNumberStruct.Watermark.Appid != o.AppId {
		return nil, errors.New(fmt.Sprintf("数据水印appid不同%v|%v", getPhoneNumberStruct.Watermark.Appid, o.AppId))
	}
	return &getPhoneNumberStruct, nil

}
func (o *_wx) unPad(s []byte) []byte {
	if len(s) == 0 {
		return []byte("{\"countryCode\":\"\",\"phoneNumber\":\"\",\"purePhoneNumber\":\"\",\"watermark\":{\"appid\":\"\",\"timestamp\":1621344576}}")
	}
	end := len(s) - int(s[len(s)-1])
	if end < 0 {
		return []byte("{\"countryCode\":\"\",\"phoneNumber\":\"\",\"purePhoneNumber\":\"\",\"watermark\":{\"appid\":\"\",\"timestamp\":1621344576}}")
	}
	return s[:end]
}
func (o *_wx) pkcs7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("wrong encryption parameters")
	} else {
		unPadding := int(origData[length-1])
		return origData[:(length - unPadding)], nil
	}
}
