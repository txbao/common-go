package utils

/**
Rq前端传值处理
txbao
*/
import (
	des2 "github.com/txbao/common-go/common/utils/des"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type RqJson struct {
	Unsigned    string `json:"unsigned"`
	ActivityId  string `json:"activity_id"`
	Code        string `json:"code"`
	ShareUserId string `json:"share_user_id"`
}

/**
解密为结构体
*/
func DesCbcDecryStruct(str string, key []byte, iv []byte) RqJson {
	cbdDesSdk := des2.NewCbc(key, iv)
	jsonStr, _ := cbdDesSdk.DesCbcDecrypt(str)
	var des RqJson

	if IsJSONString(jsonStr) {
		// 将字符串反解析为结构体
		json.Unmarshal([]byte(jsonStr), &des)
	} else {
		//a=1&b=2&c=3
		mapStr, _ := url.ParseQuery(jsonStr)
		//fmt.Println(mapStr)
		//RqJson = ss.Get("c")
		des.ActivityId = mapStr.Get("activity_id")
		des.Code = mapStr.Get("code")
		des.ShareUserId = mapStr.Get("share_user_id")
	}
	return des
}

/**
解密为字典
*/
func DesCbcDecryptMap(str string, key []byte, iv []byte) map[string]interface{} {
	cbdDesSdk := des2.NewCbc(key, iv)
	jsonStr, _ := cbdDesSdk.DesCbcDecrypt(str)
	if jsonStr == "" {
		//URL 解码
		str, _ = url.QueryUnescape(str)
		jsonStr, _ = cbdDesSdk.DesCbcDecrypt(str)
	}
	log.Println("[jsonStr]", jsonStr)
	//fmt.Println("IsJSONString",IsJSONString(jsonStr))
	//fmt.Println("IsJSON",IsJSON(jsonStr))
	if IsJSON(jsonStr) {
		//fmt.Println("JSON",jsonStr)
		// 将字符串反解析为字典
		var dict map[string]interface{}
		//json.Unmarshal([]byte(jsonStr), &dict)
		dict = Json2map(jsonStr)
		return dict
	} else {
		//fmt.Println("不是JSON",jsonStr)
		dict2, _ := url.ParseQuery(jsonStr)
		//fmt.Println("dict2",dict2)
		dict := make(map[string]interface{})
		for k, v := range dict2 {
			dict[k] = fmt.Sprint(v[0])
		}
		fmt.Println("dict", dict)
		return dict
	}
}
