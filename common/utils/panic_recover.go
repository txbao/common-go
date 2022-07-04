package utils

import "fmt"

//协程一定要调用，否则panic后就停止服务
func RecoverName() {
	if err := recover(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
