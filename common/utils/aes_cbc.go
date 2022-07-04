package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

//Aes CBC加觖密
type _aescbc struct {
	key string
	iv  string
}

func NewAesCbcSdk(key string, iv string) *_aescbc {
	return &_aescbc{
		key: key,
		iv:  iv,
	}
}

// Aes CBC加密 =================================
func (o *_aescbc) AesCbcEncrypt(text []byte) (string, error) {
	//生成cipher.Block 数据块
	block, err := aes.NewCipher([]byte(o.key))
	if err != nil {
		fmt.Println("错误 -" + err.Error())
		return "", err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := o.pad2(text, blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, []byte(o.iv))
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (o *_aescbc) pad2(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (o *_aescbc) AesCbcDecrypt(text string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
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
	block, _ := aes.NewCipher([]byte(o.key))
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(o.iv))
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回

	return string(o.unpad2(origin_data)), nil
}

func (o *_aescbc) unpad2(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	if length-unpadding < 0 {
		return ciphertext
	}

	return ciphertext[:(length - unpadding)]
}
