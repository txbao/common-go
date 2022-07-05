package errorrpc

//卡券系统服务code
const CouponSystemError = 16000

var (
	//CouponSystemBatchMakeFail = "制券失败"

	CouponSystemBatchMakeFail       = NewCodeError(CouponSystemError+1, "制券失败")
	CouponSystemNameExistErr        = NewCodeError(CouponSystemError+2, "卡券批次名称已存在")
	CouponSystemFrequentRequestsErr = NewCodeError(CouponSystemError+3, "请求频繁，请稍后再试！")

	//CouponRelationErr      = NewCodeError(CouponSystemError+3, "关联失败")
	//CouponGiveErr          = NewCodeError(CouponSystemError+4, "卡券领取失败")
	//CouponConsumErr        = NewCodeError(CouponSystemError+5, "卡券核销失败")
	//CouponDiscountQueryErr = NewCodeError(CouponSystemError+6, "卡券优惠查询失败")
	//CouponSaveOrderInfoErr = NewCodeError(CouponSystemError+7, "卡券保存订单失败")
	//CouponMultipleErr      = NewCodeError(CouponSystemError+8, "卡券多发处理失败")
	//CouponExchangeErr      = NewCodeError(CouponSystemError+9, "卡券兑换失败")
	//UpdateStatusByCodeErr  = NewCodeError(CouponSystemError+10, "卡券券码状态更改失败")
	//CouponAppendErr        = NewCodeError(CouponSystemError+11, "卡券追加错误")
	//CouponAppendOverrunErr        = NewCodeError(CouponSystemError+12, "卡券追加次数超限")
)
