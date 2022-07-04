package errorrpc

//基础服务code
const BASE_ERROR = 10000

var (
	BaseSmsInitErr         = NewCodeError(BASE_ERROR+1, "短信初始化错误")
	BaseSmsCodeErr         = NewCodeError(BASE_ERROR+2, "验证码错误")
	BaseSmsSendErr         = NewCodeError(BASE_ERROR+3, "短信发送错误")
	BaseAuthJumpPlateErr   = NewCodeError(BASE_ERROR+5, "跳转板块错误")
	BaseAuthModeErr        = NewCodeError(BASE_ERROR+6, "授权模式错误")
	BaseDecryptErr         = NewCodeError(BASE_ERROR+7, "解密错误")
	BaseDBErr              = NewCodeError(BASE_ERROR+8, "DB错误")
	BaseThirdMobileInfoErr = NewCodeError(BASE_ERROR+9, "第三方获取手机信息错误")
	WhiteInterfaceErr      = NewCodeError(BASE_ERROR+10, "白名单接口错误")
	BaseSaveLogErr      = NewCodeError(BASE_ERROR+11, "保存日志失败")
)
