package errorrpc

//支付服务code
const PAY_ERROR = 14000

var (
	PayActivityNoActive       = NewCodeError(PAY_ERROR+1, "没有可用的优惠信息")
	PayActivityOutRangeTime   = NewCodeError(PAY_ERROR+2, "该时间段内暂无活动参与")
	PayActivityOutRangeMobile = NewCodeError(PAY_ERROR+3, "该用户暂无资格参与活动")
	PayParamsValidErr         = NewCodeError(PAY_ERROR+4, "参数验证错误")
	PayInitErr                = NewCodeError(PAY_ERROR+5, "支付SDK错误")
	PayRefundErr              = NewCodeError(PAY_ERROR+6, "退款错误")
	PaySharingErr             = NewCodeError(PAY_ERROR+7, "分账错误")

	//ActivityOutRangeMobile = NewCodeError(PAY_ERROR+3, "该用户暂无资格参与活动")

)
