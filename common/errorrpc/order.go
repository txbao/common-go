package errorrpc

//订单服务code
const ORDER_ERROR = 13000

var (
	OrderLianLianOrderErr       = NewCodeError(ORDER_ERROR+1, "联联下单失败")
	OrderConsumErr              = NewCodeError(ORDER_ERROR+2, "核销失败")
	OrderRequestNOErr           = NewCodeError(ORDER_ERROR+3, "领取机会已用完")
	OrderLimitErr               = NewCodeError(ORDER_ERROR+4, "次数超限")
	ExpressErr                  = NewCodeError(ORDER_ERROR+5, "物流查询错误")
	ExpressDeliverErr           = NewCodeError(ORDER_ERROR+6, "物流发货错误")
	ExpressParameterErr         = NewCodeError(ORDER_ERROR+7, "主订单与子订单不能同时为空")
	OrderUpdateErr              = NewCodeError(ORDER_ERROR+7, "订单修改错误")
	FindOneBySkuIdActivityIdErr = NewCodeError(ORDER_ERROR+8, "根据SKU&活动ID来获取订单信息错误")
	StockNotEnoughErr           = NewCodeError(ORDER_ERROR+9, "库存不足")
)
