package dumovie

import (
	"bank-activity/common/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
)

//嘟电影

type _dumovie struct {
	AppKey     string
	PrivateKey string
	RequestUrl string
}

func NewSdk(appKey string, privateKey string, requestUrl string) (*_dumovie, error) {
	//sdk, _ := dumovie.NewSdk("shenque", "4c51ef97e80e6363f023e0b3ec15c12e", "http://testopen.mayiufu.com/route")
	//测试地址：http://testopen.mayiufu.com/route
	//线上地址：https://open.dumovie.com/route
	return &_dumovie{
		AppKey:     appKey,
		PrivateKey: privateKey,
		RequestUrl: requestUrl,
	}, nil
}

//请求
func (obj *_dumovie) Request(dataMap map[string]interface{}) (string, error) {
	//组织请求参数
	reqMp := make(map[string]interface{})
	reqMp["appkey"] = obj.AppKey
	reqMp["timestamp"] = utils.DateNowFormatStr()
	reqMp["v"] = "1.0"
	//reqMp["method"] = "dumovie.merchant.account.info"
	for k, v := range dataMap {
		reqMp[k] = v
	}
	var pList = make([]string, 0, 0)
	for key, value := range reqMp {
		if key != "sign" {
			pList = append(pList, fmt.Sprintf("%v=%v", key, value))
		}
	}
	sort.Strings(pList)
	var signStr = strings.Join(pList, "&")
	fmt.Println("dumovie请求signStr：", signStr)
	reqMp["sign"] = utils.Md5(fmt.Sprintf("%v%v", signStr, obj.PrivateKey))
	res, err := utils.HttpPost(obj.RequestUrl, utils.Map2QueryString(reqMp))
	fmt.Println("请求结果:", res, "请求地址：", obj.RequestUrl, "请求数据：", utils.Map2Json(reqMp))
	if err != nil {
		return "", err
	}
	return res, err
}

func (obj *_dumovie) CheckSign(params map[string]string) error {
	var pList []string
	var sign string
	for k, v := range params {
		if k == "sign" {
			sign = v
		} else if v == "" {
			continue
		} else {
			pList = append(pList, fmt.Sprintf("%v=%v", k, v))
		}
	}
	if sign == "" {
		log.Println("DuMovieCheckSignIsNull", params)
		return errors.New("sign为空")
	}

	sort.Strings(pList)
	signStr := strings.ToUpper(utils.Md5(fmt.Sprintf("%v%v", strings.Join(pList, "&"), obj.PrivateKey)))
	if sign != signStr {
		log.Println("DuMovieCheckSignErr", params, signStr)
		return errors.New("签名错误")
	}

	return nil
}

//城市列表
func (obj *_dumovie) CityList() (*CityList, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.city.list"
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var cityList CityList
	err = json.Unmarshal([]byte(res), &cityList)
	if err != nil {
		return nil, err
	}
	return &cityList, nil
}

//获取购票城市区县列表
func (obj *_dumovie) CountyList(citycode int64) (*CountyList, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.county.list"
	param["citycode"] = citycode
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var countyList CountyList
	err = json.Unmarshal([]byte(res), &countyList)
	if err != nil {
		return nil, err
	}
	return &countyList, nil
}

//获取购票城市热映电影列表
func (obj *_dumovie) MovieHotList(citycode int64) (*MovieHotList, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.movie.hot.list"
	param["citycode"] = citycode
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var movieHotList MovieHotList
	err = json.Unmarshal([]byte(res), &movieHotList)
	if err != nil {
		return nil, err
	}
	return &movieHotList, nil
}

//获取即将上映电影列表
func (obj *_dumovie) MovieFutureList(citycode int64) (*MovieFutureList, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.movie.future.list"
	param["citycode"] = citycode
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var movieFutureList MovieFutureList
	err = json.Unmarshal([]byte(res), &movieFutureList)
	if err != nil {
		return nil, err
	}
	return &movieFutureList, nil
}

//影片详细信息
func (obj *_dumovie) MovieDetail(citycode int64, movieid int64) (*MovieDetail, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.movie.detail"
	param["movieid"] = movieid // 1378993
	param["citycode"] = citycode
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var movieDetail MovieDetail
	err = json.Unmarshal([]byte(res), &movieDetail)
	if err != nil {
		return nil, err
	}
	return &movieDetail, nil
}

//获取电影影院列表
func (obj *_dumovie) CinemaList(citycode int64, pageNo int64, per int64) (*CinemaList, error) {
	if per == 0 {
		per = 20
	}
	param := make(map[string]interface{})
	param["method"] = "dumovie.cinema.list"
	param["citycode"] = citycode
	param["page_no"] = pageNo // 查询分页的页数,从0开始
	param["per"] = per        //  	查询分页的数量，默认20， 最大100
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var cinemaList CinemaList
	err = json.Unmarshal([]byte(res), &cinemaList)
	if err != nil {
		return nil, err
	}
	return &cinemaList, nil
}

//获取电影影院列表
func (obj *_dumovie) MoviePlayCinemaList(citycode int64, movieid int64, playdate string) (*MoviePlayCinemaList, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.movie.play.cinema.list"
	param["movieid"] = movieid
	param["citycode"] = citycode
	param["playdate"] = playdate // 放映日期
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var moviePlayCinemaList MoviePlayCinemaList
	err = json.Unmarshal([]byte(res), &moviePlayCinemaList)
	if err != nil {
		return nil, err
	}
	return &moviePlayCinemaList, nil
}

//获取电影影院列表
func (obj *_dumovie) CinemaDetail(cinemaid int64) (*CinemaDetail, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.cinema.detail"
	param["cinemaid"] = cinemaid
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var cinemaDetail CinemaDetail
	err = json.Unmarshal([]byte(res), &cinemaDetail)
	if err != nil {
		return nil, err
	}
	return &cinemaDetail, nil
}

//获取电影影院列表
func (obj *_dumovie) CinemaHallList(cinemaid int64) (*CinemaHallList, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.cinema.hall.list"
	param["cinemaid"] = cinemaid
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var cinemaHallList CinemaHallList
	err = json.Unmarshal([]byte(res), &cinemaHallList)
	if err != nil {
		return nil, err
	}
	return &cinemaHallList, nil
}

//影院排片列表-V2
func (obj *_dumovie) OpiListV2(cinemaid int64) (*OpiListV2, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.opi.list.v2"
	param["cinemaid"] = cinemaid
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	//utils.FilePutContents("11.txt",res,0777)
	var opiListV2 OpiListV2
	err = json.Unmarshal([]byte(res), &opiListV2)
	if err != nil {
		return nil, err
	}
	return &opiListV2, nil
}

//影院排片详情-V2
func (obj *_dumovie) OpiDetailV2(opiid string) (*OpiDetailV2, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.opi.detail.v2"
	param["opiid"] = opiid
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var opiDetailV2 OpiDetailV2
	err = json.Unmarshal([]byte(res), &opiDetailV2)
	if err != nil {
		return nil, err
	}
	return &opiDetailV2, nil
}

//影院排片详情-V2
func (obj *_dumovie) OpiSeat(opiid string) (*OpiSeat, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.opi.seat"
	param["opiid"] = opiid
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var opiSeat OpiSeat
	err = json.Unmarshal([]byte(res), &opiSeat)
	if err != nil {
		return nil, err
	}
	return &opiSeat, nil
}

//创建订单-V2
func (obj *_dumovie) OrderCreateV2(req OrderCreateV2Req) (*OrderCreateV2, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.order.create.v2"
	param["opiid"] = req.Opiid                         //场次ID
	param["seatids"] = req.Seatids                     // 	座位ID.多个座位以 |分隔
	param["seatnames"] = req.Seatnames                 // 	座位名称.多个座位以 |分隔,该字段必须使用以下进行转码：java.net.URLEncoder.encode(seatnames,”UTF-8”);
	param["mobile"] = req.Mobile                       // 	用户购票手机号
	param["totalSettleAmount"] = req.TotalSettleAmount //
	// 1、总的订单结算金额（渠道通知出票需要扣除渠道的金额)
	//2、渠道传递的总的订单结算金额与平台计算金额不一致，则下单失败, 接口返回错误编码: 6999
	//3、如果渠道对场次价格自行进行了调整，请保存好原始场次价格
	//4、导致双方计算的金额不一致原因：渠道缓存了排期，缓存期间影院调整了价格
	//如果渠道接口不传递该金额，则双方对账按照平台的订单金额结算，由于渠道接口没有传递该金额，由此造成的财务问题，由渠道自己承担
	param["pushUrl"] = req.PushUrl                     // 	订单出票成功或者订单退款，平台会推送给渠道订单状态变化
	param["supportChangeSeat"] = req.SupportChangeSeat // 	是否接收换座  Y支持，N不支持

	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var orderCreateV2 OrderCreateV2
	err = json.Unmarshal([]byte(res), &orderCreateV2)
	if err != nil {
		return nil, err
	}
	return &orderCreateV2, nil
}

//订单详情--V2
func (obj *_dumovie) OrderDetailV2(tradeno string) (*OrderDetailV2, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.order.detail.v2"
	param["tradeno"] = tradeno //订单号
	res, err := obj.Request(param)
	log.Println("DuMovieOrderDetailV2", res, err)
	if err != nil {
		return nil, err
	}
	var orderDetailV2 OrderDetailV2
	err = json.Unmarshal([]byte(res), &orderDetailV2)
	if err != nil {
		return nil, err
	}
	return &orderDetailV2, nil
}

//获取爆米花商品列表
func (obj *_dumovie) MovieCinemaBmhList(cinemaid int64) (*MovieCinemaBmhList, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.movie.cinema.bmh.list"
	param["cinemaid"] = cinemaid //影院id
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var movieCinemaBmhList MovieCinemaBmhList
	err = json.Unmarshal([]byte(res), &movieCinemaBmhList)
	if err != nil {
		return nil, err
	}
	return &movieCinemaBmhList, nil
}

//第三方平台与嘟电影做联名登录
func (obj *_dumovie) GetAuthCode(mobile string) (*GetAuthCodeResponse, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.member.get.authcode"
	param["mobile"] = mobile
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var response GetAuthCodeResponse
	err = json.Unmarshal([]byte(res), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//渠道绑定观影券
func (obj *_dumovie) CouponBind(mobile string, cardPass string) (*CouponBindResponse, error) {
	param := make(map[string]interface{})
	param["method"] = "dumovie.member.coupon.bind"
	param["mobile"] = mobile
	param["cardpass"] = cardPass
	res, err := obj.Request(param)
	if err != nil {
		return nil, err
	}
	var response CouponBindResponse
	err = json.Unmarshal([]byte(res), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
