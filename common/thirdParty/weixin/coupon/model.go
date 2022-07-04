package coupon

type FixedNormalCoupon struct {
	CouponAmount       int64 `json:"coupon_amount"`       // 面额;单位：分。示例值：100
	TransactionMinimum int64 `json:"transaction_minimum"` // 门槛;使用券金额门槛，单位：分。示例值：100
}

type StockUseRule struct {
	MaxCoupons        int64             `json:"max_coupons"`          // Y 发放总上限;最大发券数 示例值：100
	MaxAmount         int64             `json:"max_amount"`           // Y 总预算;总消耗金额，单位：分。示例值：5000
	MaxAmountByDay    int64             `json:"max_amount_by_day"`    // Y 单天发放上限金额;单天最高消耗金额，单位：分。示例值：400
	FixedNormalCoupon FixedNormalCoupon `json:"fixed_normal_coupon"`  // N 固定面额批次特定信息;固定面额发券批次特定信息。
	MaxCouponsPerUser int64             `json:"max_coupons_per_user"` // Y 单个用户可领个数; 示例值：3
	CouponType        string            `json:"coupon_type"`          // N 券类型; 枚举值： NORMAL：满减券 CUT_TO：减至券
	GoodsTag          []string          `json:"goods_tag"`            // N 订单优惠标记;订单优惠标记 (该字段暂未开放返回) 特殊规则：单个优惠标记的字符长度为【1，128】,条目个数限制为【1，50】。 示例值：{'123456','23456'}
	TradeType         []string          `json:"trade_type"`           // Y 支付方式; 默认不限制，枚举值： MICROAPP：小程序支付 APPPAY：APP支付 PPAY：免密支付 CARD：付款码支付 FACE：人脸支付 OTHER：（公众号、扫码等） 示例值： ["OTHER","APPPAY"]
	CombineUse        bool              `json:"combine_use"`          // N 是否可叠加其他优惠; 枚举值： true：是 false：否 示例值：true
}

type CutToMessage struct {
	SinglePriceMax int64 `json:"single_price_max"` // Y 可用优惠的商品最高单价; 可用优惠的商品最高单价，单位：分。 示例值：100
	CutToPrice     int64 `json:"cut_to_price"`     // Y 减至后的优惠单价; 减至后的优惠单价，单位：分。 示例值：100
}

type StockInfo struct {
	StockID            string       `json:"stock_id"`             // Y 批次号;微信为每个代金券批次分配的唯一id。示例值：9836588
	StockCreatorMchid  string       `json:"stock_creator_mchid"`  // Y 创建批次的商户号;微信为创建方商户分配的商户号。示例值：123456
	StockName          string       `json:"stock_name"`           // Y 批次名称;示例值：微信支付批次
	Status             string       `json:"status"`               // Y 批次状态;枚举值：unactivated：未激活 audit：审核中 running：运行中 stoped：已停止 paused：暂停发放
	CreateTime         string       `json:"create_time"`          // Y 创建时间;批次创建时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。
	Description        string       `json:"description"`          // Y 使用说明;批次描述信息 示例值：微信支付营销
	StockUseRule       StockUseRule `json:"stock_use_rule"`       // N 满减券批次使用规则;普通发券批次特定信息。
	AvailableBeginTime string       `json:"available_begin_time"` // Y 可用开始时间; 遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。 示例值：2015-05-20T13:29:35.120+08:00
	AvailableEndTime   string       `json:"available_end_time"`   // Y 可用结束时间; 遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。 示例值：2015-05-20T13:29:35.120+08:00
	DistributedCoupons int64        `json:"distributed_coupons"`  // Y 已发券数量; 示例值：100
	NoCash             bool         `json:"no_cash"`              // Y 是否无资金流; 枚举值： true：是 false：否 示例值：true
	StartTime          string       `json:"start_time"`           // N 激活批次的时间; 批次激活开启时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。 示例值：2015-05-20T13:29:35.120+08:00
	StopTime           string       `json:"stop_time"`            // N 终止批次的时间; 批次永久停止时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。 示例值：2015-05-20T13:29:35.120+08:00
	CutToMessage       CutToMessage `json:"cut_to_message"`       // N 减至批次特定信息; 单品优惠特定信息。
	Singleitem         bool         `json:"singleitem"`           // Y 是否单品优惠; 枚举值： true：是 false：否 示例值：true
	StockType          string       `json:"stock_type"`           // Y 批次类型; 枚举值： NORMAL：代金券批次 DISCOUNT_CUT：立减与折扣 OTHER：其他 示例值：NORMAL
}

type StockList struct {
	TotalCount int64        `json:"total_count"` // Y 批次总数;经过条件筛选，查询到的批次总数量。 示例值：10
	Data       []*StockInfo `json:"data"`        // N 批次详情;
	Limit      int64        `json:"limit"`       // Y 分页大小;最大10。示例值：8
	Offset     int64        `json:"offset"`      // Y 分页页码;页码从0开始，默认第0页。示例值：1
}

type StockListReq struct {
	Offset            int64  `json:"offset"`              // Y 分页页码; 页码从0开始，默认第0页。 示例值：0
	Limit             int64  `json:"limit"`               // Y 分页大小; 分页大小，最大10。 示例值：8
	StockCreatorMchid string `json:"stock_creator_mchid"` // Y 创建批次的商户号; 批次创建方商户号。校验规则：接口传入的批次号需由stock_creator_mchid所创建。 示例值：9856888
	CreateStartTime   string `json:"create_start_time"`   // N 起始时间; 起始创建时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。 校验规则：get请求，参数在 url中，需要进行 url 编码传递 示例值：2015-05-20T19%3A29%3A35.120%2B08%3A00
	CreateEndTime     string `json:"create_end_time"`     // N 终止时间; 终止创建时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示，北京时间2015年5月20日 13点29分35秒。 校验规则：get请求，参数在 url中，需要进行 url 编码传递 示例值：2015-05-20T19%3A29%3A35.120%2B08%3A00
	Status            string `json:"status"`              // N 批次状态; 枚举值： unactivated：未激活 audit：审核中 running：运行中 stoped：已停止 paused：暂停发放 示例值：paused
}

type SendCouponReq struct {
	Openid            string `json:"-"`                   // Y 用户openid;openid信息，用户在appid下的唯一标识。 校验规则：该openid需要与接口传入中的appid有对应关系。 示例值：2323dfsdf342342
	StockId           string `json:"stock_id"`            // Y 批次号;微信为每个批次分配的唯一id。 校验规则：必须为代金券（全场券或单品券）批次号，不支持立减与折扣。 示例值：9856000
	OutRequestNo      string `json:"out_request_no"`      // Y 商户单据号;商户此次发放凭据号（格式：商户id+日期+流水号），可包含英文字母，数字，|，_，*，-等内容，不允许出现其他不合法符号，商户侧需保持唯一性。 示例值： 89560002019101000121
	Appid             string `json:"appid"`               // Y 公众账号ID;微信为发券方商户分配的公众账号ID，接口传入的所有appid应该为公众号的appid或者小程序的appid（在mp.weixin.qq.com申请的）或APP的appid（在open.weixin.qq.com申请的）。 校验规则： 1、该appid需要与接口传入中的openid有对应关系； 2、该appid需要与调用接口的商户号（即请求头中的商户号）有绑定关系，若未绑定，可参考该指引完成绑定（商家商户号与AppID账号关联管理） 示例值：wx233544546545989
	StockCreatorMchid string `json:"stock_creator_mchid"` // Y 创建批次的商户号;批次创建方商户号。 校验规则：接口传入的批次号需由stock_creator_mchid所创建。 示例值：8956000
}

type SendCouponResp struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	CouponId string `json:"coupon_id"`
}

//查询代金券批次信息
type StockInfoReq struct {
	StockId           string `json:"stock_id"`
	StockCreatorMchid string `json:"stock_creator_mchid"`
}
type StockInfoRes struct {
	Code               string `json:"code"`
	Message            string `json:"message"`
	AvailableBeginTime string `json:"available_begin_time"`
	AvailableEndTime   string `json:"available_end_time"`
	CardID             string `json:"card_id"`
	CreateTime         string `json:"create_time"`
	Description        string `json:"description"`
	DistributedCoupons int64  `json:"distributed_coupons"`
	NoCash             bool   `json:"no_cash"`
	Singleitem         bool   `json:"singleitem"`
	Status             string `json:"status"`
	StockCreatorMchid  string `json:"stock_creator_mchid"`
	StockID            string `json:"stock_id"`
	StockName          string `json:"stock_name"`
	StockType          string `json:"stock_type"`
	StockUseRule       struct {
		CombineUse        bool   `json:"combine_use"`
		CouponType        string `json:"coupon_type"`
		FixedNormalCoupon struct {
			CouponAmount       int64 `json:"coupon_amount"`
			TransactionMinimum int64 `json:"transaction_minimum"`
		} `json:"fixed_normal_coupon"`
		GoodsTag          []interface{} `json:"goods_tag"`
		MaxAmount         int64         `json:"max_amount"`
		MaxAmountByDay    int64         `json:"max_amount_by_day"`
		MaxCoupons        int64         `json:"max_coupons"`
		MaxCouponsPerUser int64         `json:"max_coupons_per_user"`
		TradeType         []interface{} `json:"trade_type"`
	} `json:"stock_use_rule"`
}

//查询优惠ID详情
type CouponInfoReq struct {
	MchID    string `json:"mch_id"`
	AppId    string `json:"app_id"`
	Openid   string `json:"openid"`
	CouponId string `json:"coupon_id"`
}
type CouponInfoRes struct {
	Code                    string `json:"code"`
	Message                 string `json:"message"`
	AvailableBeginTime      string `json:"available_begin_time"`
	AvailableEndTime        string `json:"available_end_time"`
	CouponID                string `json:"coupon_id"`
	CouponName              string `json:"coupon_name"`
	CouponType              string `json:"coupon_type"`
	CreateTime              string `json:"create_time"`
	Description             string `json:"description"`
	NoCash                  bool   `json:"no_cash"`
	NormalCouponInformation struct {
		CouponAmount       int64 `json:"coupon_amount"`
		TransactionMinimum int64 `json:"transaction_minimum"`
	} `json:"normal_coupon_information"`
	Singleitem        bool   `json:"singleitem"`
	Status            string `json:"status"`
	StockCreatorMchid string `json:"stock_creator_mchid"`
	StockID           string `json:"stock_id"`
}

//下载批次核销明细API
type StockUseFlowReq struct {
	MchID   string `json:"mch_id"`
	StockId string `json:"stock_id"`
}
type StockUseFlowRes1 struct {
	HashType  string `json:"hash_type"`
	HashValue string `json:"hash_value"`
	URL       string `json:"url"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

//数据结构体
type StockUseFlowResData struct {
	StockId      string `json:"stock_id"`      //批次id
	CouponId     string `json:"coupon_id"`     //优惠id
	StockType    string `json:"stock_type"`    //优惠类型
	CouponAmount string `json:"coupon_amount"` //优惠金额（元）
	TotalAmount  string `json:"total_amount"`  //订单总金额（元）
	TradeType    string `json:"trade_type"`    //交易类型
	OutTradeNo   string `json:"out_trade_no"`  //支付单号
	ConsumTime   string `json:"consum_time"`   //消耗时间
	ConsumMchid  string `json:"consum_mchid"`  //消耗商户号
	DevNo        string `json:"dev_no"`        //设备号
	BankTradeNo  string `json:"bank_trade_no"` //银行流水号
	GoodsDetail  string `json:"goods_detail"`  //单品信息
}
type StockUseFlowRes2 struct {
	Count int                   `json:"count"`
	Data  []StockUseFlowResData `json:"data"`
}

//下载批次退款明细API
type StockRefundFlowReq struct {
	MchID   string `json:"mch_id"`
	StockId string `json:"stock_id"`
}

type StockRefundFlowRes1 struct {
	HashType  string `json:"hash_type"`
	HashValue string `json:"hash_value"`
	URL       string `json:"url"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

//数据结构体
type StockRefundFlow2DataStruct struct {
	StockId      string `json:"stock_id"`      //批次id
	CouponId     string `json:"coupon_id"`     //优惠id
	StockType    string `json:"stock_type"`    //优惠类型
	CouponAmount string `json:"coupon_amount"` //优惠金额（元）
	TotalAmount  string `json:"total_amount"`  //订单总金额（元）
	TradeType    string `json:"trade_type"`    //交易类型
	OutTradeNo   string `json:"out_trade_no"`  //支付单号
	ConsumTime   string `json:"consum_time"`   //消耗时间
	ConsumMchid  string `json:"consum_mchid"`  //消耗商户号
	DevNo        string `json:"dev_no"`        //设备号
	BankTradeNo  string `json:"bank_trade_no"` //银行流水号
	GoodsDetail  string `json:"goods_detail"`  //单品信息
}
type StockRefundFlowRes2 struct {
	Count int                          `json:"count"`
	Data  []StockRefundFlow2DataStruct `json:"data"`
}

//设置消息通知地址API
type SetCallbackUrlReq struct {
	Mchid     string `json:"mchid"`
	NotifyUrl string `json:"notify_url"`
	Wwitch    bool   `json:"wwitch"`
}
type SetCallbackUrlRes struct {
	NotifyURL  string `json:"notify_url"`
	UpdateTime string `json:"update_time"`
}

//核销事件回调通知API
type ConsumeNotifyReq struct {
	CreateTime string `json:"create_time"`
	EventType  string `json:"event_type"`
	ID         string `json:"id"`
	Resource   struct {
		Algorithm      string `json:"algorithm"`
		AssociatedData string `json:"associated_data"`
		Ciphertext     string `json:"ciphertext"`
		Nonce          string `json:"nonce"`
		OriginalType   string `json:"original_type"`
	} `json:"resource"`
	ResourceType string `json:"resource_type"`
	Summary      string `json:"summary"`
}
type ConsumeNotifyRes struct {
	AvailableBeginTime string `json:"available_begin_time"`
	AvailableEndTime   string `json:"available_end_time"`
	ConsumeInformation struct {
		ConsumeMchid string `json:"consume_mchid"`
		ConsumeTime  string `json:"consume_time"`
		GoodsDetail  []struct {
			DiscountAmount int64  `json:"discount_amount"`
			GoodsID        string `json:"goods_id"`
			Price          int64  `json:"price"`
			Quantity       int64  `json:"quantity"`
		} `json:"goods_detail"`
		TransactionID string `json:"transaction_id"`
	} `json:"consume_information"`
	CouponID    string `json:"coupon_id"`
	CouponName  string `json:"coupon_name"`
	CouponType  string `json:"coupon_type"`
	CreateTime  string `json:"create_time"`
	Description string `json:"description"`
	DiscountTo  struct {
		CutToPrice int64 `json:"cut_to_price"`
		MaxPrice   int64 `json:"max_price"`
	} `json:"discount_to"`
	NoCash                  bool `json:"no_cash"`
	NormalCouponInformation struct {
		CouponAmount       int64 `json:"coupon_amount"`
		TransactionMinimum int64 `json:"transaction_minimum"`
	} `json:"normal_coupon_information"`
	Singleitem            bool `json:"singleitem"`
	SingleitemDiscountOff struct {
		SinglePriceMax int64 `json:"single_price_max"`
	} `json:"singleitem_discount_off"`
	Status            string `json:"status"`
	StockCreatorMchid string `json:"stock_creator_mchid"`
	StockID           string `json:"stock_id"`
}
