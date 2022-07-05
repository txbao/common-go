package authority

import (
	"common-go/common/utils"
	"common-go/common/utils/des"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//{
//"success": true,
//"statusCode": 200,
//"statusText": "OK",
//"code": "0",
//"msg": "ok",
//"timestamp": 1607655759543,
//"data": {
//	"token": "185d831c611d98441a50a049f14948be",
//	"expire": 1607659359
//	}
//}
// 公共返回结构体
type ReturnCommonStruct struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Code       string      `json:"code"`
	Msg        string      `json:"msg"`
	Timestamp  int         `json:"timestamp"`
	Data       interface{} `json:"data"`
}

var SrvAuth = &_auth{
	users: map[int]string{0: ""},
}

type _auth struct {
	users     map[int]string
	usersLock sync.RWMutex
	url       string
}

func NewAuth(url string) _auth {
	return _auth{
		users: map[int]string{0: ""},
		url:   url,
	}
}

// 获取登录用户信息数据
func (o *_auth) GetAuthorityToken(result *ReturnCommonStruct, params map[string]string, body map[string]string) error {
	apiUrl := o.url + "/v1/api/login"
	// 增加token
	//params["access-token"] = token
	client := resty.New()
	res, resErr := client.R().
		SetQueryParams(params).
		SetFormData(body).
		SetResult(result).
		Post(apiUrl)
	fmt.Println("authLoginRes", res)
	if resErr != nil {
		fmt.Println("authLogin请求权限登录接口错误：", resErr)
		return resErr
	}
	return nil
}

// 获取权限系统对用用户菜单数据
func (o *_auth) GetAuthorityUserInfo(result *ReturnCommonStruct, params map[string]string) error {
	apiUrl := o.url + "/v1/api/userList"
	// 增加token
	//params["access-token"] = token
	client := resty.New()
	_, resErr := client.R().
		SetQueryParams(params).
		SetResult(result).
		ForceContentType("application/json").
		Get(apiUrl)

	if resErr != nil {
		fmt.Println("请求权限用户信息接口错误：", resErr)
		return resErr
	}
	return nil
}

// 获取用户信息
func (o *_auth) GetAuthorityUserOne(result *ReturnCommonStruct, token string) error {
	apiUrl := o.url + "/v1/api/userInfo"

	body := map[string]string{
		"token": token,
	}
	jsonStr, _ := json.Marshal(body)
	rq, _ := des.DesCbcEncrypt(jsonStr)
	params := map[string]string{
		"rq": rq,
	}

	client := resty.New()
	_, resErr := client.R().
		SetQueryParams(params).
		SetResult(result).
		ForceContentType("application/json").
		Get(apiUrl)

	if resErr != nil {
		fmt.Println("请求权限用户信息接口错误：", resErr)
		return resErr
	}
	return nil
}

//修改用户密码
func (o *_auth) GetAuthorityUserPassword(result *ReturnCommonStruct, token, password, oldPassword string) error {
	apiUrl := o.url + "/v1/api/userPassword"

	body := map[string]string{
		"token":       token,
		"password":    password,
		"oldPassword": oldPassword,
	}
	jsonStr, _ := json.Marshal(body)
	rq, _ := des.DesCbcEncrypt(jsonStr)
	params := map[string]string{
		"rq": rq,
	}
	fmt.Println(params)
	client := resty.New()
	_, resErr := client.R().
		SetFormData(params).
		SetResult(result).
		Post(apiUrl)

	if resErr != nil {
		fmt.Println("请求权限修改密码接口错误：", resErr)
		return resErr
	}
	return nil
}

func (o *_auth) getUser(id int) (string, bool) {
	o.usersLock.RLock()
	defer o.usersLock.RUnlock()
	name, exist := o.users[id]
	return name, exist
}

func (o *_auth) setUsers(users map[int]string) {
	o.usersLock.Lock()
	defer o.usersLock.Unlock()
	for k, v := range users {
		o.users[k] = v
	}
}

func (o *_auth) GetUsers(ids []int) map[int]string {
	defer utils.RecoverName()

	var idStr []string
	data := make(map[int]string)
	for _, id := range ids {
		if name, exist := o.getUser(id); exist {
			data[id] = name
		} else {
			idStr = append(idStr, strconv.Itoa(id))
		}
	}

	if len(idStr) > 0 {
		res := &ReturnCommonStruct{}
		rq, _ := des.DesCbcEncrypt([]byte(`{"ids":"` + strings.Join(idStr, ",") + `"}`))
		err := o.GetAuthorityUserInfo(res, map[string]string{"rq": rq})
		if err == nil {
			resData := res.Data.(map[string]interface{})
			if len(resData) > 0 {
				users := make(map[int]string)
				var id int
				var name string
				for k, v := range resData {
					id = utils.StringToInt(k)
					name = v.(string)
					users[id] = name
					data[id] = name
				}
				o.setUsers(users)
			}
		}
	}

	return data
}

func (o *_auth) GetUserById(id int) string {
	data := o.GetUsers([]int{id})
	return data[id]
}
