package coupon

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/txbao/common-go/utils"
	"github.com/txbao/common-go/utils/logs"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//微信H5提交地址
	URL_H5_GATEWAY = "https://api.mch.weixin.qq.com/v3/pay/transactions/h5"
	//微信H5提交地址 -- 服务商版
	URL_H5_PARTNER_GATEWAY = "https://api.mch.weixin.qq.com/v3/pay/partner/transactions/h5"

	//微信JSAPI提交地址
	URL_JSAPI_GATEWAY = "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"
	//微信JSAPI提交地址 -- 服务商版
	URL_JSAPI_PARTNER_GATEWAY = "https://api.mch.weixin.qq.com/v3/pay/partner/transactions/jsapi"

	//微信扫码提交地址
	URL_NATIVE_GATEWAY = "https://api.mch.weixin.qq.com/v3/pay/transactions/native"

	//微信App提交地址
	URL_APP_GATEWAY = "https://api.mch.weixin.qq.com/v3/pay/transactions/app"

	//微信V3退款接口
	URL_V3_REFUND = "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds"
)

// GetPrivateKey 获取私钥
// filename 私钥的地址
func GetPrivateKey(filename string) (*rsa.PrivateKey, error) {
	keybuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	logs.Info("keybuffer", string(keybuffer))
	block, _ := pem.Decode([]byte(keybuffer))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	privatekey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privatekey.(*rsa.PrivateKey), nil
}

// 解析私钥
func ParsePKCS1PrivateKey(filename string) *rsa.PrivateKey {
	privateKey, err := ioutil.ReadFile(filename)
	if err != nil {
		logs.Info(err)
		return nil
	}
	block, _ := pem.Decode(privateKey)
	privateInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logs.Info(err)
		return nil
	}
	return privateInterface
}

//AES-256-GCM解密
func RsaDecrypt(ciphertext, nonce2, associatedData2 string, ApiV3Key string) (plaintext string, err error) {
	key := []byte(ApiV3Key) //key是APIv3密钥，长度32位，由管理员在商户平台上自行设置的
	additionalData := []byte(associatedData2)
	nonce := []byte(nonce2)

	block, err := aes.NewCipher(key)
	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	plaindata, err := aesgcm.Open(nil, nonce, cipherdata, additionalData)
	logs.Info("plaintext: ", string(plaindata))

	return string(plaindata), err
}

/**
RsaWithSHA256Base64加密
https://blog.csdn.net/weixin_42390256/article/details/94392258
*/
func RsaWithSHA256Base64(origData string, block []byte) (sign string, err error) {
	blocks, _ := pem.Decode(block)
	privateKey, _ := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	h := sha256.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)
	return
}

//V3请求
/**
stockCreatorMchid := "1578813281"

	offset := "0"
	limit := "10"
	urlApi := `https://api.mch.weixin.qq.com/v3/marketing/favor/stocks?stock_creator_mchid=`+ stockCreatorMchid +`&offset=` + offset + `&limit=` + limit

	res := weixin.V3Request(stockCreatorMchid,urlApi,"GET")
*/
func V3Request(mch_id string, urlApi string, http_method string, postData string) string {
	//记录开始时间
	start := time.Now() // 获取当前时间

	body := postData
	headerParam := CreateAuthorization(urlApi, mch_id, body, http_method)
	logs.Info("message:", headerParam)

	res, err := utils.HttpCURL(urlApi, http_method, postData, headerParam)
	if err != nil {
		fmt.Println("远程请求错误：URL："+urlApi, "参数："+postData, "错误："+utils.GetErrorMsg(err))
	}

	//记录结束时间
	elapsed := time.Since(start)
	logs.Info("微信WeixinReq请求数据：URL:", urlApi, "商户号："+mch_id+"请求数据："+postData, "\n", "响应数据：", res, "\n执行完成耗时：", elapsed)
	return res
}

//生成v3 Authorization
func CreateAuthorization(urlApi string, mchId string, body string, http_method string) url.Values {
	//设置证书
	//使用证书：cert 与 key 分别属于两个.pem文件
	//证书文件请放入服务器的非web目录下
	globalExcPath := utils.GetGlobalExcPath("./")
	sslCertPath := globalExcPath + `etc/cert/weixin/` + mchId + `/apiclient_cert.pem`
	mchPrivateKey := globalExcPath + `etc/cert/weixin/` + mchId + `/apiclient_key.pem`
	logs.Info(sslCertPath, mchPrivateKey)

	u, err := url.Parse(urlApi)
	if err != nil {
		panic(err)
	}

	canonicalUrl := u.Path + "?" + u.RawQuery
	if u.RawQuery == "" {
		canonicalUrl = u.Path
	}
	//logs.Info("u:",u.Path)
	//logs.Info("u2:",u.RawQuery)

	//当前时间戳
	timestamp := utils.DateUnix()
	//随机字符串
	nonce := utils.RandomString(32, 1)

	//POST请求时
	message := http_method + "\n"
	message += canonicalUrl + "\n"
	message += utils.IntToString(timestamp) + "\n"
	message += nonce + "\n"
	message += body + "\n"

	logs.Info("message===", "|"+message+"|\n")

	pemKey := utils.ReadFile(mchPrivateKey)
	sign, err := RsaWithSHA256Base64(message, []byte(pemKey))
	if err != nil {
		panic(err)
	}

	serialNo := GetSerialNumberHex(mchId)
	logs.Info("serialNo", serialNo, "\n")

	//Authorization 类型
	schema := "WECHATPAY2-SHA256-RSA2048"
	//生成token
	token := fmt.Sprintf(`mchid="%s",serial_no="%s",nonce_str="%s",timestamp="%d",signature="%s"`, mchId, serialNo, nonce, timestamp, sign)

	var param = url.Values{}
	param.Add("Content-Type", "application/json")
	param.Add("Accept", "application/json")
	param.Add("User-Agent", "*/*")
	param.Add("Authorization", schema+" "+token)
	logs.Info("schema", schema, "token2", token, "ALL==", schema+" "+token)

	return param

}

//生成CreatePaySign
func CreatePaySign(mchId string, appId string, timestamp int, nonceStr string, prepayId string) string {

	//设置证书
	//使用证书：cert 与 key 分别属于两个.pem文件
	//证书文件请放入服务器的非web目录下
	globalExcPath := utils.GetGlobalExcPath("./")
	sslCertPath := globalExcPath + `etc/cert/weixin/` + mchId + `/apiclient_cert.pem`
	mchPrivateKey := globalExcPath + `etc/cert/weixin/` + mchId + `/apiclient_key.pem`
	logs.Info(sslCertPath, mchPrivateKey)

	//POST请求时
	message := appId + "\n"
	message += utils.IntToString(timestamp) + "\n"
	message += nonceStr + "\n"
	message += prepayId + "\n"

	logs.Info("message===", "|"+message+"|\n")

	pemKey := utils.ReadFile(mchPrivateKey)
	sign, err := RsaWithSHA256Base64(message, []byte(pemKey))
	if err != nil {
		panic(err)
	}
	return sign
}

/**
获取证书序列号
*/
func GetSerialNumberHex(mch string) string {
	globalExcPath := utils.GetGlobalExcPath("./")
	filePath := globalExcPath + `etc/cert/weixin/` + mch + `/apiclient_cert.pem`

	_, x509Cert, _ := GetCertificate(filePath)
	logs.Info("XXXX:=========", x509Cert.SerialNumber, "==========")

	SerialNumber := utils.BigIntToHex(x509Cert.SerialNumber)
	return SerialNumber
}

/**
解析证书
*/
func GetCertificate(filePath string) (bool, *x509.Certificate, error) {
	certPEMBlock, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, nil, err
	}
	//获取证书信息 -----BEGIN CERTIFICATE-----   -----END CERTIFICATE-----
	//这里返回的第二个值是证书中剩余的 block, 一般是rsa私钥 也就是 -----BEGIN RSA PRIVATE KEY 部分
	//一般证书的有效期，组织信息等都在第一个部分里
	block, _ := pem.Decode([]byte(certPEMBlock))
	if block == nil {
		logs.Info("failed to parse certificate PEM")
		return false, nil, errors.New("failed to parse certificate PEM")
	}
	x509Cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		logs.Info("failed to parse certificate: " + err.Error())
		return false, nil, err
	}
	logs.Info("AAA", x509Cert.SerialNumber)
	return true, x509Cert, errors.New("ok")
}

func Decrypt(ciphertext string, rsaPrivateKey *rsa.PrivateKey) string {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	logs.Info("cipherdata", cipherdata)

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, rsaPrivateKey, cipherdata, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return ""
	}

	//fmt.Printf("Plaintext: %s\n", string(plaintext))
	return string(plaintext)
}

// 转换8进制utf-8字符串到中文
// eg: `\346\200\241` -> 怡
func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
