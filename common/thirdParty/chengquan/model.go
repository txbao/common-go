package chengquan

// 公共返回参数
type CommonResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// 卡券品牌列表
type CouponType struct {
	GoodsTypeId   int64  `json:"goods_type_id"`
	GoodsTypeName string `json:"goods_type_name"`
	SellStatus    int64  `json:"sell_status"`
}
type CouponTypeListResponse struct {
	CommonResponse
	Data []CouponType `json:"data"`
}

// 卡券品牌产品列表
type CouponGoods struct {
	Name          string  `json:"name"`
	GoodsNo       string  `json:"goods_no"`
	Price         float64 `json:"price"`
	OfficialPrice float64 `json:"official_price"`
	GoodsTypeId   int64   `json:"goods_type_id"`
	GoodsTypeName string  `json:"goods_type_name"`
	SellStatus    int64   `json:"sell_status"`
}
type CouponTypeGoodsListRequest struct {
	GoodsTypeId int64 `json:"goods_type_id"`
}
type CouponTypeGoodsListResponse struct {
	CommonResponse
	Data []CouponGoods `json:"data"`
}

// 卡券产品库存
type CouponGoodsStock struct {
	Name     string `json:"name"`
	GoodsNo  string `json:"goods_no"`
	StockNum int64  `json:"stock_num"`
}
type CouponGoodsStockRequest struct {
	GoodsNo string `json:"goods_no"`
}
type CouponGoodsStockResponse struct {
	CommonResponse
	Data CouponGoodsStock `json:"data"`
}

// 卡券使用须知
type CouponGoodsNotesRequest struct {
	GoodsNo string `json:"goods_no"`
}
type CouponGoodsNotesResponse struct {
	CommonResponse
	Data string `json:"data"`
}

// 卡券购买接口
type CouponOrderPayRequest struct {
	UserOrderNo string `json:"user_order_no"`
	GoodsNo     string `json:"goods_no"`
	Count       int64  `json:"count"`
	Mobile      string `json:"mobile"`
}
type CouponOrderPayResponse struct {
	Code int64 `json:"code"`
	Data []struct {
		UserOrderNo   string  `json:"user_order_no"`  //商户提交订单号
		OrderNo       string  `json:"order_no"`       //橙券平台订单号
		State         string  `json:"state"`          //状态码 	状态码说明 SUCCESS 	成功 FAILURE 	失败
		Price         float64 `json:"price"`          //单价(单位：元)，保留小数点后四位
		GoodsNo       string  `json:"goods_no"`       //产品编号
		GoodsName     string  `json:"goods_name"`     //产品名称
		GoodsType     string  `json:"goods_type"`     //产品类型(LINK 链接，PICTURE 图片，NUMBER_PASSWORD 卡号和密码，PASSWORD 密码)
		GoodsNumber   string  `json:"goods_number"`   //卡号(产品类型为NUMBER_PASSWORD时有值)，AES加密
		GoodsPassword string  `json:"goods_password"` //卡密(产品类型为NUMBER_PASSWORD或PASSWORD时有值)，AES加密
		GoodsLink     string  `json:"goods_link"`     //链接(产品类型为LINK或PICTURE时有值)，AES加密
		CompleteTime  int64   `json:"complete_time"`  //订单创建时间(时间戳，单位：毫秒)
		CreateTime    int64   `json:"create_time"`    ///订单完成时间(时间戳，单位：毫秒)
		EffectiveTime int64   `json:"effective_time"` //有效期(时间戳，单位：毫秒)
	} `json:"data"`
	Message string `json:"message"`
}

//卡券订单状态查询接口
type CouponOrderQueryRequest struct {
	UserOrderNo string `json:"user_order_no"`
}
