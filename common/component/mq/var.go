package mq

import "fmt"

//公共队列名称，多服务共用的

//微信立减金支付宝红包发送
type BankGoodsCouponNotifyStruct struct {
	OrderNo      string `json:"order_no"`      //订单号
	SubOrderNo   string `json:"sub_order_no"`  //子订单号
	SupplierId   int64  `json:"supplier_id"`   //供应商
	StockId      string `json:"stock_id"`      //批次号
	CouponId     string `json:"coupon_id"`     //优惠ID
	CardType     int64  `json:"card_type"`     //卡券类型，1链接，2图片(二维码)，3卡密，4密码、5条形码、6联联
	CardNumber   string `json:"card_number"`   //卡号
	CardPassword string `json:"card_password"` //卡密
	CardLink     string `json:"card_link"`     //卡链接
	Status       int64  `json:"status"`        //状态：0：待发券，1发券成功，2发券失败
	ErrMsg       string `json:"err_msg"`       //错误提示
	Time         string `json:"time"`          //时间
	Times        int64  `json:"times"`         //次数默认1
}

//微信立减金支付宝红包核销通知
type BankGoodsCouponConsumStruct struct {
	SubOrderNo  string `json:"sub_order_no"` //子订单号
	ConsumeTime int64  `json:"consume_time"` //核销时间
	Time        string `json:"time"`         //时间
	Times       int64  `json:"times"`        //次数默认1
}

//获取微信立减金支付宝红包发送结果队列名称
func GetCouponNotifyQueueName(env string) string {
	return fmt.Sprintf("bankGoodsCouponNotify_%s", env)
}

//获取微信立减金支付宝红包核销队列名称
func GetCouponConsumQueueName(env string) string {
	return fmt.Sprintf("bankGoodsCouponConsum_%s", env)
}

//延迟分账队列名称
func GetSharingDelayQueueName(env string) string {
	return fmt.Sprintf("bankSharingDealy_%s", env)
}

// 延迟队列名称
func GetRefundDelayQueueName(env string) string {
	return fmt.Sprintf("bankRefundDealy_%s", env)
}

// 采购卡券队列名称
func GetCouponPurchaseQueueName(env string) string {
	return fmt.Sprintf("bankCouponPurchase_%s", env)
}

//定时分账队列名称
func GetSharingCronQueueName(env string) string {
	return fmt.Sprintf("bankSharingCron_%s", env)
}

//购买订单优惠卡券核销
func GetBuyCouponCodeConsumName(env string) string {
	return fmt.Sprintf("bankBuyCouponCodeConsum_%s", env)
}
