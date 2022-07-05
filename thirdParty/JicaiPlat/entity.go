package JicaiPlat

//OA客户表
type OaCustomerList struct {
	Code string `json:"code"`
	Data struct {
		Count int64 `json:"count"`
		List  []struct {
			BusinessPeopleName string `json:"businessPeopleName"`
			BusinessPeopleNo   string `json:"businessPeopleNo"`
			CustomerName       string `json:"customerName"`
			CustomerNo         string `json:"customerNo"`
			ID                 int64  `json:"id"`
		} `json:"list"`
		PageIndex int64 `json:"pageIndex"`
		PageSize  int64 `json:"pageSize"`
	} `json:"data"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}

//制券通知
type MakeCouponNotifyResp struct {
	Code string `json:"code"`
	Data struct {
		ProjectNo string `json:"project_no"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type IntegralResponse struct {
	Code string `json:"code"`
	Data struct {
		Count int64 `json:"count"`
		List  []struct {
			Id             int64  `json:"id"`             //集采商品表主键
			ProductPic     string `json:"productPic"`     //商品图片
			SpuName        string `json:"spuName"`        //商品名称
			SpuCode        string `json:"spuCode"`        //商品编码
			MarketTaxPrice string `json:"marketTaxPrice"` //市场含税价
			NetPrice       string `json:"netPrice"`       //净价
			Tax            string `json:"tax"`            //税额
			MinOrder       int64  `json:"minOrder"`       //起订量
			TaxRate        string `json:"taxRate"`        //税率
			SkuId          int64  `json:"skuId"`          //集采skuId
			SkuName        string `json:"skuName"`        //sku名称
			SkuCode        string `json:"skuCode"`        //集采skuCode
			Sale           int64  `json:"sale"`           //销量
			ProductNumber  string `json:"productNumber"`  //商品货号
			SupplySkuId    int64  `json:"supplySkuId"`    //供应链sku_id
			OaCustomerId   int64  `json:"oaCustomerId"`   //oa客户方Id
			ProductType    string `json:"productType"`    //商品类型
			Qty            int64  `json:"qty"`            //库存
			SkuPrice       int64  `json:"skuPrice"`       //单价,单位:分
			ProductContent int64  `json:"productContent"` //商品详情
		} `json:"list"`
		PageIndex int64 `json:"pageIndex"`
		PageSize  int64 `json:"pageSize"`
	} `json:"data"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}

type CreateActivity struct {
	Code      string `json:"code"`
	Data      string `json:"data"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}

type GoodsStock struct {
	Code string `json:"code"`
	Data struct {
		Count int64 `json:"count"`
		List  []struct {
			BusinessSide   string `json:"businessSide"`
			BusinessSideNo string `json:"businessSideNo"`
			CustomerID     int64  `json:"customerId"`
			IsUse          int64  `json:"isUse"`
			IsVoucher      int64  `json:"isVoucher"`
			ItemID         int64  `json:"itemId"`
			ProductPic     string `json:"productPic"`
			Quantity       int64  `json:"quantity"`
			SaleMan        string `json:"saleMan"`
			SaleManNo      string `json:"saleManNo"`
			SerialNo       string `json:"serialNo"`
			SkuCode        string `json:"skuCode"`
			SkuID          int64  `json:"skuId"`
			SkuName        string `json:"skuName"`
			SpuName        string `json:"spuName"`
			SupplySkuID    int64  `json:"supplySkuId"`
		} `json:"list"`
		PageIndex int64 `json:"pageIndex"`
		PageSize  int64 `json:"pageSize"`
	} `json:"data"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
}

type ReturnGoodsStock struct {
	Code string `json:"code"`
	Data struct {
		ProjectNo string `json:"project_no"`
	} `json:"data"`
	Msg string `json:"msg"`
}
