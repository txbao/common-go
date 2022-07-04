package orderVar

const (
	//状态：0:待支付，2:已支付，4:已采购(已预约)，6:已发货，8：已收货（已核销）,10:已完成，12:已取消（含已退款）,14采购失败
	StatusWaitPay         = 0  //0：等待支付
	StatusPaySuccess      = 2  //5:2:已支付
	StatusPurchaseSuccess = 4  //4:已采购(已预约)
	StatusDelivered       = 6  //6:已发货
	StatusConsume         = 8  //8：已收货（已核销）
	StatusComplete        = 10 //10:已完成
	StatusCannel          = 12 //12:已取消（含已退款）
	StatusCancelTimeOut   = 13 //13:已取消（超时）
	StatusPurchaseFail    = 14 //14采购(发券)失败
)
