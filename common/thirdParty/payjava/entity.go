package payjava

//中行解密
type BocDecryptMessage struct {
	Code string `json:"code"`
	Data struct {
		Base64msg string `json:"base64msg"`
	} `json:"data"`
	Msg string `json:"msg"`
}

//中行解密客户信息
type BocDecryptMessageBase64msg struct {
	BranchID       string `json:"branchId"`
	CifNumber      string `json:"cifNumber"`
	CreateDate     string `json:"createDate"`
	CustomerID     string `json:"customerId"`
	CustomerName   string `json:"customerName"`
	Gender         string `json:"gender"`
	Ibknum         string `json:"ibknum"`
	IdentityNumber string `json:"identityNumber"`
	IdentityType   string `json:"identityType"`
	Mobile         string `json:"mobile"`
	OrgID          string `json:"orgId"`
}
