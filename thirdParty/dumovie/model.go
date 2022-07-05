package dumovie

//城市列表
type CityList struct {
	Data struct {
		CityList []struct {
			Citycode int64  `json:"citycode"`
			Ishot    bool   `json:"ishot"`
			Name     string `json:"name"`
			Pinyin   string `json:"pinyin"`
		} `json:"cityList"`
	} `json:"data"`
	Success bool `json:"success"`
}

//获取购票城市区县列表
type CountyList struct {
	Data struct {
		CountyList []struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"countyList"`
	} `json:"data"`
	Success bool `json:"success"`
}

//获取购票城市热映电影列表
type MovieHotList struct {
	Data struct {
		MovieList []struct {
			Actors      string `json:"actors"`
			Boughtcount int64  `json:"boughtcount"`
			Buytype     string `json:"buytype"`
			Canbuy      bool   `json:"canbuy"`
			Director    string `json:"director"`
			Duration    string `json:"duration"`
			Englishname string `json:"englishname"`
			ID          int64  `json:"id"`
			Language    string `json:"language"`
			Logo        string `json:"logo"`
			Name        string `json:"name"`
			Playdate    string `json:"playdate"`
			Rating      string `json:"rating"`
			Showmark    string `json:"showmark"`
			State       string `json:"state"`
			Type        string `json:"type"`
		} `json:"movieList"`
	} `json:"data"`
	Success bool `json:"success"`
}

//获取即将上映电影列表
type MovieFutureList struct {
	Data struct {
		MovieList []struct {
			Actors         string `json:"actors"`
			Canbuy         bool   `json:"canbuy"`
			Collectedtimes int64  `json:"collectedtimes"`
			Director       string `json:"director"`
			Edition        string `json:"edition"`
			ID             int64  `json:"id"`
			Logo           string `json:"logo"`
			Name           string `json:"name"`
			Playdate       string `json:"playdate"`
			Rating         string `json:"rating"`
			Showmark       string `json:"showmark"`
			State          string `json:"state"`
			Type           string `json:"type"`
		} `json:"movieList"`
	} `json:"data"`
	Success bool `json:"success"`
}

//影片详细信息
type MovieDetail struct {
	Data struct {
		Movie struct {
			Actors         string `json:"actors"`
			Canbuy         bool   `json:"canbuy"`
			Collectedtimes int64  `json:"collectedtimes"`
			Content        string `json:"content"`
			Director       string `json:"director"`
			Duration       string `json:"duration"`
			Englishname    string `json:"englishname"`
			Headlogo       string `json:"headlogo"`
			ID             int64  `json:"id"`
			Language       string `json:"language"`
			Logo           string `json:"logo"`
			Name           string `json:"name"`
			Playdate       string `json:"playdate"`
			Rating         string `json:"rating"`
			State          string `json:"state"`
			Trailers       string `json:"trailers"`
			Type           string `json:"type"`
		} `json:"movie"`
	} `json:"data"`
	Success bool `json:"success"`
}

//获取电影影院列表
type CinemaList struct {
	Data struct {
		CinemaList []struct {
			Address      string  `json:"address"`
			Citycode     int64   `json:"citycode"`
			Cityname     string  `json:"cityname"`
			Contactphone string  `json:"contactphone"`
			Countycode   int64   `json:"countycode"`
			Countyname   string  `json:"countyname"`
			ID           int64   `json:"id"`
			Lat          float64 `json:"lat"`
			Lon          float64 `json:"lon"`
			Name         string  `json:"name"`
			Status       string  `json:"status"`
		} `json:"cinemaList"`
		Pagination struct {
			CurrentPage int64 `json:"current_page"`
			ItemTotal   int64 `json:"item_total"`
			Next        bool  `json:"next"`
			PageTotal   int64 `json:"page_total"`
		} `json:"pagination"`
	} `json:"data"`
	Success bool `json:"success"`
}

//影片开放购票的影院
type MoviePlayCinemaList struct {
	Data struct {
		Cinemaids string `json:"cinemaids"`
	} `json:"data"`
	Success bool `json:"success"`
}

//影院详细信息
type CinemaDetail struct {
	Data struct {
		Cinema struct {
			Address      string  `json:"address"`
			Citycode     int64   `json:"citycode"`
			Cityname     string  `json:"cityname"`
			Contactphone string  `json:"contactphone"`
			Countycode   int64   `json:"countycode"`
			Countyname   string  `json:"countyname"`
			ID           int64   `json:"id"`
			Lat          float64 `json:"lat"`
			Lon          float64 `json:"lon"`
			Name         string  `json:"name"`
			Status       string  `json:"status"`
		} `json:"cinema"`
	} `json:"data"`
	Success bool `json:"success"`
}

//获取影厅列表
type CinemaHallList struct {
	Data struct {
		HallList []struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"hallList"`
	} `json:"data"`
	Success bool `json:"success"`
}

//影院排片列表-V2 OpiListV2
type OpiListV2 struct {
	Data struct {
		Opilist []struct {
			AreaPrice []struct {
				Area        string  `json:"area"`
				ServiceFee  float64 `json:"serviceFee"`
				SettlePrice float64 `json:"settlePrice"`
				ShowPrice   float64 `json:"showPrice"`
			} `json:"areaPrice"`
			Cinemaid    int64  `json:"cinemaid"`
			Cinemaname  string `json:"cinemaname"`
			Closetime   string `json:"closetime"`
			Edition     string `json:"edition"`
			Hallid      int64  `json:"hallid"`
			Hallname    string `json:"hallname"`
			Maxcanbuy   int64  `json:"maxcanbuy"`
			Movieid     int64  `json:"movieid"`
			Moviename   string `json:"moviename"`
			Opiid       string `json:"opiid"`
			Playtime    string `json:"playtime"`
			Roomname    string `json:"roomname"`
			Servicefee  int64  `json:"servicefee"`
			SettlePrice int64  `json:"settlePrice"`
			ShowPrice   int64  `json:"showPrice"`
		} `json:"opilist"`
	} `json:"data"`
	Success bool `json:"success"`
}

//影院排片详情-V2
type OpiDetailV2 struct {
	Data struct {
		Opi struct {
			AreaPrice []struct {
				Area        string  `json:"area"`
				ServiceFee  float64 `json:"serviceFee"`
				SettlePrice float64 `json:"settlePrice"`
				ShowPrice   float64 `json:"showPrice"`
			} `json:"areaPrice"`
			Cinemaid    int64  `json:"cinemaid"`
			Cinemaname  string `json:"cinemaname"`
			Closetime   string `json:"closetime"`
			Edition     string `json:"edition"`
			Hallid      int64  `json:"hallid"`
			Hallname    string `json:"hallname"`
			Maxcanbuy   int64  `json:"maxcanbuy"`
			Movieid     int64  `json:"movieid"`
			Moviename   string `json:"moviename"`
			Opiid       string `json:"opiid"`
			Playtime    string `json:"playtime"`
			Roomname    string `json:"roomname"`
			Servicefee  int64  `json:"servicefee"`
			SettlePrice int64  `json:"settlePrice"`
			ShowPrice   int64  `json:"showPrice"`
		} `json:"opi"`
	} `json:"data"`
	Success bool `json:"success"`
}

//排片座位图
type OpiSeat struct {
	Data struct {
		Opi struct {
			Maxcanbuy  int64  `json:"maxcanbuy"`
			Maxcolumn  int64  `json:"maxcolumn"`
			Maxrow     int64  `json:"maxrow"`
			Mincolumn  int64  `json:"mincolumn"`
			Minrow     int64  `json:"minrow"`
			Notice     string `json:"notice"`
			Opiid      string `json:"opiid"`
			Regular    bool   `json:"regular"`
			Seatcount  int64  `json:"seatcount"`
			Soldcount  int64  `json:"soldcount"`
			TipMessage string `json:"tip_message"`
		} `json:"opi"`
		SeatList []struct {
			Attach    interface{} `json:"attach"`
			Column    int64       `json:"column"`
			Flag      int64       `json:"flag"`
			Leftpx    interface{} `json:"leftpx"`
			Lovegroup interface{} `json:"lovegroup"`
			Name      string      `json:"name"`
			Opiid     string      `json:"opiid"`
			Row       int64       `json:"row"`
			Rowname   string      `json:"rowname"`
			Seatid    string      `json:"seatid"`
			Status    int64       `json:"status"`
			Toppx     interface{} `json:"toppx"`
		} `json:"seatList"`
	} `json:"data"`
	Success bool `json:"success"`
}

//创建订单-V2请求
type OrderCreateV2Req struct {
	Opiid             string  `json:"opiid"`             //场次ID
	Seatids           string  `json:"seatids"`           // 	座位ID.多个座位以 |分隔
	Seatnames         string  `json:"seatnames"`         //座位名称.多个座位以 |分隔,该字段必须使用以下进行转码：java.net.URLEncoder.encode(seatnames,”UTF-8”);
	Mobile            string  `json:"mobile"`            //用户购票手机号
	TotalSettleAmount float64 `json:"totalSettleAmount"` //总的订单结算金额
	PushUrl           string  `json:"pushUrl"`           //订单出票成功或者订单退款，平台会推送给渠道订单状态变化
	SupportChangeSeat string  `json:"supportChangeSeat"` //是否接收换座  Y支持，N不支持
}

//创建订单-V2请求
type OrderCreateV2 struct {
	Data struct {
		Tradeno string `json:"tradeno"`
	} `json:"data"`
	Error struct {
		Code string `json:"code"`
		Desp string `json:"desp"`
	} `json:"error"`
	Success bool `json:"success"`
}

//订单详情--V2
type OrderDetailV2 struct {
	Data struct {
		Order struct {
			Addtime        string `json:"addtime"`
			Cinemaid       int64  `json:"cinemaid"`
			Cinemaname     string `json:"cinemaname"`
			Citycode       string `json:"citycode"`
			Cityname       string `json:"cityname"`
			Edition        string `json:"edition"`
			Expiretime     string `json:"expiretime"`
			Mobile         string `json:"mobile"`
			Movieid        int64  `json:"movieid"`
			Movielogo      string `json:"movielogo"`
			Moviename      string `json:"moviename"`
			OrderMobile    string `json:"order_mobile"`
			Orderstatus    string `json:"orderstatus"`
			Playtime       string `json:"playtime"`
			Quantity       string `json:"quantity"`
			Roomname       string `json:"roomname"`
			Seats          string `json:"seats"`
			NewSeats       string `json:"newSeats"`
			SeatsPrice     string `json:"seats_price"`
			SettleAmount   string `json:"settleAmount"`
			ShowAmount     string `json:"showAmount"`
			Smscontent     string `json:"smscontent"`
			TicketcodeList [][]struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"ticketcodeList"`
			Tradeno      string `json:"tradeno"`
			RefundReason string `json:"refundReason"`
		} `json:"order"`
	} `json:"data"`
	Success bool `json:"success"`
}

//获取爆米花商品列表 - 测试商户暂未开通
type MovieCinemaBmhList struct {
	Data struct {
		ProductList []struct {
			Description string `json:"description"`
			ID          int64  `json:"id"`
			Logo        string `json:"logo"`
			Name        string `json:"name"`
			Price       int64  `json:"price"`
		} `json:"productList"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetAuthCodeResponse struct {
	Data struct {
		AuthCode string `json:"auth_code"`
	} `json:"data"`
	Success bool `json:"success"`
}

type CouponBindResponse struct {
	Success bool `json:"success"`
}
