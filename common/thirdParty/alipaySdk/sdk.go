package alipaySdk

import (
	"common-go/common/thirdParty/alipaySdk/alipay"
	"common-go/common/utils"
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	//"github.com/go-pay/gopay"
	//"github.com/go-pay/gopay/alipay"
	//"github.com/go-pay/gopay/pkg/xlog"
)

const (
	//AppId      = "2019091167209566"
	//PrivateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCkuYetJ/EnzuqS9RhumWk5356X0AARFZAMJ0eha2AeflbPRot94P8aV0ZkuV4mqHwWDY2Nxoil0MZu/4sv6RgwmcQMab0MMjCKdMEAb2/fFwR6iiu8g8/Ct+4UK6Xiav48hBMe3rjMChiRuwzR4xa78x46x5n4DaLmJDHgAcUUWm9q9KmoL/Vjp3IUc1y2E2AofhxtH3qxCozX8XeM2dLDpyUKKUXPLgWAv6x8RUv0IWFqo6vWjhGkgYPBjMAEilcN+wzvgMZMRMln46j8Xz/nYtuejC9yvAYcwhjrBn01OECMmXO45o8E+SSvmJlmN5ifKxLyMTq9ZM73kUwoGssbAgMBAAECggEAN1TSzElEajjI9sA5ir0haX1CCoCl5Rc3Ib34A7LLwLJzeoCZzpjLWA/E54SUqauiss2upNbxg7FTVmmkWV2U8I5WGk1SYUxinb+GLR1BmNBkgrzy4AnuuduKr1/SAvIpoC5FNfNeY7tocVtBfScotc1+dAQkJfx+oYu67SHrD2CfWPLCcvYM8a7nD+Muh2nJHNtG9ldpXXksSd6Aa5ueMQsA6z/ZimKEavzojuPiQyp9CVgSCGhbId54FMOfat1jkIZYHhAcOlmheMaGO3YICBVHCH/yQHwVxk9alRVfraOjVkePm4B3rXyO+y1z/puMm6ofJ9pXXJoCouyCXdmDAQKBgQDSU6xYfCvxu4jkZLDXUMe4daQhA+pA6t3PAzPuFiigGos8BTLvSG3qW7LVU5t9ib/NwSeckSKI7Gg/1N5YWImSwPer9eGUdFI7721zRMmov/yufi2PNxoH6H6Q8TcmL4JT9FvSPNkS6zrYNX8bU67rpvC2oQnzGgUwNARv44zkQQKBgQDIfsdtvGUj9w78OOi0rBDUZJWDScg4HM3mAjNPzlPdvPsMFuKbKGRv7PW6NH8XYv/qFyf3W0E1wWgQUvkbeSbbuPpBFW+JK/TtT86UBzdm2Ioeu30VxgN2xMaCaX64FtAryZVoU9a+wiDEccSoPQIXVlgk/JftVwZpYEWnJHuoWwKBgGzLc17h/CkyhID/xKnb0zOLRrb9O3MRCmNGmNoTBDitAlCtr8cuhAyyGjNW6Btr4ZcgzsiiGmcRQmuiRNEWGzOaNYLx0RnlgZQXKVpyvo1XofDwu8hVDFYC4VVAMPLDqHqc79I0P2UV4H5uTa1pABZNaD5P1sG1N/HTNmhaIEsBAoGAJMv/1ggJl/wz596Z9u7Nd+2t8xaLhKDkuR4WRMNdaQjSGnnoxQk4xcH7p10TJjupNFFjqGY7JZ6FdtmDlqM5moHsjB1fBxHmNde71jcs19dOIi3O2zwoTpf2xuCNvUOSmTa6EReyBfijAurOaQIpbBcQvlkNmSWrq5NbtT1g5JkCgYAsOdvs2U8P4zIQ8ZBqLKXrYktqTbX/QOcq8WsrBEhda0MtcYExchZ4JZC/NwRfwyXVDluGeVMHPkX5BmpBAobKRjjKjUa18fi2EHRwewX2H64crmMM5ow8BlZxO0YlzKElo5fEcF0InBY3hL21rDzFoCLCgAjN6DEB4zsNw+FLLg=="
	//PublibKey  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApLmHrSfxJ87qkvUYbplpOd+el9AAERWQDCdHoWtgHn5Wz0aLfeD/GldGZLleJqh8Fg2NjcaIpdDGbv+LL+kYMJnEDGm9DDIwinTBAG9v3xcEeoorvIPPwrfuFCul4mr+PIQTHt64zAoYkbsM0eMWu/MeOseZ+A2i5iQx4AHFFFpvavSpqC/1Y6dyFHNcthNgKH4cbR96sQqM1/F3jNnSw6clCilFzy4FgL+sfEVL9CFhaqOr1o4RpIGDwYzABIpXDfsM74DGTETJZ+Oo/F8/52LbnowvcrwGHMIY6wZ9NThAjJlzuOaPBPkkr5iZZjeYnysS8jE6vWTO95FMKBrLGwIDAQAB"

	StatusSuccess = "SUCCESS" //成功
	StatusFail    = "FAIL"    //失败（具体失败原因请参见error_code以及fail_reason返回值）
	StatusDealing = "DEALING" //处理中
	StatusRefund  = "REFUND"  //退票
)

type AliPaySdk struct {
	AppId      string         `json:"app_id"`
	PrivateKey string         `json:"private_key"`
	PublicKey  string         `json:"public_key"`
	Client     *alipay.Client `json:"client"`
	ctx        context.Context
}

func NewSdk(ctx context.Context, appId string, privateKey string, PublicKey string) (*AliPaySdk, error) {
	//初始化
	// 初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(appId, privateKey, true)
	if err != nil {
		fmt.Println("alipayNewClientErr:", err.Error())
		return nil, err
	}

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	globalExcPath := utils.GetGlobalExcPath("./")
	appCertPublicKey := globalExcPath + `etc/cert/alipay/` + appId + `/appCertPublicKey_` + appId + `.crt`
	alipayRootCertPath := globalExcPath + `etc/cert/alipay/` + appId + `/alipayRootCert.crt`
	alipayCertPublicKeyRSA2 := globalExcPath + `etc/cert/alipay/` + appId + `/alipayCertPublicKey_RSA2.crt`
	rootCertSN, err := alipay.GetRootCertSN(alipayRootCertPath)
	certSN, err := alipay.GetCertSN(appCertPublicKey)
	publicCertSN, err := alipay.GetCertSN(alipayCertPublicKeyRSA2)

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation("Asia/Shanghai"). // 设置时区，不设置或出错均为默认服务器时间
						SetAliPayRootCertSN(rootCertSN).     // 设置支付宝根证书SN，通过 alipaySdk.GetRootCertSN() 获取
						SetAppCertSN(certSN).                // 设置应用公钥证书SN，通过 alipaySdk.GetCertSN() 获取
						SetAliPayPublicCertSN(publicCertSN). // 设置支付宝公钥证书SN，通过 alipaySdk.GetCertSN() 获取
						SetCharset("utf-8").                 // 设置字符编码，不设置默认 utf-8
						SetSignType(alipay.RSA2).            // 设置签名类型，不设置默认 RSA2
						SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
						SetNotifyUrl("https://www.fmm.ink"). // 设置异步通知URL
						SetAppAuthToken("")                  // 设置第三方应用授权

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	//client.AutoVerifySign("alipayCertPublicKey_RSA2 bytes")

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath(appCertPublicKey, alipayRootCertPath, alipayCertPublicKeyRSA2)
	if err != nil {
		fmt.Println("clientSetCertSnByPathErr:", err.Error())
		return nil, err
	}

	return &AliPaySdk{
		AppId:      appId,
		PrivateKey: privateKey,
		PublicKey:  PublicKey,
		Client:     client,
		ctx:        ctx,
	}, nil
}

//转账
func (o *AliPaySdk) Transfer(identity string, name string, outBizNo string, transAmount float64) (*alipay.FundTransUniTransferResponse, error) {
	//转账
	mp := make(map[string]interface{})
	mp["identity_type"] = "ALIPAY_LOGON_ID"
	mp["identity"] = identity
	mp["name"] = name
	bm := make(gopay.BodyMap)
	bm.Set("trans_amount", transAmount).
		Set("out_biz_no", outBizNo).
		Set("product_code", "TRANS_ACCOUNT_NO_PWD").
		Set("order_title", "转账给"+name).
		Set("remark", "转账给"+name).
		Set("biz_scene", "DIRECT_TRANSFER").
		Set("payee_info", mp)

	res, err := o.Client.FundTransUniTransfer(o.ctx, bm)
	if err != nil {
		xlog.Errorf("client.FundTransUniTransfer(%+v),error:%+v", bm, err)
		return nil, err
	}
	xlog.Debug("Transfer:", res)
	return res, nil
}

//支付
func (o *AliPaySdk) TradeWapPay() {
	//支付
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "手机网站测试支付").
		Set("out_trade_no", "GZ201909081743431443").
		Set("quit_url", "https://www.fmm.ink").
		Set("total_amount", "3.00").
		Set("product_code", "QUICK_WAP_WAY")

	// 手机网站支付请求
	payUrl, err := o.Client.TradeWapPay(o.ctx, bm)
	if err != nil {
		xlog.Errorf("client.TradeWapPay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("payUrl:", payUrl)
}

//转账业务单据查询接口
func (o *AliPaySdk) FundTransOrderQuery(outBizNo string) (*alipay.FundTransOrderQueryResponse, error) {
	//支付
	// 请求参数
	bm := make(gopay.BodyMap).
		Set("out_biz_no", outBizNo)

	// 手机网站支付请求
	res, err := o.Client.FundTransOrderQuery(o.ctx, bm)
	if err != nil {
		xlog.Errorf("client.FundTransOrderQuery(%+v),error:%+v", bm, err)
		return nil, err
	}
	xlog.Debug("FundTransOrderQuery:", res)
	return res, nil
}

//转账
func (o *AliPaySdk) AlipayUserDtbankcustChannelvoucherSend(logonId string, phoneId string, activityId string, outBizNo string) (*alipay.AlipayUserDtbankcustChannelvoucherSendResponse, error) {
	//转账
	//mp := make(map[string]interface{})
	//mp["identity_type"] = "ALIPAY_LOGON_ID"
	//mp["identity"] = identity
	//mp["name"] = name
	bm := make(gopay.BodyMap)
	bm.Set("phone_id", phoneId).
		Set("logon_id", logonId).
		Set("activity_id", activityId).
		Set("out_biz_no", outBizNo)

	res, err := o.Client.AlipayUserDtbankcustChannelvoucherSend(o.ctx, bm)
	if err != nil {
		xlog.Errorf("client.AlipayUserDtbankcustChannelvoucherSend(%+v),error:%+v", bm, err)
		return nil, err
	}

	xlog.Debug("Transfer:", res)
	return res, nil
}

// 数字分行红包活动配置查询接口
func (o *AliPaySdk) AlipayUserDtbankcustChannelvoucherconfigQuery(activityId string) (*alipay.AlipayUserDtbankcustChannelvoucherconfigQueryResponse, error) {
	bm := gopay.BodyMap{
		"activity_id": activityId,
	}
	res, err := o.Client.AlipayUserDtbankcustChannelvoucherconfigQuery(o.ctx, bm)
	if err != nil {
		xlog.Errorf("client.AlipayUserDtbankcustChannelvoucherconfigQuery(%+v),error:%+v", bm, err)
		return nil, err
	}

	xlog.Debug("Transfer:", res)
	return res, nil
}
