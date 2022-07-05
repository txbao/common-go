package scm

type SupplierListResponse struct {
	Code    int64  `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    []struct {
		SupplierId   int64  `json:"supplierId"`
		SupplierName string `json:"supplierName"`
	} `json:"data"`
}

type SpuListResponse struct {
	Code    int64  `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		CurrPage   int64 `json:"currPage"`
		PageSize   int64 `json:"pageSize"`
		TotalCount int64 `json:"totalCount"`
		TotalPage  int64 `json:"totalPage"`
		List       []struct {
			Id           int64       `json:"id"`
			SupplierId   int64       `json:"supplierId"`
			SupplierName string      `json:"supplierName"`
			SpuName      string      `json:"spuName"`
			BrandId      int64       `json:"brandId"`
			BrandName    string      `json:"brandName"`
			UseNotice    string      `json:"useNotice"`
			UseStatus    int64       `json:"useStatus"`
			ProductCount int64       `json:"productCount"`
			Products     interface{} `json:"products"`
			CreateTime   string      `json:"createTime"`
			UpdateTime   string      `json:"updateTime"`
			CreateId     int64       `json:"createId"`
			UpdateId     int64       `json:"updateId"`
			Remarks      string      `json:"remarks"`
			Enable       int64       `json:"enable"`
			Version      int64       `json:"version"`
		} `json:"list"`
	} `json:"data"`
}

type SkuListResponse struct {
	Code    int64  `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    []struct {
		Id                  int64       `json:"id"`
		ProductCode         string      `json:"productCode"`
		ProductName         string      `json:"productName"`
		RechangeProductType interface{} `json:"rechangeProductType"`
		SpuId               int64       `json:"spuId"`
		SpuName             string      `json:"spuName"`
		BrandId             int64       `json:"brandId"`
		BrandName           string      `json:"brandName"`
		ProductIcon         string      `json:"productIcon"`
		ProductPic          string      `json:"productPic"`
		SupplierId          int64       `json:"supplierId"`
		SupplierName        string      `json:"supplierName"`
		OfficialPrice       int64       `json:"officialPrice"`
		Price               int64       `json:"price"`
		SqPrice             int64       `json:"sqPrice"`
		Stock               int64       `json:"stock"`
		SellStatus          int64       `json:"sellStatus"`
		ProductStatus       int64       `json:"productStatus"`
		CardType            interface{} `json:"cardType"`
		CreateTime          string      `json:"createTime"`
		UpdateTime          string      `json:"updateTime"`
		CreateId            int64       `json:"createId"`
		UpdateId            int64       `json:"updateId"`
		Remarks             string      `json:"remarks"`
		Enable              int64       `json:"enable"`
		Version             int64       `json:"version"`
		ValidityTime        int64       `json:"validityTime"`
		UseNotice           string      `json:"useNotice"`
		OtherRemark         string      `json:"otherRemark"`
		Disclaimer          string      `json:"disclaimer"`
		ProductDesc         string      `json:"productDesc"`
		ValidityNum         int64       `json:"validityNum"`
		PriceDivide100      interface{} `json:"priceDivide100"`
		SqPriceDivide100    interface{} `json:"sqPriceDivide100"`
	} `json:"data"`
}

type LaunchPerformanceReq struct {
	UserId          string `json:"userId"`          // 用户id
	TransactionNo   string `json:"transactionNo"`   // 卡号/支付交易流水号
	ProductId       int64  `json:"productId"`       // 商品id
	Quantity        int64  `json:"quantity"`        // 数量
	ActivityId      string `json:"activityId"`      // 活动Id
	TransactionType int64  `json:"transactionType"` // 交易类型【1.卡券、2.订单】
}
type ApplyPerformanceUrlRequest struct {
	RequestId    string               // 必填；事务Id 单次唯一 用于幂等（不超过20位长度字符串）
	ActivityName string               // 必填；活动名称
	ProductName  string               // 商品名称
	ProductPic   string               // 商品图片
	ProductDesc  string               // 商品描叙
	Data         LaunchPerformanceReq // 加密信息
}
type ApplyPerformanceUrlResponse struct {
	Code    int64  `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    string `json:"data"`
}
