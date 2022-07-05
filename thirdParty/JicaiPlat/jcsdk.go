package JicaiPlat

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/txbao/common-go/errorx"
	"github.com/txbao/common-go/utils"

	"io"
	"net/http"
	"sort"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	//成功code
	CodeSuccess = "00000"
)

//集采相关
type _jc struct {
	Host      string
	AppKey    string
	AppSecret string
}

func NewSdk(host string, appKey string, appSecret string) *_jc {
	return &_jc{
		Host:      host,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (l *_jc) getSignParm(m map[string]interface{}) (map[string]interface{}, string) {
	m["appKey"] = l.AppKey
	//m["timestamp"] = fmt.Sprintf("%v", utils.DateUnix64())
	sign := l.CreateSign(m, l.AppSecret)
	m["sign"] = sign
	queryStr := ""
	for k, v := range m {
		if queryStr != "" {
			queryStr += "&"
		}
		queryStr = fmt.Sprintf("%v%v=%v", queryStr, k, v)
	}
	return m, queryStr
}

//CreateSign 创建验签
func (l *_jc) CreateSign(params map[string]interface{}, appSecret string) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sign" {
			key = append(key, k)
		}
	}
	key = append(key, "app_secret")
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if key[i] == "app_secret" {
			str = str + appSecret
		} else {
			switch f := params[key[i]].(type) {
			case float64:
				str = str + utils.Float64ToStr(f)
			case map[string]interface{}:
				paramsJson, _ := json.Marshal(f)
				str = str + string(paramsJson)
			default:
				str = str + fmt.Sprintf("%v", f)
			}
		}
	}
	// 自定义签名算法
	sign := l.Md5String(str)
	return sign
}
func (l *_jc) Md5String(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

//post请求
func (l *_jc) httpPost(apiUrl string, params map[string]interface{}) (string, error) {
	//记录开始时间
	start := time.Now() // 获取当前时间
	client := resty.New()
	resp, err := client.R().
		SetBody(params).
		SetHeader("Content-Type", "application/json").
		Post(apiUrl)
	//记录结束时间
	elapsed := time.Since(start)
	logx.Info("集采平台请求数据：", apiUrl, "参数：", utils.Map2Json(params), "\n", "响应数据：", resp, "\n执行完成耗时：", elapsed)
	if err != nil {
		return "", err
	}
	return resp.String(), err
}
func (l *_jc) httpGet(url string) (response string, err error) {
	//记录开始时间
	start := time.Now() // 获取当前时间
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}

	response = result.String()
	//记录结束时间
	elapsed := time.Since(start)
	logx.Info("集采平台请求数据：", url, "\n", "响应数据：", response, "\n执行完成耗时：", elapsed)
	return
}

//OA客户表
func (l *_jc) OaCustomerList() (*OaCustomerList, error) {
	_, queryStr := l.getSignParm(map[string]interface{}{})
	url := l.Host + "/api/v1/third/oa-customer/list?" + queryStr
	resp, err := l.httpGet(url)
	if err != nil {
		return nil, err
	}
	var oaCustomerList OaCustomerList
	if err = json.Unmarshal([]byte(resp), &oaCustomerList); err != nil {
		return nil, err
	}
	return &oaCustomerList, nil
}

//制券回调
func (l *_jc) MakeCouponNotify(itemId int64, batchId int64, isVoucher int64) (*MakeCouponNotifyResp, error) {
	getSignParm, _ := l.getSignParm(map[string]interface{}{
		"itemId":    itemId,
		"serailNo":  batchId,
		"isVoucher": isVoucher,
	})
	url := l.Host + "/api/v1/open-erp/gb-store-purchase-order-child-item-voucher/voucher"
	res, err := l.httpPost(url, getSignParm)
	if err != nil {
		return nil, err
	}
	var makeCouponNotifyResp MakeCouponNotifyResp
	if err = json.Unmarshal([]byte(res), &makeCouponNotifyResp); err != nil {
		return nil, err
	}
	return &makeCouponNotifyResp, nil
}

func (l *_jc) GoodsList(customerId int64, spuName string, spuCode string, pageIndex int64, pageSize int64) (*IntegralResponse, error) {
	host := l.Host

	tmp := map[string]interface{}{
		"appKey":     l.AppKey,
		"customerId": customerId,
		"spuName":    spuName,
		"spuCode":    spuCode,
		"pageIndex":  pageIndex,
		"pageSize":   pageSize,
	}
	sign := utils.CreateSign(tmp, l.AppSecret)
	url := fmt.Sprintf("%s/api/v1/open-erp/gb-store-purchase-order-child-item/goods?appKey=%s&customerId=%d&spuName=%s&spuCode=%s&pageIndex=%d&pageSize=%d&sign=%s", host, l.AppKey, customerId, spuName, spuCode, pageIndex, pageSize, sign)
	resp, err := utils.HttpGet(url)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	fmt.Println("集采商品接口============", resp)

	res := &IntegralResponse{}
	_ = json.Unmarshal([]byte(resp), res)

	return res, err
}

func (l *_jc) GoodsStock(customerId int64, productId int64, pageIndex int64, pageSize int64) (*GoodsStock, error) {
	host := l.Host

	tmp := map[string]interface{}{
		"appKey":     l.AppKey,
		"customerId": customerId,
		"productId":  productId,
		"pageIndex":  pageIndex,
		"pageSize":   pageSize,
	}
	sign := utils.CreateSign(tmp, l.AppSecret)
	url := fmt.Sprintf("%s/api/v1/open-erp/gb-store-purchase-order-child-item/qty?appKey=%s&customerId=%d&productId=%d&pageIndex=%d&pageSize=%d&sign=%s", host, l.AppKey, customerId, productId, pageIndex, pageSize, sign)
	resp, err := utils.HttpGet(url)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	fmt.Println("集采商品库存接口============", resp)

	res := &GoodsStock{}
	json.Unmarshal([]byte(resp), res)
	return res, err
}

func (l *_jc) ReturnGoodsStock(itemId int64, isUse int64) error {
	getSignParm, _ := l.getSignParm(map[string]interface{}{
		"itemId": itemId,
		"isUse":  isUse,
	})
	url := l.Host + "/api/v1/open-erp/gb-store-purchase-order-child-item-voucher/qty"
	resp, err := l.httpPost(url, getSignParm)
	if err != nil {
		return errorx.NewDefaultError(err.Error())
	}
	fmt.Println("集采商品库存修改回调接口============", resp)
	res := &ReturnGoodsStock{}
	json.Unmarshal([]byte(resp), res)

	if res.Code != "00000" {
		return errorx.NewDefaultError(res.Msg)
	}
	return nil
}
