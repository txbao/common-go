package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func RemoveOmitemptyTag(filename string) {
	//filename := "./order/order.pb.go"
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ReadFile err: %v\n", err)
		return
	}
	data := strings.ReplaceAll(string(fileData), ",omitempty", "")
	fileData = []byte(data)
	err = ioutil.WriteFile(filename, fileData, 0644)
	if err != nil {
		fmt.Printf("WriteFile err: %v\n", err)
		return
	}
	fmt.Println("RemoveOmitemptyTag successfully")
}
