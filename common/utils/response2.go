package utils

/**
txbao 输出json格式
*/

import (
	code2 "bank-activity/common/utils/code"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Response2 struct {
	Success    bool        `json:"success"`
	StatusCode int64       `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Code       string      `json:"code"`
	Msg        string      `json:"msg"`
	Timestamp  int64       `json:"timestamp"`
	Data       interface{} `json:"data"`
}

type ResponsePage struct {
	Success    bool        `json:"success"`
	StatusCode int64       `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Code       string      `json:"code"`
	Msg        string      `json:"msg"`
	Count      int64       `json:"count"`
	TotalPage  int64       `json:"totalPage"`
	Limit      int64       `json:"limit"`
	Page       int64       `json:"paging"`
	Timestamp  int64       `json:"timestamp"`
	Data       interface{} `json:"data"`
}

//返回数据
func ResJSON(ctx *gin.Context, code string, data interface{}, msg string) {
	if msg == "" {
		msg = code2.Map[code]
	}
	rsp := &Response2{
		Success:    true,
		StatusCode: http.StatusOK,
		StatusText: http.StatusText(http.StatusOK),
		Code:       code,
		Msg:        msg,
		Timestamp:  time.Now().UnixNano() / 1000000,
		Data:       data,
	}
	ctx.JSON(http.StatusOK, rsp)
}

//返回数据
func ResHtml(ctx *gin.Context, title string, msg string, returnUrl string) {
	if title == "" {
		title = "页面错误"
	}
	ctx.HTML(200, "err_tips.html", gin.H{
		"title":      title,
		"msg":        msg,
		"return_url": returnUrl,
	})
}

//返回分页数据
func ResPageJSON(ctx *gin.Context, code string, model interface{}, count int64, limitInt int64, page int64) {
	//总页面
	var totalPage int64 = int64(math.Ceil(float64(count) / float64(limitInt)))

	rsp := &ResponsePage{
		Success:    true,
		StatusCode: http.StatusOK,
		StatusText: http.StatusText(http.StatusOK),
		Code:       code,
		Msg:        code2.Map[code],
		TotalPage:  totalPage,
		Limit:      limitInt,
		Page:       page,
		Count:      count,
		Timestamp:  time.Now().UnixNano() / 1000000,
		Data:       model,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func ResPageJSONMap(ctx *gin.Context, code string, model interface{}, count int64, limitInt int64, page int64) interface{} {

	//总页面
	var totalPage int64 = int64(math.Ceil(float64(count) / float64(limitInt)))

	rsp := &ResponsePage{
		Success:    true,
		StatusCode: http.StatusOK,
		StatusText: http.StatusText(http.StatusOK),
		Code:       code,
		Msg:        code2.Map[code],
		TotalPage:  totalPage,
		Limit:      limitInt,
		Page:       page,
		Count:      count,
		Timestamp:  time.Now().UnixNano() / 1000000,
		Data:       model,
	}
	return rsp
}

// 获取分页信息
func GetPageInfo(ctx *gin.Context) (int64, int64, int64) {
	page := "1"
	limit := "10"
	if ctx.Request.Method == "GET" {
		limit = ctx.DefaultQuery("limit", limit)
		page = ctx.DefaultQuery("paging", page)
	} else {
		limit = ctx.DefaultPostForm("limit", limit)
		page = ctx.DefaultPostForm("paging", page)
	}
	//总页数

	limitInt, _ := strconv.ParseInt(limit, 10, 64)
	pageInt, _ := strconv.ParseInt(page, 10, 64)

	offset := (pageInt - 1) * limitInt
	return limitInt, offset, pageInt
}

// 获取分页信息\用于RQ请求
func GetPagingOffset(limit int, page int) int {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	//总页数
	offset := (page - 1) * limit
	return offset
}
