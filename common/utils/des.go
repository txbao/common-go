package utils

/**
des加密 用于前后端加密
txbao
*/

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	"log"
)

func padding(src []byte, blocksize int) []byte {
	n := len(src)
	padnum := blocksize - n%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	dst := append(src, pad...)
	return dst
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	dst := src[:n-unpadnum]
	return dst
}

func EncryptDES(src []byte) []byte {
	key := []byte("qimiao66")
	block, _ := des.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func DecryptDES(src []byte) []byte {
	key := []byte("qimiao66")
	block, _ := des.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src

}

// DES CBC加密 =================================
var key = []byte("l49ij@Dz0z3X$Gza")
var iv = []byte("FdO!$8M4bbpTmzxw")

func DesCbcEncrypt(text []byte) (string, error) {
	//生成cipher.Block 数据块
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("错误 -" + err.Error())
		return "", err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad(text, blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func DesCbcDecrypt(text string) (string, error) {
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
	block, _ := aes.NewCipher(key)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回

	return string(unpad(origin_data)), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	if length-unpadding < 0 {
		return ciphertext
	}

	return ciphertext[:(length - unpadding)]
}

//获取加密RQ ，用法 DesCbcRq("a=b&c=d")
func DesCbcRq(q string) string {
	rqStr, _ := DesCbcEncrypt([]byte(q))
	return rqStr
}

// DES CBC加密 End =================================
