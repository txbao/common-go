package utils

import (
	"fmt"
	"reflect"
)

//结构体获取URL Query
func UrlStruct2Query(structObj interface{}) string {
	var valInfo = reflect.ValueOf(structObj)
	st := reflect.TypeOf(structObj)
	urlStr := ""
	for k := 0; k < st.NumField(); k++ {
		if urlStr != "" {
			urlStr += "&"
		}
		field := st.Field(k)
		//fmt.Println("数据--",field.Tag.Get("json"),field.Name,valInfo.Field(k).String(),field.Type)
		urlStr += fmt.Sprintf("%v=%v", field.Tag.Get("json"), valInfo.Field(k))
	}
	return urlStr
}
