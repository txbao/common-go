package code

//异常码规则 20_02_1_00001 为十三位
// 0为正确 ,异常为 20_02_1_00001
// 18为系统编号 支付平台、 17为在线收银系统 16为系统编号 合伙人
// 06 为模块或活动编辑，00：后台及公共异常 01：商品兑换码 02:微信立减金 03：虚拟卡兑换 04:兑换码核销 05权益活动 06：礼品兑换 以此类推
// 1为异常级别  1：错误 2：警告 3：其它
// 00001 为5位异常码

const (
	Ok           = "0"
	SysErr       = "400"
	LoginTimeout = "401"

	Limit               = "16_01_1_00001"
	ParamsError         = "16_01_1_00002"
	MobileError         = "16_01_1_00003"
	SmsCodeError        = "16_01_1_00004"
	SmsCodeStillValid   = "16_01_1_00005"
	LoginFail           = "16_01_1_00006"
	LoginFailed         = "16_01_1_00007"
	NonExistGoods       = "16_01_1_00010"
	NonExistRecords     = "16_01_1_00011"
	OperationFail       = "16_01_1_00012"
	WeixinTokenFail     = "16_01_1_00013"
	GetPhoneDecryptFail = "16_01_1_00014"
	TokenFail           = "16_01_1_00015"
	IDAuthFail          = "16_01_1_00016"
	RegisterFail        = "16_01_1_00017"
	ValidFail           = "16_01_1_00018"
	AchieveErr          = "16_01_1_00019"

	PartnerNonExist   = "16_01_1_00020"
	PartnerNoApproval = "16_01_1_00021"
	RepeatRequest     = "16_01_1_00022"
	PartnerIsDisable  = "16_01_1_00030"

	ApplyCardSignErr         = "16_01_1_00023"
	ApplyCardQrErr           = "16_01_1_00024"
	WithdrawErr              = "16_01_1_00025"
	UserInfoSaveErr          = "16_01_1_00026"
	PartnerDelFail           = "16_01_1_00027"
	PartnerChangeAccountFail = "16_01_1_00028"
	SignEditBankcardFail     = "16_01_1_00029"
)

var Map = map[string]string{
	Ok:                  "ok",
	SysErr:              "系统异常",
	Limit:               "全局限流",
	ParamsError:         "请求参数错误",
	MobileError:         "手机号码格式错误",
	SmsCodeError:        "验证码错误",
	SmsCodeStillValid:   "验证码依然有效",
	LoginFail:           "登录失败",
	LoginFailed:         "登录失败",
	LoginTimeout:        "登录超时",
	NonExistGoods:       "此活动没有商品",
	NonExistRecords:     "没有记录",
	OperationFail:       "操作失败",
	WeixinTokenFail:     "登录凭证校验异常！",
	GetPhoneDecryptFail: "获取手机号解密失败",
	TokenFail:           "获取Token失败",
	IDAuthFail:          "身份证格式错误",
	RegisterFail:        "注册失败",
	ValidFail:           "验证失败",
	AchieveErr:          "业绩统计错误",
	PartnerNonExist:     "合伙人还没注册",
	PartnerNoApproval:   "合伙人还未审核完成",
	RepeatRequest:       "请勿重复请求",

	ApplyCardSignErr:         "办卡签名有误",
	ApplyCardQrErr:           "申请办卡二维码有误",
	WithdrawErr:              "提现失败",
	UserInfoSaveErr:          "用户信息保存失败",
	PartnerDelFail:           "合伙人注销失败",
	PartnerChangeAccountFail: "合伙人切换失败",
	PartnerIsDisable:         "当前账户被禁用，请联系管理员",
	SignEditBankcardFail:     "修改签约的银行卡失败",
}
