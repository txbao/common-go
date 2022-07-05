package tao

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/txbao/common-go/utils"
	"log"
)

type Api struct {
	Host string
}

func (o *Api) GetBrandList(supplierId int64) ([]*SupplierBrand, error) {
	apiUrl := o.Host + fmt.Sprintf("/sqqmall-product/product/third/getSupplierBrands?supplierId=%d", supplierId)
	resp, err := utils.HttpGet(apiUrl)
	if err != nil {
		log.Println("TaoApiGetBrandListHttpGetErr:", supplierId, err)
		return nil, err
	}
	log.Println("TaoApiGetBrandList:", resp)
	type response struct {
		CommonResponse
		Data []*SupplierBrand `json:"data"`
	}
	res := &response{}
	err = json.Unmarshal([]byte(resp), res)
	if err != nil {
		log.Println("TaoApiGetBrandListUnmarshalErr:", resp, err)
		return nil, err
	} else if res.Code != 200 {
		log.Println("TaoApiGetBrandListResErr:", res)
		return nil, errors.New(res.Msg)
	}
	return res.Data, nil
}

func (o *Api) GetProductList(req ProductListRequest) (*ProductListResponse, error) {
	apiUrl := o.Host + fmt.Sprintf("/sqqmall-product/product/third/list?brandId=%d&supplierId=%d&limit=%d&page=%d", req.BrandId, req.SupplierId, req.Limit, req.Page)
	resp, err := utils.HttpGet(apiUrl)
	if err != nil {
		log.Println("TaoApiGetProductListHttpGetErr:", req, err)
		return nil, err
	}
	log.Println("TaoApiGetProductList:", resp)
	type response struct {
		CommonResponse
		Data ProductListResponse
	}
	res := &response{}
	err = json.Unmarshal([]byte(resp), res)
	if err != nil {
		log.Println("TaoApiGetProductListUnmarshalErr:", resp, err)
		return nil, err
	} else if res.Code != 200 {
		log.Println("TaoApiGetProductListResErr:", res)
		return nil, errors.New(res.Msg)
	}
	return &res.Data, nil
}
