package qrcode

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/txbao/common-go/utils"
	"image/color"
)

type _qrcode struct {
	key string
}

func NewQrcode() *_qrcode {
	return &_qrcode{
	}
}

func (o *_qrcode) CreateFile(text string) (string, error) {
	fileName := fmt.Sprintf("%v%v", utils.DateUnix64(), ".png")
	err := qrcode.WriteColorFile(text, qrcode.Medium, 256, color.White, color.Black, fileName)
	if err != nil {
		fmt.Println("png, err:", err)
		return "", err
	}
	return fileName, nil
}

func (o *_qrcode) CreateOss(text string) (string, error) {
	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	fileName := fmt.Sprintf("%v%v", utils.DateUnix64(), ".png")
	return utils.SvcOss.UploadFileByte(fileName, "png", png, "dev")
}
