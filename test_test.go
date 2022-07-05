package common_go

import (
	"fmt"
	"github.com/txbao/common-go/utils"
	"testing"
)

func Test(t *testing.T) {
	//
	//h,_ := os.Hostname()
	//fmt.Println("AA:",h)
	fmt.Println(utils.Getenv("sss") + "dd")
}
