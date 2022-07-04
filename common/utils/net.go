package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
	maxBodyLen      = 8 << 20 // 8MB
	emptyJson       = "{}"
)

/*
有关Http协议GET和POST请求的封装
//ss := utils.HttpGet("https://www.cnblogs.com/mafeng/p/7068837.html")

	var param2 = url.Values{}
	param2.Add("sdeee", "txx")
	param2.Add("timestamp", utils.Int64ToString(time.Now().Unix()))

	ss := utils.HttpPost("http://payment.com/payment/index.html","sdeee=ab3&33=ss&ll=555")
	fmt.Println("结果",ss)

	ssc := utils.HttpPostUrl("http://payment.com/payment/index.html",param2)
	fmt.Println("结果",ssc)

*/

//发送GET请求
//url:请求地址
//response:请求返回的内容
func HttpGet(url string) (response string, err error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	if error != nil {
		return "", error
		panic(error)
	}
	defer resp.Body.Close()

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
			panic(err)
		}
	}

	response = result.String()
	return
}

// POST请求
func HttpPost(url string, data string) (string, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		// handle error
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", err
	}

	return string(body), nil
}

// Http请求
func HttpCURL(urlGateway string, method string, data string, headerParam url.Values) (string, error) {
	if method == "" {
		method = "GET"
	}
	client := &http.Client{}

	var req *http.Request
	var err error
	if method == "GET" {
		req, err = http.NewRequest(method, urlGateway, nil)
	} else {
		req, err = http.NewRequest(method, urlGateway, strings.NewReader(data))
	}
	if err != nil {
		// handle error
		return "", err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cookie", "name=anny")

	if headerParam == nil {
		headerParam = make(url.Values, 0)
	}
	for key := range headerParam {
		var value = strings.TrimSpace(headerParam.Get(key))
		if len(value) > 0 {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("HttpCURL_ERR:", err)
		return "", err
	}

	return string(body), nil
}

//post提交，参数为URL
func HttpPostUrl(url string, param url.Values) (string, error) {
	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		pList = append(pList, key+"="+value)
	}
	sort.Strings(pList)
	var data = strings.Join(pList, "&")
	return HttpPost(url, data)
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求放回的内容
func HttpPost3(url string, data interface{}, contentType string) (content string, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if contentType == "" {
		contentType = "application/json"
		//contentType = "application/x-www-form-urlencoded"
	}
	req.Header.Add("content-type", contentType)
	if err != nil {
		return "", err
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		return "", error
		panic(error)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	content = string(result)
	return
}

/*
func HttpPostJson2(api_url string,jsonStr string) (bool,error,string)  {
	//Golang发送post请求　　 post := "{"待发送"："json"}"
	var jsonstr = []byte(jsonStr) //转换二进制
	buffer:= bytes.NewBuffer(jsonstr)
	request, err := http.NewRequest("POST", api_url, buffer)
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
		return false, err,""
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")  //添加请求头
	client := http.Client{} //创建客户端
	resp, err := client.Do(request.WithContext(context.TO DO())) //发送请求
	if err != nil {
		fmt.Printf("client.Do%v", err)
		return false, err,""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll%v", err)
		return false, err,respBytes.
	}
}
*/

//POST json
func HttpPostJson(url string, jsonStr string) (string, error) {

	jsonStrBtype := []byte(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStrBtype))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//POST请求
func HttpPost2(url string, data string) (string, error) {
	//data = "sdeee=cjb&b=123"
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", err
	}

	return string(body), nil
}

func GetPost(ctx *gin.Context) map[string]interface{} {
	PostMap := make(map[string]interface{})
	for k, v := range ctx.Request.PostForm {
		PostMap[k] = v[0]
		//fmt.Printf("k:%v\n", k)
		//fmt.Printf("v:%v\n", v)
	}
	return PostMap
}
func withJsonBody(r *http.Request) bool {
	return r.ContentLength > 0 && strings.Contains(r.Header.Get(ContentType), ApplicationJson)
}

//获取BODY
func GetPostBody(r *http.Request) ([]byte, error) {
	var reader io.Reader
	if withJsonBody(r) {
		reader = io.LimitReader(r.Body, maxBodyLen)
	} else {
		reader = strings.NewReader(emptyJson)
	}
	return ioutil.ReadAll(reader)

}

func GetGet(ctx *gin.Context) map[string]interface{} {
	GetMap := make(map[string]interface{})
	for k, v := range ctx.Request.URL.Query() {
		GetMap[k] = v[0]
		//fmt.Printf("k:%v\n", k)
		//fmt.Printf("v:%v\n", v)
	}
	return GetMap
}

//获取所有GET + POST数据
func GetRequest(ctx *gin.Context) map[string]interface{} {
	var dataMapRes map[string]interface{}
	fmt.Println("ctx.Request.Method", ctx.Request.Method)
	if ctx.Request.Method == "GET" {
		dataMapRes = GetGet(ctx)
	} else {
		dataMapRes = GetPost(ctx)
	}
	return dataMapRes
}

func HttpGetResty(apiUrl string, params map[string]interface{}, headerParams map[string]interface{}, result interface{}) (*resty.Response, error) {
	client := resty.New()
	respRequest := client.R().
		SetQueryParams(MapInteface2string(params)).
		ForceContentType("application/json")
	for k, v := range headerParams {
		respRequest = respRequest.SetHeader(k, Interface2string(v))
	}
	if result != nil {
		respRequest = respRequest.SetResult(result)
	}
	resp, resErr := respRequest.
		Get(apiUrl)

	return resp, resErr
}

func HttpGetResty2(apiUrl string, params map[string]interface{}) (string, error) {
	client := resty.New()
	resp, resErr := client.R().
		SetQueryParams(MapInteface2string(params)).
		ForceContentType("application/json").
		Get(apiUrl)
	if resErr != nil {
		return "", resErr
	}
	return resp.String(), nil
}

func HttpPostResty(apiUrl string, params map[string]interface{}, headerParams map[string]interface{}, result interface{}) (*resty.Response, error) {
	client := resty.New()
	respRequest := client.R().
		SetBody(params).
		SetHeader("Content-Type", "application/json; charset=UTF-8")
	for k, v := range headerParams {
		respRequest = respRequest.SetHeader(k, Interface2string(v))
	}
	if result != nil {
		respRequest = respRequest.SetResult(result)
	}
	resp, resErr := respRequest.Post(apiUrl)
	fmt.Println("HttpPostResty请求地址：", apiUrl, "参数", Map2Json(params), "结果：", resp)
	return resp, resErr
}
func HttpPostResty2(apiUrl string, params map[string]interface{}) (string, error) {
	client := resty.New()
	resp, resErr := client.R().
		SetBody(params).
		SetHeader("Content-Type", "application/json").
		//SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(apiUrl)

	fmt.Println("resp:", resp, resErr)
	if resp == nil {
		return "", resErr
	}
	//fmt.Println("result==:",result)
	return resp.String(), resErr
}
