package utils

import (
	"io/ioutil"
	"os"
)

//读取文本
func ReadFile(filePath string) string {
	// 读取文本文件的内容
	txt, err := ioutil.ReadFile(filePath)
	//defer txt.Close()
	if err != nil {
		panic(err)
	}

	// 将字节流转换为字符串
	return string(txt)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
