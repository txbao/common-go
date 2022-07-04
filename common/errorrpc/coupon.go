package errorrpc

//卡券服务code
const CouponError = 15000

var (
	CouponCreateErr        = NewCodeError(CouponError+1, "卡券生成错误")
	CouponNameExistErr     = NewCodeError(CouponError+2, "卡券批次名称已存在")
	CouponRelationErr      = NewCodeError(CouponError+3, "关联失败")
	CouponGiveErr          = NewCodeError(CouponError+4, "卡券领取失败")
	CouponConsumErr        = NewCodeError(CouponError+5, "卡券核销失败")
	CouponDiscountQueryErr = NewCodeError(CouponError+6, "卡券优惠查询失败")
	CouponSaveOrderInfoErr = NewCodeError(CouponError+7, "卡券保存订单失败")
	CouponMultipleErr      = NewCodeError(CouponError+8, "卡券多发处理失败")
	CouponExchangeErr      = NewCodeError(CouponError+9, "卡券兑换失败")
	UpdateStatusByCodeErr  = NewCodeError(CouponError+10, "卡券券码状态更改失败")
	CouponAppendErr        = NewCodeError(CouponError+11, "卡券追加错误")
	CouponAppendOverrunErr        = NewCodeError(CouponError+12, "卡券追加次数超限")
)
