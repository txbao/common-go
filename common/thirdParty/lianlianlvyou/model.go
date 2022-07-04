package lianlianlvyou

// 产品item
type ProductItem struct {
	ID                    int64  `json:"id"`
	SubTitle              string `json:"subTitle"`
	ItemChannelStock      int64  `json:"itemChannelStock"`
	ChannelPrice          int64  `json:"channelPrice"`
	ItemChannelSaleAmount int64  `json:"itemChannelSaleAmount"`
	SalePrice             int64  `json:"salePrice"`
	OriginPrice           int64  `json:"originPrice"`
	SingleMax             int64  `json:"singleMax"`
	AllAreaID             int64  `json:"allAreaId"`
	Memo                  string `json:"memo"`
	CodeAmount            int64  `json:"codeAmount"`
}

// 经纬度
type LatAndLong struct {
	Latitude  string
	Longitude string
}

// 商家信息
type ShopInfo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

// 产品图片
type ProductImg struct {
	Url  string
	Sort int64
}

// 产品信息
type ProductInfo struct {
	ID                   int64         `json:"id"`
	LocationID           int64         `json:"locationId"`
	OnlyName             string        `json:"onlyName"`
	ProductTitle         string        `json:"productTitle"`
	Title                string        `json:"title"`
	ShareText            string        `json:"shareText"`
	FaceImg              string        `json:"faceImg"`
	Address              string        `json:"address"`
	Tel                  string        `json:"tel"`
	BeginTime            int64         `json:"beginTime"`
	EndTime              int64         `json:"endTime"`
	ValidBeginDate       int64         `json:"validBeginDate"`
	ValidEndDate         int64         `json:"validEndDate"`
	SingleMin            int64         `json:"singleMin"`
	SingleMax            int64         `json:"singleMax"`
	ChannelStock         int64         `json:"channelStock"`
	ChannelSaleAmount    int64         `json:"channelSaleAmount"`
	ItemStock            int64         `json:"itemStock"`
	BookingType          int64         `json:"bookingType"`
	BookingBeginDate     int64         `json:"bookingBeginDate"`
	BookingShowAddress   int64         `json:"bookingShowAddress"`
	OrderShowIdCard      int64         `json:"orderShowIdCard"`
	OrderShowDate        int64         `json:"orderShowDate"`
	BookingText          string        `json:"bookingText"`
	Attention            string        `json:"attention"`
	SoldOutTime          int64         `json:"soldOutTime"`
	IsSoldOut            int64         `json:"isSoldOut"`
	City                 string        `json:"city"`
	CityCode             string        `json:"cityCode"`
	Latitude             float64       `json:"latitude"`
	Longitude            float64       `json:"longitude"`
	BookingShowPostTime  int64         `json:"bookingShowPostTime"`
	ChannelVisible       interface{}   `json:"channelVisible"`
	PosterURL            string        `json:"posterUrl"`
	Items                []ProductItem `json:"items"`
	LatAndLongList       []LatAndLong  `json:"latAndLongList"`
	Shops                []ShopInfo    `json:"shops"`
	CategoryPath         string        `json:"categoryPath"`
	CategoryName         string        `json:"categoryName"`
	ProductCategoryID    int64         `json:"productCategoryId"`
	IsReFund             int64         `json:"isReFund"`
	LoopImg              []ProductImg  `json:"loopImg"`
	ContractID           int64         `json:"contractId"`
	QualificationsList   []string      `json:"qualificationsList"`
	ReleaseTime          int64         `json:"releaseTime"`
	ChannelMallPosterImg string        `json:"channelMallPosterImg"`
	CodeDelay            int64         `json:"codeDelay"`
	Name                 string        `json:"name"`
	Ecommerce            int64         `json:"ecommerce"`
}

//查询产品列表
type ProductListRes struct {
	Data      []ProductInfo `json:"data"`
	PageCount int64         `json:"pageCount"`
	PageIndex int64         `json:"pageIndex"`
	PageSize  int64         `json:"pageSize"`
	RowCount  int64         `json:"rowCount"`
}

//创建订单&订单回调-重发
type CreateOrderRes struct {
	ChannelOrderID string `json:"channelOrderId"`
	OrderList      []struct {
		Address             string      `json:"address"`
		BookingURL          interface{} `json:"bookingUrl"`
		Code                interface{} `json:"code"`
		CorderID            string      `json:"corderId"`
		Count               int64       `json:"count"`
		CustomerName        string      `json:"customerName"`
		CustomerPhoneNumber string      `json:"customerPhoneNumber"`
		DetailURL           interface{} `json:"detailUrl"`
		ID                  int64       `json:"id"`
		IDCard              string      `json:"idCard"`
		MorderID            string      `json:"morderId"`
		OrderID             string      `json:"orderId"`
		ProductItemID       int64       `json:"productItemId"`
		ProductItemName     string      `json:"productItemName"`
		QrCodeImg           interface{} `json:"qrCodeImg"`
		QrCodeURL           interface{} `json:"qrCodeUrl"`
		Salt                interface{} `json:"salt"`
		SendSms             int64       `json:"sendSms"`
		Status              int64       `json:"status"`
	} `json:"orderList"`
	PurchaseTime      int64  `json:"purchaseTime"`
	ThirdPartyOrderNo string `json:"thirdPartyOrderNo"`
	ValidBeginDate    int64  `json:"validBeginDate"`
	ValidEndDate      int64  `json:"validEndDate"`
}

// 图文详情
type ProductDetail struct {
	HtmlContent string `json:"htmlContent"`
}

//站点信息
type LocationInfo struct {
	City      string `json:"city"`
	CityCode  string `json:"cityCode"`
	FirstWord string `json:"firstWord"`
	ID        int64  `json:"id"`
	Pid       int64  `json:"pid"`
}

// 产品分类
type ProductCategory struct {
	Id        int64             `json:"id"`
	Name      string            `json:"name"`
	ParentId  int64             `json:"parentId"`
	Level     int64             `json:"level"`
	Sort      int64             `json:"sort"`
	ChildList []ProductCategory `json:"childList"`
}

//核销通知
type ConsumCallbackStruct struct {
	CheckTime         string `json:"checkTime"`
	Num               int64  `json:"num"`
	OrderID           string `json:"orderId"`
	ThirdPartyOrderNo string `json:"thirdPartyOrderNo"`
}

//产品上架/下架/售罄通知
type GoodsCallbackStruct struct {
	Items     []int64 `json:"items"`
	ProductID int64   `json:"productId"`
	Type      int64   `json:"type"`
}

//订单退款
type RefundCallbackStruct struct {
	Count         int64 `json:"count"`
	HandlingPrice int64 `json:"handlingPrice"`
	OrderAllPrice int64 `json:"orderAllPrice"`
	OrderIds      []struct {
		OrderID    string      `json:"orderId"`
		RefundDate interface{} `json:"refundDate"`
	} `json:"orderIds"`
	RefundDate        string `json:"refundDate"`
	ThirdPartyOrderNo string `json:"thirdPartyOrderNo"`
}
