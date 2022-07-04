package tao

type CommonResponse struct {
	Code    int64  `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	BizCode int64  `json:"bizCode"`
}

type SupplierBrand struct {
	Id                  int64  `json:"id"`
	BrandName           string `json:"brandName"`
	SupplierId          int64  `json:"supplierId"`
	RechangeAccountType int64  `json:"rechangeAccountType"`
	UseStatus           int64  `json:"useStatus"`
	BrandOrder          int64  `json:"brandOrder"`
	LogoUrl             string `json:"logoUrl"`
	BrandDesc           string `json:"brandDesc"`
	Keyword             string `json:"keyword"`
	BrandTypes          string `json:"brandTypes"`
	Remarks             string `json:"remarks"`
	Version             int64  `json:"version"`
	Enable              int64  `json:"enable"`
	CreateTime          string `json:"createTime"`
	UpdateTime          string `json:"updateTime"`
	CreateId            int64  `json:"createId"`
	UpdateId            int64  `json:"updateId"`
}

type ProductListRequest struct {
	BrandId    int64 `json:"brandId"`
	SupplierId int64 `json:"supplierId"`
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
}

type ProductListResponse struct {
	TotalCount int64         `json:"totalCount"`
	PageSize   int64         `json:"pageSize"`
	TotalPage  int64         `json:"totalPage"`
	CurrPage   int64         `json:"currPage"`
	List       []ProductInfo `json:"list"`
}

type ProductInfo struct {
	Id                    int64  `json:"id"`
	ProductId             string `json:"productId"`
	ProductName           string `json:"productName"`
	BandProductName       string `json:"bandProductName"`
	BrandId               int64  `json:"brandId"`
	BrandName             string `json:"brandName"`
	ProductIcon           string `json:"productIcon"`
	ProductPic            string `json:"productPic"`
	SupplierId            int64  `json:"supplierId"`
	SupplierName          string `json:"supplierName"`
	RechangeProductType   string `json:"rechangeProductType"`
	RechangeBrandTypeId   int64  `json:"rechangeBrandTypeId"`
	RechangeBrandTypeName string `json:"rechangeBrandTypeName"`
	OfficialPrice         int64  `json:"officialPrice"`
	Price                 int64  `json:"price"`
	SqPrice               int64  `json:"sqPrice"`
	Stock                 int64  `json:"stock"`
	SellStatus            int64  `json:"sellStatus"`
	ProductStatus         int64  `json:"productStatus"`
	CardType              int64  `json:"cardType"`
	Remarks               string `json:"remarks"`
	Version               int64  `json:"version"`
	Enable                int64  `json:"enable"`
	CreateTime            string `json:"createTime"`
	UpdateTime            string `json:"updateTime"`
	CreateId              int64  `json:"createId"`
	UpdateId              int64  `json:"updateId"`
	ValidityTime          int64  `json:"validityTime"`
	UseNotice             string `json:"useNotice"`
	OtherRemark           string `json:"otherRemark"`
	Disclaimer            string `json:"disclaimer"`
	ProductDesc           string `json:"productDesc"`
	ValidityNum           int64  `json:"validityNum"`
}
