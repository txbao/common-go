package coupon

import (
	"common-go/utils"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
)

/**
 * error code 说明.
 * <ul>

 *    <li>-41001: encodingAesKey 非法</li>
 *    <li>-41003: aes 解密失败</li>
 *    <li>-41004: 解密后得到的buffer非法</li>
 *    <li>-41005: base64加密失败</li>
 *    <li>-41016: base64解密失败</li>
 * </ul>
 */
func WXBizDataCrypt(encryptedData string, iv string, sessionKey string) (map[string]interface{}, error) {

	esKey, err1 := utils.Base64Decode(sessionKey) //密钥

	aesIV, err2 := utils.Base64Decode(iv) //偏移量

	aesCipher, err3 := utils.Base64Decode(encryptedData) //密文

	if err1 != nil || err2 != nil || err3 != nil {
		return nil, errors.New("base64解密出错！")
	}

	//AES-128-CBC解密
	cipherBlock, err := aes.NewCipher(esKey)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(cipherBlock, aesIV)

	origData := make([]byte, len(aesCipher))

	blockMode.CryptBlocks(origData, aesCipher)

	//填充PKCS#7
	decrypted := unPad(origData)

	userData := make(map[string]interface{})

	err = json.Unmarshal(decrypted, &userData)

	if err != nil {
		return nil, err
	}
	// map[countryCode:86 phoneNumber:13107211253 purePhoneNumber:13107211253 watermark:map[appid:wxa465dbbc622fc3f5 timestamp:1.606459857e+09]]
	//appid := userData["watermark"].(map[string]interface{})["appid"]
	//if appid != global.Config.Weixin.AppId {
	//	return nil, errors.New("数据水印appid不同")
	//}

	return userData, nil

}

func unPad(s []byte) []byte {
	return s[:(len(s) - int(s[len(s)-1]))]
}

func pkcs7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("wrong encryption parameters")
	} else {
		unPadding := int(origData[length-1])
		return origData[:(length - unPadding)], nil
	}
}
