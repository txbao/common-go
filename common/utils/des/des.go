package des

/**
des加密 用于前后端加密
txbao
*/

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

type _des struct {
	key []byte
}

func NewDes(key []byte) *_des {
	return &_des{
		key: key,
	}
}

func (o *_des) padding(src []byte, blocksize int) []byte {
	n := len(src)
	padnum := blocksize - n%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	dst := append(src, pad...)
	return dst
}

func (o *_des) unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	dst := src[:n-unpadnum]
	return dst
}

func (o *_des) EncryptDES(src []byte) []byte {
	block, _ := des.NewCipher(o.key)
	src = o.padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, o.key)
	blockmode.CryptBlocks(src, src)
	return src
}

func (o *_des) DecryptDES(src []byte) []byte {
	block, _ := des.NewCipher(o.key)
	blockmode := cipher.NewCBCDecrypter(block, o.key)
	blockmode.CryptBlocks(src, src)
	src = o.unpadding(src)
	return src

}
