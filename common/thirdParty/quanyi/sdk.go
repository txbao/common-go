package quanyi

import (
	"github.com/txbao/common-go/common/utils"
	"encoding/json"
	"fmt"
	"log"
)

const (
	ResCodePass = 7000 // 请求通过
)

type SDK struct {
	AppId     string
	Key       string
	ApiUrl    string
	NotifyUrl string
	AesKey    string
	AesIv     string
}

func (o *SDK) ActivityList(page int64, limit int64, name string, status int64) (int64, int64, int64, int64, interface{}, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	if status == 0 {
		status = 0
	}
	url := o.ApiUrl + "rights/api/queryActList?page=" + utils.Int64ToString(page) + "&limit=" + utils.Int64ToString(limit) + "&title=" + utils.URLEncode(name) + "&status=" + utils.Int64ToString(status)
	fmt.Println("url", url)
	res, err := utils.HttpGet(url)
	fmt.Println("res", res)
	response := &ActivityListResponse{}
	err = json.Unmarshal([]byte(res), response)
	if err != nil {
		log.Println("QuanyiActivityListErr", err, res)
		return 0, 0, 0, 0, nil, err
	}
	fmt.Println("response", response)

	return response.Data.TotalCount, response.Data.TotalPage, response.Data.PageSize, response.Data.CurrPage, response.Data.List, nil
}
