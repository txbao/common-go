package common_go

import (
	"common-go/utils"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//
	//h,_ := os.Hostname()
	//fmt.Println("AA:",h)
	fmt.Println(utils.Getenv("sss"))
}
