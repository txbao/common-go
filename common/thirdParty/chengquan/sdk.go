package chengquan

import (
	"github.com/txbao/common-go/common/utils"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ResCodePass = 7000 // 请求通过
)

type SDK struct {
	AppId     string
	Key       string
	ApiUrl    string
	NotifyUrl string
	AesKey    string
	AesIv     string
}

func (o *SDK) sign(params map[string]string) {
	params["app_id"] = o.AppId
	params["timestamp"] = strconv.FormatInt(time.Now().UnixNano(), 10)[:13]

	var keys, kvList []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if params[k] != "" {
			kvList = append(kvList, k+"="+params[k])
		}
	}
	kvList = append(kvList, "key="+o.Key)
	params["sign"] = fmt.Sprintf("%X", md5.Sum([]byte(strings.Join(kvList, "&"))))
}

func (o *SDK) post(api string, params map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	o.sign(params)
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	req, err := http.NewRequest("POST", o.ApiUrl+"/"+api, strings.NewReader(values.Encode()))
	if err != nil {
		log.Println("ChengQuanPostNewRequestErr", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ChengQuanPostClientDoErr", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ChengQuanPostReadAllErr", err, "body:", string(body))
		return nil, err
	}
	return body, nil
}

// 卡券品牌列表
func (o *SDK) CouponTypeList() (*CouponTypeListResponse, error) {
	api := "coupon/type/list"
	params := make(map[string]string)
	b, err := o.post(api, params)
	if err != nil {
		return nil, err
	}
	response := &CouponTypeListResponse{}
	err = json.Unmarshal(b, response)
	if err != nil {
		log.Println("ChengQuanCouponTypeListUnmarshalErr", err, string(b))
		return nil, err
	}
	return response, nil
}

// 卡券品牌产品列表
func (o *SDK) CouponTypeGoodsList(request CouponTypeGoodsListRequest) (*CouponTypeGoodsListResponse, error) {
	api := "coupon/type/goods/list"
	params := make(map[string]string)
	params["goods_type_id"] = strconv.FormatInt(request.GoodsTypeId, 10)
	b, err := o.post(api, params)
	if err != nil {
		return nil, err
	}
	response := &CouponTypeGoodsListResponse{}
	err = json.Unmarshal(b, response)
	if err != nil {
		log.Println("ChengQuanCouponTypeGoodsListUnmarshalErr", err, string(b))
		return nil, err
	}
	return response, nil
}

// 卡券产品库存
func (o *SDK) CouponGoodsStock(request CouponGoodsStockRequest) (*CouponGoodsStockResponse, error) {
	api := "coupon/goods/stock"
	params := make(map[string]string)
	params["goods_no"] = request.GoodsNo
	b, err := o.post(api, params)
	if err != nil {
		return nil, err
	}
	response := &CouponGoodsStockResponse{}
	err = json.Unmarshal(b, response)
	if err != nil {
		log.Println("ChengQuanCouponGoodsStockUnmarshalErr", err, string(b))
		return nil, err
	}
	return response, nil
}

// 卡券使用须知
func (o *SDK) CouponGoodsNotes(request CouponGoodsNotesRequest) (*CouponGoodsNotesResponse, error) {
	api := "coupon/goods/notes"
	params := make(map[string]string)
	params["goods_no"] = request.GoodsNo
	b, err := o.post(api, params)
	if err != nil {
		return nil, err
	}
	response := &CouponGoodsNotesResponse{}
	err = json.Unmarshal(b, response)
	if err != nil {
		log.Println("ChengQuanCouponGoodsNotesUnmarshalErr", err, string(b))
		return nil, err
	}
	return response, nil
}

// 卡券购买接口
func (o *SDK) CouponOrderPay(request CouponOrderPayRequest) (*CouponOrderPayResponse, error) {
	api := "coupon/order/pay"
	params := make(map[string]string)
	params["user_order_no"] = request.UserOrderNo        //商户提交的订单号
	params["goods_no"] = request.GoodsNo                 //产品编号
	params["count"] = utils.Int64ToString(request.Count) //购买数量(只能为正整数，最大等于50)
	if request.Mobile != "" {
		params["mobile"] = request.Mobile //接受卡密短信手机号
	}
	fmt.Println("参数：", utils.Map2Json(params))
	b, err := o.post(api, params)
	if err != nil {
		return nil, err
	}
	response := &CouponOrderPayResponse{}
	fmt.Println("b:", string(b))
	err = json.Unmarshal(b, response)
	if err != nil {
		logx.Error("ChengQuanCouponOrderPayErr", err, string(b), "订单号："+request.UserOrderNo)
		return nil, err
	}
	return response, nil
}

// 卡券订单状态查询接口
func (o *SDK) CouponOrderQuery(request CouponOrderQueryRequest) (*CouponOrderPayResponse, error) {
	api := "coupon/order/query"
	params := make(map[string]string)
	params["user_order_no"] = request.UserOrderNo //商户提交的订单号
	fmt.Println("参数：", utils.Map2Json(params))
	b, err := o.post(api, params)
	if err != nil {
		return nil, err
	}
	response := &CouponOrderPayResponse{}
	fmt.Println("b:", string(b))
	err = json.Unmarshal(b, response)
	if err != nil {
		logx.Error("ChengQuanCouponOrderQueryErr", err, string(b), "订单号："+request.UserOrderNo)
		return nil, err
	}
	return response, nil
}

// Aes解密
func (o *SDK) AesDecrypt(text string) string {
	if text == "" {
		return ""
	}
	AesCbcSdk := utils.NewAesCbcSdk(o.AesKey, o.AesIv)
	res, err := AesCbcSdk.AesCbcDecrypt(text)
	if err != nil {
		logx.Error("橙券解密错误：", err.Error())
		return ""
	}
	return res
}
