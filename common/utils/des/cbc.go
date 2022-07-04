package des

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
)

type _cbc struct {
	key []byte
	iv  []byte
}

func NewCbc(key []byte, iv []byte) *_cbc {
	return &_cbc{
		key: key,
		iv:  iv,
	}
}

// DES CBC加密 =================================
func (o *_cbc) DesCbcEncrypt(text []byte) (string, error) {
	//生成cipher.Block 数据块
	block, err := aes.NewCipher(o.key)
	if err != nil {
		log.Println("错误 -" + err.Error())
		return "", err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := o.pad(text, blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, o.iv)
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (o *_cbc) pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (o *_cbc) DesCbcDecrypt(text string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	if text == "" {
		return "", errors.New("加密字符为空！")
	}
	decode_data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher(o.key)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, o.iv)
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回

	return string(o.unpad(origin_data)), nil
}

func (o *_cbc) unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	if length-unpadding < 0 {
		return ciphertext
	}

	return ciphertext[:(length - unpadding)]
}

//获取加密RQ ，用法 DesCbcRq("a=b&c=d")
func (o *_cbc) DesCbcRq(q string) string {
	rqStr, _ := o.DesCbcEncrypt([]byte(q))
	return rqStr
}

// DES CBC加密 End =================================
