package lianlianlvyou

//联联周边游

import (
	"common-go/utils"
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	CodeOk = 200
)

type LianLian struct {
	ChannelId int
	Key       string
	Url       string
	secretKey []byte
}

type lianLianRequest struct {
	Sign          string `json:"sign"`
	EncryptedData string `json:"encryptedData"`
	Timestamp     string `json:"timestamp"`
	ChannelId     string `json:"channelId"`
}

type lianLianData struct {
	Sign          string `json:"sign"`
	EncryptedData string `json:"encryptedData"`
	Timestamp     int64  `json:"timestamp"`
	ChannelId     int    `json:"channelId"`
}

type lianLianResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    lianLianData `json:"data"`
}

type lianLianRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (o *LianLian) getSecretKey() []byte {
	if o.secretKey == nil {
		var err error
		o.secretKey, err = base64.StdEncoding.DecodeString(o.Key)
		if err != nil {
			log.Println("LianLianDecodeKeyErr:", err)
		}
	}
	return o.secretKey
}

func (o *LianLian) getEncryptedData(data []byte) string {
	return base64.StdEncoding.EncodeToString(o.aesEcbEncrypt(data))
}

func (o *LianLian) aesEcbEncrypt(src []byte) []byte {
	block, err := aes.NewCipher(o.getSecretKey())
	if err != nil {
		log.Println("LianLianAesEcbEncryptNewCipherErr", err)
		return nil
	}
	blockSize := block.BlockSize()
	src = o.pkcs5Padding(src, blockSize)
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}

	return out
}

func (o *LianLian) aesEcbDecrypt(src []byte) ([]byte, error) {
	block, err := aes.NewCipher(o.getSecretKey())
	if err != nil {
		log.Println("LianLianAesEcbDecryptNewCipherErr", err)
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(src)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Decrypt(dst, src[:blockSize])
		src = src[blockSize:]
		dst = dst[blockSize:]
	}

	return o.pkcs5UnPadding(out), nil
}

func (o *LianLian) pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func (o *LianLian) pkcs5UnPadding(cipherText []byte) []byte {
	length := len(cipherText)
	unPadding := int(cipherText[length-1])
	return cipherText[:(length - unPadding)]
}

func (o *LianLian) formatParams(params []byte) string {
	data := lianLianRequest{
		EncryptedData: o.getEncryptedData(params),
		ChannelId:     strconv.Itoa(o.ChannelId),
		Timestamp:     strconv.FormatInt(time.Now().Unix(), 10),
	}
	h := md5.New()
	io.WriteString(h, data.EncryptedData)
	io.WriteString(h, data.ChannelId)
	io.WriteString(h, data.Timestamp)
	data.Sign = fmt.Sprintf("%x", h.Sum(nil))
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("LianLianFormatParamsErr", err)
	}
	return string(b)
}

func (o *LianLian) post(api string, params map[string]interface{}) (*lianLianRes, error) {
	paramsByte := []byte(utils.Map2Json(params))
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := o.Url + api
	req, err := http.NewRequest("POST", url, strings.NewReader(o.formatParams(paramsByte)))
	if err != nil {
		log.Println("LianLianPostNewRequestErr", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("LianLianPostClientDoErr", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("LianLianPostReadAllErr", err, "body:", string(body))
		return nil, err
	}

	response := &lianLianResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		log.Println("LianLianPostUnmarshalErr", err, "body:", string(body))
		return nil, err
	}

	res := &lianLianRes{
		Code:    response.Code,
		Message: response.Message,
	}
	if response.Code != CodeOk {
		log.Println("LianLianResCodeNotOk", response.Code, response.Message)
		return res, nil
	}

	// 验证签名
	h := md5.New()
	io.WriteString(h, response.Data.EncryptedData)
	io.WriteString(h, strconv.FormatInt(response.Data.Timestamp, 10))
	sign := fmt.Sprintf("%x", h.Sum(nil))
	if sign != response.Data.Sign {
		log.Println("LianLianCheckSignFailed:", sign, response.Data.Sign, api)
		res.Code = 40001
		res.Message = "返回数据验签失败"
		return res, nil
	}

	data, err := base64.StdEncoding.DecodeString(response.Data.EncryptedData)
	if err != nil {
		log.Println("LianLianPostDecodeStringErr:", err)
		res.Code = 40002
		res.Message = "返回数据Base64Decode失败"
		return res, nil
	}

	res.Data, err = o.aesEcbDecrypt(data)
	if err != nil {
		log.Println("LianLianAesEcbDecryptErr:", err)
	}
	return res, err
}

//解密
func (o *LianLian) AesEcbDecrypt(src string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return o.aesEcbDecrypt(data)
}

// 查询产品列表
func (o *LianLian) ProductList(params map[string]interface{}) (*ProductListRes, error) {
	res, err := o.post("/ll/api/product/queryProductList", params)
	if err != nil {
		return nil, err
	} else if res == nil {
		return nil, errors.New("res is nil")
	} else if res.Code != CodeOk {
		return nil, errors.New(res.Message)
	}

	data := &ProductListRes{}
	err = json.Unmarshal(res.Data.([]byte), data)
	if err != nil {
		log.Println("LianLianProductListUnmarshalErr:", string(res.Data.([]byte)), err)
	}
	return data, err
}

// 3、创建渠道订单-并自动发码
func (o *LianLian) OrderCreate(params map[string]interface{}) (*CreateOrderRes, error) {
	fmt.Println("联联采购：创建渠道订单-并自动发码")
	//params := utils.Struct2Map(req)
	fmt.Println("OrderCreate-params", utils.Map2Json(params))
	res, err := o.post("/ll/api/channelOrder/createOrder", params)
	if err != nil {
		return nil, err
	} else if res == nil {
		return nil, errors.New("res is nil")
	} else if res.Code != CodeOk {
		return nil, errors.New(res.Message)
	}

	fmt.Println("orderNo_"+utils.MapValue(params, "thirdPartyOrderNo")+":联联下单结果：", string(res.Data.([]byte)))
	data := &CreateOrderRes{}
	err = json.Unmarshal(res.Data.([]byte), data)
	if err != nil {
		log.Println("LianLianOrderCreateUnmarshalErr:", string(res.Data.([]byte)), err)
	}
	return data, err
}

// 6、查询产品信息
func (o *LianLian) ProductByCondition(params map[string]interface{}) (*ProductInfo, error) {
	res, err := o.post("/ll/api/product/queryProductByCondition", params)
	if err != nil {
		return nil, err
	} else if res == nil {
		return nil, errors.New("res is nil")
	} else if res.Code != CodeOk {
		return nil, errors.New(res.Message)
	}

	data := &ProductInfo{}
	err = json.Unmarshal(res.Data.([]byte), data)
	if err != nil {
		log.Println("LianLianProductByConditionUnmarshalErr:", string(res.Data.([]byte)), err)
	}
	return data, err
}

// 查询产品图文详情(文案)
func (o *LianLian) ProductDetailHtml(params map[string]interface{}) (*ProductDetail, error) {
	res, err := o.post("/ll/api/product/detail/html", params)
	if err != nil {
		return nil, err
	} else if res == nil {
		return nil, errors.New("res is nil")
	} else if res.Code != CodeOk {
		return nil, errors.New(res.Message)
	}

	data := &ProductDetail{}
	err = json.Unmarshal(res.Data.([]byte), data)
	if err != nil {
		log.Println("LianLianProductDetailHtmlUnmarshalErr:", string(res.Data.([]byte)), err)
	}
	return data, err
}

// 查询站点列表
func (o *LianLian) LocationList(params map[string]interface{}) ([]*LocationInfo, error) {
	res, err := o.post("/ll/api/location/getLocationList", params)
	if err != nil {
		return nil, err
	} else if res == nil {
		return nil, errors.New("res is nil")
	} else if res.Code != CodeOk {
		return nil, errors.New(res.Message)
	}

	var data []*LocationInfo
	err = json.Unmarshal(res.Data.([]byte), &data)
	if err != nil {
		log.Println("LianLianLocationListUnmarshalErr:", string(res.Data.([]byte)), err)
	}
	return data, err
}

// 查询站点列表
func (o *LianLian) ProductCategory(params map[string]interface{}) ([]ProductCategory, error) {
	res, err := o.post("/ll/api/product/queryProductCategory", params)
	if err != nil {
		return nil, err
	} else if res == nil {
		return nil, errors.New("res is nil")
	} else if res.Code != CodeOk {
		return nil, errors.New(res.Message)
	}

	var data []ProductCategory
	err = json.Unmarshal(res.Data.([]byte), &data)
	if err != nil {
		log.Println("LianLianProductCategoryUnmarshalErr:", string(res.Data.([]byte)), err)
	}
	return data, err
}
