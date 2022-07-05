package errorrpc

const SYS_ERROR = 20000

var (
	OK               = NewCodeError(0, "成功")
	Default          = NewCodeError(SYS_ERROR+1, "其他错误")
	TokenExpired     = NewCodeError(SYS_ERROR+2, "token 已经过期")
	TokenNotValidYet = NewCodeError(SYS_ERROR+3, "token还未生效")
	TokenMalformed   = NewCodeError(SYS_ERROR+4, "这不是一个token")
	TokenInvalid     = NewCodeError(SYS_ERROR+5, "违法的token")
	Parameter        = NewCodeError(SYS_ERROR+6, "参数错误")
	System           = NewCodeError(SYS_ERROR+7, "系统繁忙")
	Database         = NewCodeError(SYS_ERROR+8, "数据库错误")
	NotFind          = NewCodeError(SYS_ERROR+9, "未查询到")
	Duplicate        = NewCodeError(SYS_ERROR+10, "参数重复")
	SignatureExpired = NewCodeError(SYS_ERROR+11, "签名已经过期")
	Permissions      = NewCodeError(SYS_ERROR+12, "权限不足")
	Method           = NewCodeError(SYS_ERROR+13, "method不支持")
	Type             = NewCodeError(SYS_ERROR+14, "参数的类型不对")
	OutRange         = NewCodeError(SYS_ERROR+15, "参数的值超出范围")
	TimeOut          = NewCodeError(SYS_ERROR+16, "等待超时")
	Server           = NewCodeError(SYS_ERROR+17, "本实例处理不了该信息")
	DBUpdateFail     = NewCodeError(SYS_ERROR+18, "数据库更新失败")
)
