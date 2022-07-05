package coupon

import (
	"common-go/utils"
	"common-go/utils/logs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhnxin/csvreader"
	"net/http"
)

type WxCouponApi struct {
}

//批次列表
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_4.shtml
func (o *WxCouponApi) StockList(req StockListReq) (*StockList, error) {
	if req.Offset < 0 || req.Limit < 1 || req.StockCreatorMchid == "" {
		logs.ErrorTag("WxCouponApiStockListParamsErr", req)
		return nil, errors.New("params error")
	}
	urlApi := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/stocks?stock_creator_mchid=%v&offset=%d&limit=%d", req.StockCreatorMchid, req.Offset, req.Limit)
	if req.CreateEndTime != "" {
		urlApi += fmt.Sprintf("&create_start_time=%v", req.CreateEndTime)
	}
	if req.CreateEndTime != "" {
		urlApi += fmt.Sprintf("&create_end_time=%v", req.CreateEndTime)
	}
	if req.Status != "" {
		urlApi += fmt.Sprintf("&status=%v", req.Status)
	}

	resJsons := V3Request(req.StockCreatorMchid, urlApi, "GET", "")
	stockList := &StockList{}
	err := json.Unmarshal([]byte(resJsons), stockList)
	if err != nil {
		logs.ErrorTag("WxCouponApiStockListUnmarshalErr", map[string]interface{}{
			"resJsons": resJsons,
			"err":      err,
		})
		return nil, err
	}
	return stockList, nil
}

//立减金发放
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_2.shtml
func (o *WxCouponApi) SendCoupon(req SendCouponReq) (*SendCouponResp, error) {
	if req.StockId == "" || req.StockCreatorMchid == "" || req.Openid == "" || req.Appid == "" || req.OutRequestNo == "" {
		logs.ErrorTag("WxCouponApiSendCouponParamsErr", req)
		return nil, errors.New("params error")
	}

	urlApi := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/users/%s/coupons", req.Openid)
	b, _ := json.Marshal(req)
	resJsons := V3Request(req.StockCreatorMchid, urlApi, http.MethodPost, string(b))
	resp := &SendCouponResp{}
	err := json.Unmarshal([]byte(resJsons), resp)
	if err != nil {
		logs.ErrorTag("WxCouponApiSendCouponUnmarshalErr", map[string]interface{}{
			"resJsons": resJsons,
			"err":      err,
		})
		return nil, err
	}
	return resp, nil
}

//查询代金券批次信息
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_5.shtml
func (o *WxCouponApi) StockInfo(req StockInfoReq) (*StockInfoRes, error) {
	if req.StockId == "" || req.StockCreatorMchid == "" {
		logs.ErrorTag("StockInfoParamsErr", req)
		return nil, errors.New("params error")
	}
	urlApi := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/%v?stock_creator_mchid=%v", req.StockId, req.StockCreatorMchid)

	resJsons := V3Request(req.StockCreatorMchid, urlApi, "GET", "")
	logs.Info("StockInfo查询代金券批次结果", resJsons)

	var stockInfoRes StockInfoRes
	if err := json.Unmarshal([]byte(resJsons), &stockInfoRes); err != nil {
		return nil, err
	}

	return &stockInfoRes, nil
}

//查询优惠ID详情
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_6.shtml
func (o *WxCouponApi) CouponInfo(req CouponInfoReq) (*CouponInfoRes, error) {
	if req.MchID == "" || req.AppId == "" || req.Openid == "" || req.CouponId == "" {
		logs.ErrorTag("StockInfoParamsErr", req)
		return nil, errors.New("params error")
	}
	urlApi := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/users/%v/coupons/%v?appid=%v", req.Openid, req.CouponId, req.AppId)
	resJsons := V3Request(req.MchID, urlApi, "GET", "")
	logs.Info("查询代金券结果CouponInfo", resJsons)
	var couponInfoRes CouponInfoRes
	err := json.Unmarshal([]byte(resJsons), &couponInfoRes)
	if err != nil {
		return nil, err
	}
	return &couponInfoRes, nil
}

//下载批次核销明细API
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_10.shtml
func (o *WxCouponApi) StockUseFlow(req StockUseFlowReq) (*StockUseFlowRes2, error) {
	if req.StockId == "" || req.MchID == "" {
		logs.ErrorTag("StockUseFlowParamsErr", req)
		return nil, errors.New("params error")
	}
	urlApi := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/%v/use-flow", req.StockId)

	resJsons := V3Request(req.MchID, urlApi, "GET", "")
	logs.Info("下载批次核销明细API结果StockUseFlow("+req.StockId+")", urlApi, resJsons)
	var stockUseFlowRes1 StockUseFlowRes1
	if err := json.Unmarshal([]byte(resJsons), &stockUseFlowRes1); err != nil {
		return nil, err
	}
	if stockUseFlowRes1.URL == "" {
		return nil, errors.New("URL数据异常")
	}

	resJsons = V3Request(req.MchID, stockUseFlowRes1.URL, "GET", "")
	fmt.Println("下载批次核销明细API结果CouponUseFlow2("+req.StockId+"):", resJsons)

	//struct slice
	datas := []StockUseFlowResData{}
	_ = csvreader.New().
		WithHeader([]string{"stock_id", "coupon_id", "stock_type", "coupon_amount", "total_amount", "trade_type", "out_trade_no", "consum_time", "consum_mchid", "dev_no", "bank_trade_no", "goods_detail"}).
		UnMarshalBytes([]byte(resJsons), &datas)

	dataLen := 0
	if len(datas) > 0 {
		dataLen = len(datas) - 1
	}
	return &StockUseFlowRes2{
		Count: dataLen,
		Data:  datas,
	}, nil

}

//下载批次退款明细API
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_11.shtml
func (o *WxCouponApi) StockRefundFlow(req StockRefundFlowReq) (*StockRefundFlowRes2, error) {
	if req.StockId == "" || req.MchID == "" {
		logs.ErrorTag("StockRefundFlowParamsErr", req)
		return nil, errors.New("params error")
	}
	urlApi := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/marketing/favor/stocks/%v/refund-flow", req.StockId)

	resJsons := V3Request(req.MchID, urlApi, "GET", "")
	logs.Info("下载批次退款明细API结果StockRefundFlow1_"+req.StockId, resJsons, urlApi)

	var stockRefundFlowRes1 StockRefundFlowRes1
	err := json.Unmarshal([]byte(resJsons), &stockRefundFlowRes1)
	if err != nil {
		return nil, err
	}
	if stockRefundFlowRes1.URL == "" {
		return nil, errors.New("URL数据异常")
	}

	resJsons = V3Request(req.MchID, stockRefundFlowRes1.URL, "GET", "")
	fmt.Println("下载批次退款明细API结果StockRefundFlow2_"+req.StockId+":", resJsons)

	//struct slice
	datas := []StockRefundFlow2DataStruct{}
	_ = csvreader.New().
		WithHeader([]string{"stock_id", "coupon_id", "stock_type", "coupon_amount", "total_amount", "trade_type", "out_trade_no", "consum_time", "consum_mchid", "dev_no", "bank_trade_no", "goods_detail"}).
		UnMarshalBytes([]byte(resJsons), &datas)
	fmt.Println("AAAA", len(datas))

	dataLen := 0
	if len(datas) > 0 {
		dataLen = len(datas) - 1
	}

	return &StockRefundFlowRes2{
		Count: dataLen,
		Data:  datas,
	}, nil
}

//设置消息通知地址API
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_12.shtml
func (o *WxCouponApi) SetCallbackUrl(req SetCallbackUrlReq) (*SetCallbackUrlRes, error) {
	if req.Mchid == "" || req.NotifyUrl == "" {
		logs.ErrorTag("SetCallbackUrlParamsErr", req)
		return nil, errors.New("params error")
	}
	urlApi := "https://api.mch.weixin.qq.com/v3/marketing/favor/callbacks"

	jsonStr, err := utils.StructToJsonStr(req)
	if err != nil {
		return nil, err
	}
	resJsons := V3Request(req.Mchid, urlApi, "POST", string(jsonStr))
	logs.Info("设置消息通知地址APISetCallbackUrl", resJsons, urlApi)

	var setCallbackUrlRes SetCallbackUrlRes
	err = json.Unmarshal([]byte(resJsons), &setCallbackUrlRes)
	if err != nil {
		return nil, err
	}
	return &setCallbackUrlRes, nil
}

//设置消息通知地址API
//https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_1_12.shtml
func (o *WxCouponApi) ConsumeNotify(req ConsumeNotifyReq) (*ConsumeNotifyRes, error) {

	//return &setCallbackUrlRes, nil
	return nil, nil
}
