package cashout

import (
	"common-go/common/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type _sdk struct {
	xrsa *utils.XRsa
}

type ReqStruct struct {
	AppId   int64  `json:"app_id"`   //应用ID
	BizData string `json:"biz_data"` //接口参数（json格式）加密后字符串
	Sign    string `json:"sign"`     //签名
}

type ReqBizDataStruct struct {
	Amount       int64  `json:"amount"`         //提现金额（单位分）
	RequestNo    string `json:"request_no"`     //请求号，惟一
	Remarks      string `json:"remarks"`        //备注
	ThirdUserId  int64  `json:"third_user_id"`  //业务系统用户ID
	ThirdOrderNo string `json:"third_order_no"` //业务订单号
	Timestamp    int64  `json:"timestamp"`      //时间戳，如1569135887
	TradeType    int64  `json:"trade_type"`     //交易类型;1:消费、2:提现、3:返佣
	MergeData    string `json:"merge_data"`     //合并的数据
}

//通用格式化
func formatKey(key string) string {
	keyLen := len(key)
	keyStr := "\n"
	for i := 0; i <= keyLen; i = i + 64 {
		s := 0 + i
		e := 64 + i
		if e > keyLen {
			e = keyLen
		}
		keyStr += key[s:e] + "\n"
	}
	return keyStr
}

//格式化私钥
func FormatPrivateKey(privateKey string) string {
	return "-----BEGIN RSA PRIVATE KEY-----" + formatKey(privateKey) + "-----END RSA PRIVATE KEY-----"
}

// 格式化公钥
func FormatPublicKey(publicKey string) string {
	return "-----BEGIN PUBLIC KEY-----" + formatKey(publicKey) + "-----END PUBLIC KEY-----"
}

// map转 json str
func Map2Json(result interface{}) string {
	jsonBytes, _ := json.Marshal(result)
	jsonStr := string(jsonBytes)
	return jsonStr
}
func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func NewSdk(appId int64, privateKey string, publicKey string) (*_sdk, error) {
	priKey := bytes.NewBufferString(FormatPrivateKey(privateKey))
	pubKey := bytes.NewBufferString(FormatPublicKey(publicKey))

	var err error
	xrsa, err := utils.XrsaNewXRsa(pubKey.Bytes(), priKey.Bytes())
	if err != nil {
		return nil, err
	}
	return &_sdk{
		xrsa: xrsa,
	}, nil
}

//post请求
func (obj *_sdk) HttpPost(apiUrl string, params map[string]interface{}) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(params).
		SetHeader("Content-Type", "application/json").
		Post(apiUrl)
	if err != nil {
		return "", err
	}
	return resp.String(), err
}

//请求
func (obj *_sdk) Request(reqUrl string, appId int64, reqJson string) (string, error) {
	sign, _ := obj.xrsa.XrsaSign(reqJson)
	data, _ := obj.xrsa.XrsaPublicEncrypt(reqJson)

	//组织请求参数
	reqMp := make(map[string]interface{})
	reqMp["app_id"] = appId
	reqMp["biz_data"] = data
	reqMp["sign"] = sign

	res, err := obj.HttpPost(reqUrl, reqMp)
	fmt.Println("请求结果:", res, "请求地址：", reqUrl, "请求数据：", utils.Map2Json(reqMp), "加密数据：", reqJson)

	return res, err
}

//RSA验证
func (obj *_sdk) VerifyRsa(bizData string, sign string) (string, error) {
	decrypted, err := obj.xrsa.XrsaPrivateDecrypt(bizData)
	if err != nil {
		return "", err
	}

	err = obj.xrsa.XrsaVerify(decrypted, sign)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}
