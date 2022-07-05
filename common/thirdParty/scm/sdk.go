package scm

import (
	"bytes"
	"github.com/txbao/common-go/common/utils"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ResCodeOk           int64 = 200
	ResCodeCreated      int64 = 201
	ResCodeUnauthorized int64 = 401
	ResCodeForbidden    int64 = 403
	ResCodeNotFound     int64 = 404
)

type SDK struct {
	ApiUrl     string
	SecretKey  string
	BusinessId string
	Iv         string
}

func (o *SDK) post(api string, rq string) ([]byte, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	url := o.ApiUrl + api
	req, err := http.NewRequest("POST", url, strings.NewReader(rq))
	if err != nil {
		log.Println("ScmPostNewRequestErr", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("ScmPostClientDoErr", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ScmPostReadAllErr", err)
		return nil, err
	}

	return body, nil
}

func (o *SDK) aesEncrypt(src []byte) (string, error) {
	block, err := aes.NewCipher([]byte(o.SecretKey))
	if err != nil {
		log.Println("ScmAesEncryptNewCipherErr", map[string]interface{}{
			"o":   *o,
			"src": string(src),
			"err": err,
		})
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, []byte(o.Iv))

	content := o.pkcs5Padding(src, blockSize)
	encrypted := make([]byte, len(content))
	blockMode.CryptBlocks(encrypted, content)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (o *SDK) pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func (o *SDK) Test() {
	//res, err := o.SupplierList()
	//log.Println(res, err)
	//res1, err := o.SpuList(3, "", 1, 10)
	//log.Println(res1, err)
	//res2, err := o.SkuList(9)
	//log.Println(res2, err)
	res3, err := o.ApplyPerformanceUrl(ApplyPerformanceUrlRequest{
		ProductDesc:  "desc1",
		ProductName:  "name1",
		ProductPic:   "pic1",
		ActivityName: "act1",
		RequestId:    strconv.FormatInt(time.Now().UnixNano(), 10),
		Data: LaunchPerformanceReq{
			UserId:          "user1",
			TransactionNo:   "TN1",
			ProductId:       1,
			Quantity:        1,
			ActivityId:      "a1",
			TransactionType: 1,
		},
	})
	log.Println(res3, err)
}

func (o *SDK) SupplierList() (*SupplierListResponse, error) {
	rq, _ := json.Marshal(map[string]interface{}{
		"businessId": o.BusinessId,
	})
	b, err := o.post("/mpf-scm-interface/call/query/supplier/list", string(rq))
	log.Println("ScmResSupplierList", string(b), err)
	if err != nil {
		return nil, err
	}
	resp := &SupplierListResponse{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		log.Println("ScmSupplierListJsonUnmarshalErr", map[string]interface{}{
			"b":   string(b),
			"err": err,
		})
		return nil, err
	}
	return resp, nil
}

func (o *SDK) SpuList(supplierId int64, spuName string, current, size int64) (*SpuListResponse, error) {
	rq, _ := json.Marshal(map[string]interface{}{
		"businessId": o.BusinessId,
		"current":    current,
		"size":       size,
		"spuName":    spuName,
		"supplierId": supplierId,
	})
	b, err := o.post("/mpf-scm-interface/call/query/spu/list", string(rq))
	log.Println("ScmResSpuList", string(b), err)
	if err != nil {
		return nil, err
	}
	resp := &SpuListResponse{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		log.Println("ScmSpuListJsonUnmarshalErr", map[string]interface{}{
			"b":   string(b),
			"err": err,
		})
		return nil, err
	}
	return resp, nil
}

func (o *SDK) SkuList(spuId int64) (*SkuListResponse, error) {
	rq, _ := json.Marshal(map[string]interface{}{
		"businessId": o.BusinessId,
		"spuId":      spuId,
	})
	b, err := o.post(fmt.Sprintf("/mpf-scm-interface/call/query/sku/%d", spuId), string(rq))
	log.Println("ScmResSkuList", string(b), err)
	if err != nil {
		return nil, err
	}
	resp := &SkuListResponse{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		log.Println("ScmSkuListJsonUnmarshalErr", map[string]interface{}{
			"b":   string(b),
			"err": err,
		})
		return nil, err
	}
	return resp, nil
}

func (o *SDK) ApplyPerformanceUrl(req ApplyPerformanceUrlRequest) (*ApplyPerformanceUrlResponse, error) {
	src, _ := json.Marshal(req.Data)
	data, err := o.aesEncrypt(src)
	if err != nil {
		return nil, err
	}

	rq, _ := json.Marshal(map[string]interface{}{
		"businessId":   o.BusinessId,
		"data":         data,
		"productDesc":  req.ProductDesc,
		"productName":  req.ProductName,
		"productPic":   req.ProductPic,
		"activityName": req.ActivityName,
		"requestId":    req.RequestId,
	})
	reqJson,_ := utils.StructToJsonStr(req)
	log.Println("ScmReqApplyPerformanceUrlReq", reqJson, string(rq))
	b, err := o.post("/mpf-scm-interface/call/query/apply/performance/url", string(rq))
	log.Println("ScmResApplyPerformanceUrl", string(b), err)
	if err != nil {
		return nil, err
	}
	fmt.Printf("SCMErr 请求供应链接口：%v,参数：%v", string(b), utils.Map2Json(rq))

	resp := &ApplyPerformanceUrlResponse{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		log.Println("ScmApplyPerformanceUrlJsonUnmarshalErr", map[string]interface{}{
			"b":   string(b),
			"err": err,
		})
		return nil, err
	}
	return resp, nil
}
