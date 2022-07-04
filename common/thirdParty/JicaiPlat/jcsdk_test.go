package JicaiPlat

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//
	//h,_ := os.Hostname()
	//fmt.Println("AA:",h)

	host := "https://dev-integral-erp-api.sqqmall.com"
	appkey := "yinhangandpay"
	appSecret := "N5D8MAOoKYvjzeyBrJlCOpWeIi%jRCw441KihOWWE8@0oqyPcAYF$Oy0li3Ua%zxy07y28G4O3hV8gPw@Avdikfm3BlVd1!s^EJ"
	sdk := NewSdk(host, appkey, appSecret)
	//resp, err := sdk.MakeCouponNotify(20,2,1)
	//fmt.Println(resp, err)

	resp, err := sdk.OaCustomerList()
	fmt.Println(resp, err)

}
