package smapiSdk

import (
	"bank-activity/common/utils"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ZZMarquis/gm/sm2"
	"github.com/go-resty/resty/v2"
)

//对外接口API-国密

type _sdk struct {
}

type ReqStruct struct {
	AppId int64  `json:"app_id"` //应用ID
	Data  string `json:"data"`   //接口参数（json格式）加密后字符串
	Sign  string `json:"sign"`   //签名
}

type ReqDataStruct struct {
	AppId     int64       `json:"app_id"`    //应用ID
	Service   string      `json:"service"`   //服务
	Method    string      `json:"method"`    //方法
	Timestamp int64       `json:"timestamp"` //时间戳，如1569135887
	Data      interface{} `json:"data"`
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

func NewSdk() (*_sdk, error) {
	//"priv_d": "e0e131531af3b9c002dc0a74bb0105923b2ce948aa21aae44fe4911522f35ffc",
	//	"pub_x": "84878cdebf26745071f444d3d38ba3ee5778afa847d324325031999b4549de97",
	//	"pub_y": "23e98128ca56ed739b0cd012d36c846028158f188c1d175120cf9f18e2894662"
	return &_sdk{
	}, nil
}

//post请求
func (o *_sdk) HttpPost(apiUrl string, params map[string]interface{}) (string, error) {
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
func (o *_sdk) Request(reqUrl string, appId int64, data string) (string, error) {
	//组织请求参数
	reqMp := make(map[string]interface{})
	reqMp["app_id"] = appId
	reqMp["data"] = data
	res, err := o.HttpPost(reqUrl, reqMp)
	fmt.Println("请求结果:", res, "请求地址：", reqUrl, "请求数据：", utils.Map2Json(reqMp))

	return res, err
}

func (o *_sdk) Encrypt(src string, pubX string, pubY string) (string, error) {
	pub1, _ := hex.DecodeString(fmt.Sprintf("%v%v", pubX, pubY))
	pub, err := sm2.RawBytesToPublicKey(pub1)
	if err != nil {
		fmt.Println("RawBytesToPublicKeyErr:", err.Error())
		return "", err
	}
	//pub,_ := sm2.RawBytesToPublicKey()
	cipherText, err := sm2.Encrypt(pub, []byte(src), sm2.C1C3C2)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	//fmt.Printf("cipher text:%s\n", hex.EncodeToString(cipherText))
	return hex.EncodeToString(cipherText), nil
}

//解密
func (o *_sdk) Decrypt(encyptStr string, privateKey string) (string, error) {
	//privateKey := "4f23ea50142147782ccae2189d01d4ddb2b98b794243541d89f17a9ec9a0659f"
	//PublicKey := "d48a271570bddcfbeff7313d86e9a39afb17427a1587b3dc42d65efb977d8c9d1a795b161d8f89f360a2425586fe88a7bad92d276c65ece7e9debfdeb9a2987c"
	//生成公钥私钥
	//priv, pub, err := sm2.GenerateKey(rand.Reader)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//导入公钥私钥
	priv1, err := hex.DecodeString(privateKey)
	if err != nil {
		fmt.Println("HexDecodeStringErr:", err.Error())
		return "", err
	}
	priv, err := sm2.RawBytesToPrivateKey(priv1)
	if err != nil {
		fmt.Println("RawBytesToPrivateKeyErr:", err.Error(), "priv1", priv1)
		return "", err
	}
	cipherText, _ := hex.DecodeString(encyptStr)
	plainText, err := sm2.Decrypt(priv, cipherText, sm2.C1C3C2)
	if err != nil {
		fmt.Println("sm2DecryptErr:", err.Error())
		return "", err
	}
	fmt.Printf("plainText:%s\n", plainText)
	return string(plainText), nil
}
