package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/txbao/common-go/utils"
	"github.com/txbao/common-go/utils/logs"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	MethodPay     = "pay"
	MethodCoupon  = "coupon"
	MethodRefund  = "refund"
	MethodSharing = "agentpay" //分账

	SharingPayType = "agentpay"

	PayTypeWeChatH5    = "weixin_h5"
	PayTypeWeChatJsApi = "weixin_jsapi"

	CouponTypeQueryInfo = "coupon_query_info"
	CouponTypeSend      = "coupon_send"
	CouponTypeLink      = "coupon_link"

	PayResultCodeSuccess    = "success"
	PayResultCodeFail       = "fail"
	PayResultCodeRefundPart = "refund_part"
	PayResultCodeRefund     = "refund"

	PayResultSuccess = "success"
	PayResultFail    = "fail"
	PayResultCancel  = "cancel"
)

//支付调用服务参数
type PayAppParams struct {
	AppKey          string //MD5 key
	PrivateKey      string //秘钥
	PublicKey       string //公钥
	AppId           string //appid
	Gateway         string //请求地址
	PlatformAccount string //平台账号
	XRsa            *utils.XRsa
}

//支付固定结构(业务参数)
type PaymentStruct struct {
	Method      string           `json:"method"`
	OutTradeNo  string           `json:"out_trade_no"`
	TotalAmount string           `json:"total_amount"`
	AppId       string           `json:"app_id"`
	ReturnUrl  string           `json:"return_url"`
	NotifyUrl  string           `json:"notify_url"`
	PayType    string           `json:"pay_type"`
	PayCode    string           `json:"pay_code"`
	SignType   string           `json:"sign_type"`
	Sign       string           `json:"sign"`
	Subject    string           `json:"subject"`
	Version    string           `json:"version"`
	Attach     string           `json:"attach"`
	Timestamp  string           `json:"timestamp"`
	BizContent BizContentStruct `json:"biz_content"`
}

type DetailStruct struct {
	GoodsDetail []GoodsDetailStruct `json:"goods_detail"`
}

type GoodsDetailStruct struct {
	GoodsId  string `json:"goods_id"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

//分账
type ReceiversStruct struct {
	Account     string `json:"account"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Rate        string `json:"rate"` //手续费承担比例
	Type        string `json:"type"`
}

//业务参数固定参数
type BizContentStruct struct {
	CloseTime string       `json:"close_time"`
	CouponId  string       `json:"coupon_id"`
	Detail    DetailStruct `json:"detail"`
	//Detail string `json:"detail"`
	Fee           string            `json:"fee"`
	FeeMode       string            `json:"fee_mode"` //分账手续费承担模式，1按收款金额承担手续费，2按比例承担手续费，3平台承担手续费
	GoodsTag      string            `json:"goods_tag"`
	Ip            string            `json:"ip"`
	LimitPayType  string            `json:"limit_pay_type"` //限制（即排除）支付类型。适用于本地生活，格式为”101-102或weixin_h5-alipay_h5”，支付应用分配的支付类开>=限制支付类型
	Mobile        string            `json:"mobile"`
	Openid        string            `json:"openid"`
	Receivers     []ReceiversStruct `json:"receivers"`
	RefundFee     string            `json:"refund_fee"` //退款金额
	Sharing       string            `json:"sharing"`
	SharingPeriod string            `json:"sharing_period"` //1为分账周期T+1，N为T+N
	StockId       string            `json:"stock_id"`
}

//数据结构
type DataStruct struct {
	AppId    string `json:"app_id"`
	SignType string `json:"sign_type"`
	Data     string `json:"data"`
}

//返回结构
type ResStruct struct {
	Gateway string      `json:"gateway"`
	Url     string      `json:"url"`
	Data    interface{} `json:"data"`
}

//回调结构体
type CallbackStruct struct {
	AgentpayTradeNo string `json:"agentpay_trade_no"`
	AppId           string `json:"app_id"`
	Attach          string `json:"attach"`
	BankType        string `json:"bank_type"`
	CashFee         string `json:"cash_fee"`
	ErrorCode       string `json:"error_code"`
	GoodsDetail     string `json:"goods_detail"`
	Openid          string `json:"openid"`
	OutTradeNo      string `json:"out_trade_no"`
	PayType         string `json:"pay_type"`
	ResultCode      string `json:"result_code"`
	SignType        string `json:"sign_type"`
	Subject         string `json:"subject"`
	SysTradeNo      string `json:"sys_trade_no"`
	ThirdTradeNo    string `json:"third_trade_no"`
	TotalAmount     string `json:"total_amount"`
	OriginalAmount  string `json:"original_amount"`
	Discount        string `json:"discount"`
	Sign            string `json:"sign"`
}

//退款 {"success":true,"statusCode":200,"statusText":"OK","return_code":"1","return_msg":"退款成功"}
//{"return_code":0,"return_msg":"PAY:【202006161302339553356290】此订单已完全退款！","code":190052}
type RefundResponse struct {
	Success    bool           `json:"success"`
	StatusCode int            `json:"statusCode"`
	StatusText string         `json:"statusText"`
	Code       string         `json:"code"`
	Msg        string         `json:"msg"`
	ReturnMsg  string         `json:"return_msg"`
	Timestamp  int64          `json:"timestamp"`
	Data       SendCouponData `json:"data"`
}
type RefundData struct {
	OutTradeNo    string `json:"out_trade_no"`
	RefundFee     string `json:"refund_fee"`
	RefundTradeNo string `json:"refund_trade_no"`
	SysTradeNo    string `json:"sys_trade_no"`
}

// 发券
type SendCouponResponse struct {
	Success    bool           `json:"success"`
	StatusCode int            `json:"statusCode"`
	StatusText string         `json:"statusText"`
	Code       string         `json:"code"`
	Msg        string         `json:"msg"`
	ReturnMsg  string         `json:"return_msg"`
	Timestamp  int64          `json:"timestamp"`
	Data       SendCouponData `json:"data"`
}
type SendCouponData struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	CouponId string `json:"coupon_id"`
}

// 查券
//type CouponInfoResponse struct {
//	Success    bool   `json:"success"`
//	StatusCode int    `json:"statusCode"`
//	StatusText string `json:"statusText"`
//	ReturnCode int    `json:"return_code"`
//	ReturnMsg  string `json:"return_msg"`
//	Data       string `json:"data"`
//}
//type CouponInfoData struct {
//	AvailableBeginTime      time.Time            `json:"code"`                      // 可用开始时间
//	AvailableEndTime        time.Time            `json:"message"`                   // 可用结束时间
//	CouponId                string               `json:"coupon_id"`                 // 代金券id
//	CouponName              string               `json:"coupon_name"`               // 代金券名称
//	CouponType              string               `json:"coupon_type"`               // 券类型 NORMAL：满减券 CUT_TO：减至券
//	CreateTime              time.Time            `json:"create_time"`               // 创建时间
//	Description             string               `json:"description"`               // 使用说明
//	NoCash                  bool                 `json:"no_cash"`                   // 是否无资金流
//	NormalCouponInformation NormalCouponInfoData `json:"normal_coupon_information"` // 满减券信息
//	SingleItem              bool                 `json:"singleitem"`                // 是否单品优惠
//	Status                  string               `json:"status"`                    // 代金券状态 SENDED：可用 USED：已实扣 EXPIRED：已过期
//	StockCreatorMchId       string               `json:"stock_creator_mchid"`       // 创建批次的商户号
//	StockId                 string               `json:"stock_id"`                  // 批次号
//}
//type NormalCouponInfoData struct {
//	CouponAmount       uint64 `json:"coupon_amount"`       // 面额
//	TransactionMinimum uint64 `json:"transaction_minimum"` // 门槛
//}

const (
	SIGN_TYPE = "rsa"
	VERSION   = "1.0"
)

var (
//gateway string
//appId   string
//scheme  string
)

//func init() {
//	privateKey := bytes.NewBufferString(utils.FormatPrivateKey(global.Config.Payment.PrivateKey))
//	publicKey := bytes.NewBufferString(utils.FormatPublicKey(global.Config.Payment.PublicKey))
//
//	//fmt.Println("公钥：",global.Config.Payment.PrivateKey,"<")
//	//fmt.Println("公钥：",global.Config.Payment.PublicKey,"<")
//
//	gateway = global.Config.Payment.Gateway
//	appId = global.Config.Payment.AppId
//	scheme = global.Config.Application.Scheme
//
//	var err error
//	xrsa, err = utils.XrsaNewXRsa(publicKey.Bytes(), privateKey.Bytes())
//	if err != nil {
//		fmt.Println("err", err)
//		return
//	}
//}

//创建实例结构体
func NewPaymentStruct(p *PayAppParams) (*PayAppParams, error) {
	var err error
	privateKey := bytes.NewBufferString(utils.FormatPrivateKey(p.PrivateKey))
	publicKey := bytes.NewBufferString(utils.FormatPublicKey(p.PublicKey))

	p.XRsa, err = utils.XrsaNewXRsa(publicKey.Bytes(), privateKey.Bytes())
	if err != nil {
		fmt.Println("payment 初始化 err::", err)
		return p, err
	}
	return p, nil
}

//支付
func (p *PayAppParams) Pay(model *PaymentStruct) (*ResStruct, error) {
	if p.Gateway == "" || p.AppId == "" || p.PublicKey == "" {
		return &ResStruct{}, fmt.Errorf("支付必要参数不能为空::%+v", p)
	}

	if model.AppId == "" {
		model.AppId = p.AppId
	}
	if model.ReturnUrl == "" {
		model.ReturnUrl = p.Gateway + "/v1/pay/return"
	}
	if model.NotifyUrl == "" {
		model.NotifyUrl = p.Gateway + "/v1/pay/notify"
	}
	if model.SignType == "" {
		model.SignType = SIGN_TYPE
	}
	if model.Version == "" {
		model.Version = VERSION
	}
	if model.Timestamp == "" {
		model.Timestamp = utils.Int64ToString(time.Now().Unix())
	}
	model.TotalAmount = fmt.Sprintf("%.2f", utils.StringToFloat64(model.TotalAmount))

	signStr := p.SortParam(p.PaymentStructToUrl(model))
	if model.SignType == "md5" {
		return p.Md5Payment(signStr, model), nil
	}
	return p.RsaPayment(signStr, model), nil
}

//struct 转 url.Values
func (p *PayAppParams) PaymentStructToUrl(payStruct *PaymentStruct) url.Values {
	bizContentJsonBytes, err := json.Marshal(payStruct.BizContent)
	fmt.Println("payStruct.BizContent", payStruct.BizContent)
	if err != nil {
		fmt.Println(err)
	}
	bizContentJson := string(bizContentJsonBytes)

	fmt.Println("bizContentJson", bizContentJson)

	var param = url.Values{}
	param.Add("method", payStruct.Method)
	param.Add("out_trade_no", payStruct.OutTradeNo)
	param.Add("total_amount", payStruct.TotalAmount)
	param.Add("app_id", payStruct.AppId)
	param.Add("return_url", payStruct.ReturnUrl)
	param.Add("notify_url", payStruct.NotifyUrl)
	param.Add("pay_type", payStruct.PayType)
	param.Add("pay_code", payStruct.PayCode)
	param.Add("sign_type", payStruct.SignType)
	param.Add("subject", payStruct.Subject)
	param.Add("version", payStruct.Version)
	param.Add("attach", payStruct.Attach)
	param.Add("timestamp", payStruct.Timestamp)
	param.Add("biz_content", bizContentJson)
	return param
}

//排序
func (p *PayAppParams) SortParam(param url.Values) string {
	if param == nil {
		param = make(url.Values, 0)
	}
	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			if SIGN_TYPE == "md5" && key == "biz_content" {
				//value = utils.URLEncode(value)
			}
			pList = append(pList, key+"="+StrReplaceEncrypt(value))
		}
	}
	sort.Strings(pList)
	var src = strings.Join(pList, "&")
	return src
}

//Md5加密
func (p *PayAppParams) Md5Payment(signStr string, model *PaymentStruct) *ResStruct {
	fmt.Println("signStr", signStr)
	sign := utils.Md5(signStr + "&key=" + p.AppKey)
	signStr += "&sign=" + sign
	model.Sign = sign

	resModel := &ResStruct{
		Gateway: p.Gateway,
		Url:     p.Gateway + "?" + signStr,
		Data:    model,
	}
	return resModel
}

//RSA加密
func (p *PayAppParams) RsaPayment(signStr string, model *PaymentStruct) *ResStruct {
	fmt.Println("signStr", signStr)
	fmt.Println("")
	sign, _ := p.XRsa.XrsaSign(signStr)
	signStr += "&sign=" + sign

	model.Sign = sign

	modelJson, _ := utils.StructToJsonStr(model)
	fmt.Println("modelJson:", modelJson)

	data, _ := p.XRsa.XrsaPublicEncrypt(modelJson)
	dataModel := &DataStruct{
		AppId:    p.AppId,
		SignType: SIGN_TYPE,
		Data:     data,
	}

	resModel := &ResStruct{
		Gateway: p.Gateway,
		Url:     p.Gateway + "?app_id=" + p.AppId + "&sign_type=" + SIGN_TYPE + "&data=" + data,
		Data:    dataModel,
	}
	return resModel
}

//远程请求
func (p *PayAppParams) GetRequest(model *PaymentStruct) string {
	resModel, err := p.Pay(model)
	if err != nil {
		logs.Error("err::", err)
	}
	dataModel := (resModel.Data).(*DataStruct)
	data := "app_id=" + dataModel.AppId + "&sign_type=" + dataModel.SignType + "&data=" + dataModel.Data
	res, _ := utils.HttpPost(resModel.Gateway, data)
	return res
}

//是否是微信浏览器访问
func IsWeiXinBrowser(ctx *gin.Context) bool {
	userAgent := ctx.GetHeader("User-Agent")
	match, _ := regexp.MatchString("MicroMessenger", userAgent)
	if match {
		return true
	} else {
		return false
	}
}

//获取支付数据
//func GetPayData(model *PaymentStruct) *ResStruct {
//	//host := global.Config.Application.Scheme + ctx.Request.Host
//	host := global.Config.Application.ApiUrl
//	if model.AppId == "" {
//		model.AppId = model.AppParams.AppId
//	}
//	if model.ReturnUrl == "" {
//		model.ReturnUrl = host + "/v1/pay/return"
//	}
//	if model.NotifyUrl == "" {
//		model.NotifyUrl = host + "/v1/pay/notify"
//	}
//	if model.SignType == "" {
//		model.SignType = SIGN_TYPE
//	}
//	if model.Version == "" {
//		model.Version = VERSION
//	}
//	if model.Timestamp == "" {
//		model.Timestamp = utils.Int64ToString(time.Now().Unix())
//	}
//	model.TotalAmount = fmt.Sprintf("%.2f", utils.StringToFloat64(model.TotalAmount))
//
//	signStr := sortParam(paymentStructToUrl(model))
//
//	if model.SignType == "md5" {
//		return md5Payment(signStr, model)
//	}
//	return rsaPayment(signStr, model)
//}

//RSA验证
func (p *PayAppParams) VerifyRsa(data string) (*CallbackStruct, error) {
	decrypted, err := p.XRsa.XrsaPrivateDecrypt(data)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse("http://www.sqqmall.com/a.php?" + decrypted)
	if err != nil {
		//panic(err)
		return nil, err
	}
	m1 := u.Query()

	var param = url.Values{}
	param.Add("agentpay_trade_no", m1.Get("agentpay_trade_no"))
	param.Add("app_id", m1.Get("app_id"))
	param.Add("attach", m1.Get("attach"))
	param.Add("bank_type", m1.Get("bank_type"))
	param.Add("cash_fee", m1.Get("cash_fee"))
	param.Add("discount", m1.Get("discount"))
	param.Add("error_code", m1.Get("error_code"))
	param.Add("goods_detail", m1.Get("goods_detail"))
	param.Add("openid", m1.Get("openid"))
	param.Add("original_amount", m1.Get("original_amount"))
	param.Add("out_trade_no", m1.Get("out_trade_no"))
	param.Add("pay_type", m1.Get("pay_type"))
	param.Add("result_code", m1.Get("result_code"))
	param.Add("sign_type", m1.Get("sign_type"))
	param.Add("subject", m1.Get("subject"))
	param.Add("sys_trade_no", m1.Get("sys_trade_no"))
	param.Add("third_trade_no", m1.Get("third_trade_no"))
	param.Add("timestamp", m1.Get("timestamp"))
	param.Add("total_amount", m1.Get("total_amount"))
	param.Add("msg", m1.Get("msg"))

	sign := m1.Get("sign")

	var callbackModel CallbackStruct
	callbackModel = CallbackStruct{
		AgentpayTradeNo: m1.Get("agentpay_trade_no"),
		AppId:           m1.Get("app_id"),
		Attach:          m1.Get("attach"),
		BankType:        m1.Get("bank_type"),
		ErrorCode:       m1.Get("error_code"),
		GoodsDetail:     m1.Get("goods_detail"),
		Openid:          m1.Get("openid"),
		OutTradeNo:      m1.Get("out_trade_no"),
		PayType:         m1.Get("pay_type"),
		ResultCode:      m1.Get("result_code"),
		SignType:        m1.Get("sign_type"),
		Subject:         m1.Get("subject"),
		SysTradeNo:      m1.Get("sys_trade_no"),
		ThirdTradeNo:    m1.Get("third_trade_no"),
		TotalAmount:     m1.Get("total_amount"),
		OriginalAmount:  m1.Get("original_amount"),
		Discount:        m1.Get("discount"),
		CashFee:         m1.Get("cash_fee"),
		Sign:            sign,
	}

	signStr := p.SortParam(param)
	logs.Info("VerifyRsaSignStr", signStr)
	err = p.XRsa.XrsaVerify(signStr, sign)
	if err != nil {
		logs.Error("VerifyRsaFail", "")
		return &callbackModel, err
	}
	return &callbackModel, nil
}

//获取支付数据
func SendData(ctx *gin.Context) {
	/*
		dataModel := payment(ctx)
		html := `<html><head><meta http-equiv="Content-Type" content="text/html;charset=UTF-8"></head><body><form action="` + gateway + `" id="paymentform" method="post" >`
		html += `<input type='hidden' name='app_id' value='` + appId + `' />`
		html += `<input type='hidden' name='sign_type' value='` + SIGN_TYPE + `' />`
		html += `<input type='hidden' name='data' value='` + dataModel.Data + `' />`
		//        html += `<input type="submit" valuse="提交" />`
		html += `</form><script>document.getElementById("paymentform").submit()</script></body></html>`
		fmt.Println(html)
	*/
}

//callback struct 转 url.Values
func callbackStructToUrl(callbackStruct *CallbackStruct) url.Values {
	var param = url.Values{}
	param.Add("app_id", callbackStruct.AppId)
	param.Add("attach", callbackStruct.Attach)
	param.Add("bank_type", callbackStruct.BankType)
	param.Add("error_code", callbackStruct.ErrorCode)
	param.Add("openid", callbackStruct.Openid)
	param.Add("out_trade_no", callbackStruct.OutTradeNo)
	param.Add("pay_type", callbackStruct.PayType)
	param.Add("result_code", callbackStruct.ResultCode)
	param.Add("sign_type", callbackStruct.SignType)
	param.Add("subject", callbackStruct.Subject)
	param.Add("sys_trade_no", callbackStruct.SysTradeNo)
	param.Add("third_trade_no", callbackStruct.ThirdTradeNo)
	param.Add("total_amount", callbackStruct.TotalAmount)
	return param
}

//附加参数分割
func AttachSlipt(attach string) (string, string) {
	activityId := "0"
	tel := "1"
	attachArr := strings.Split(attach, "_")
	if len(attachArr) > 0 {
		activityId = attachArr[0]
	}
	if len(attachArr) > 1 {
		tel = attachArr[1]
	}
	return activityId, tel
}
func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

//加密的特殊字符替换， 如"+",空格等
func StrReplaceEncrypt(subject string) string {
	subject = StrReplace(" ", "", subject, -1)
	subject = StrReplace("+", "＋", subject, -1)
	return subject
}

// 退款
//func Refund(outTradeNo string) (*RefundResponse, error) {
//	res := GetRequest(&PaymentStruct{
//		Method:     MethodRefund,
//		OutTradeNo: outTradeNo,
//		PayType:    "refund",
//		NotifyUrl:  global.Config.Application.ApiUrl + "/v1/pay/discard",
//		BizContent: BizContentStruct{
//			Ip: "0.0.0.0",
//		},
//	})
//	logs.Info("RefundRes", res)
//	refundRes := &RefundResponse{}
//	err := json.Unmarshal([]byte(res), refundRes)
//	if err != nil {
//		logs.Error("RefundError", err)
//	}
//	return refundRes, err
//}
