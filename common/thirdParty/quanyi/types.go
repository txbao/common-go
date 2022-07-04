package quanyi

type ActivityListReq struct {
	limit int64 `json:"limit"`
	page  int64 `json:"page"`
}

type Goods struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"picUrl"`
}

type Activity struct {
	Id        int64  `json:"id"`
	Code      string `json:"code"`
	Title     string `json:"title"`
	Status    int64  `json:"state"`
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
}

type ActivityData struct {
	TotalCount int64 `json:"totalCount"`
	PageSize   int64 `json:"pageSize"`
	TotalPage  int64 `json:"totalPage"`
	CurrPage   int64 `json:"currPage"`
	List       []Activity
}

type ActivityListResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
	Data ActivityData
}
