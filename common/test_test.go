package common

import (
	"github.com/txbao/common-go/common/utils"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//
	//h,_ := os.Hostname()
	//fmt.Println("AA:",h)
	fmt.Println(utils.Getenv("sss") + "dd")
}
