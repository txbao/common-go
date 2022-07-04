package mqueue

type ReleaseStock struct {
	OrderNo string `json:"order_no"` //订单号
	Time    string `json:"time"`     //时间
	Times   int64  `json:"times"`    //次数默认1
}

//订单服务
const (
	DelayTime5Min = 300 //延迟5分钟
)
const (
	//库存释放延迟队列
	ReleaseStockQueue = "BA.orderRpc.delayReleaseStock.%s"
)
