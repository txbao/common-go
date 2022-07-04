package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func IsJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

//json转map
func Json2map(jsonStr string) map[string]interface{} {
	dataMap := make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	decoder.UseNumber()
	if err := decoder.Decode(&dataMap); err != nil {
		fmt.Println("err", err)
		return nil
	}
	return dataMap
}

//json转map
func JsonArr2map(jsonStr string) []map[string]interface{} {
	var dataMap []map[string]interface{}
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	decoder.UseNumber()
	if err := decoder.Decode(&dataMap); err != nil {
		fmt.Println("err", err)
		return nil
	}
	return dataMap
}

//json转map
func Interface2map(jsonStr string) map[string]interface{} {
	dataMap := make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	decoder.UseNumber()
	if err := decoder.Decode(&dataMap); err != nil {
		fmt.Println("err", err)
		return nil
	}
	return dataMap
}

//json转interface
func Json2Interface(jsonStr string, ojb interface{}) interface{} {
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	decoder.UseNumber()
	if err := decoder.Decode(&ojb); err != nil {
		fmt.Println("err", err)
		return nil
	}
	return ojb
}

/**
格式化json到Map
*/
func JsonUnmarshal(jsonStr string) map[string]interface{} {
	var personFromJSON interface{}

	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	decoder.UseNumber()
	decoder.Decode(&personFromJSON)

	return personFromJSON.(map[string]interface{})
}
