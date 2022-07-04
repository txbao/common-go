package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"mime/multipart"
	"path"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//上传文件
var SvcOss = &_oss{}

type _oss struct {
}

//是否重命名
const IsRename = false

//固定OOS文件路径，为空则按系统规则
const OssPath = "" //南宁光大 partner/std/nngd
const (
	// oss存储
	OssAccessKeyId     = "LTAIfnfVUYnmyNbR"
	OssAccessKeySecret = "Qj4D2VRwJT38sodpqmqqVf68yscFLH"
	OssEndpointNet     = "oss-cn-shenzhen.aliyuncs.com"
	OssEndpointLoc     = "oss-cn-shenzhen-internal.aliyuncs.com"
	OssBucket          = "sqqmall-php"
	OssDomain          = "https://image.sqqmall.com/"
)

//判断切片中是否包含某元素
func (obj *_oss) IsInSlice(slice []string, str string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}

//上传文件
func (obj *_oss) Upload(file *multipart.FileHeader, fileExt []string) (string, error) {
	if len(fileExt) == 0 {
		fileExt = []string{".jpg", ".png", ".gif"}
	}
	ext := strings.ToLower(path.Ext(file.Filename))
	if !obj.IsInSlice(fileExt, ext) {
		return "", errors.New("文件格式不支持")
	}
	fileHandle, err := file.Open()
	if err != nil {
		fmt.Println("打开文件错误：", err)
		return "", err
	}

	defer fileHandle.Close()

	// 获取上传文件字节流
	fileByte, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		fmt.Println("获取上传文件字节流：", err)
		return "", err
	}

	//去除.号
	if ext != "" {
		ext = ext[1:len(ext)]
	}
	fileName := file.Filename
	if IsRename {
		fileName = ""
	}
	return obj.UploadFileByte(fileName, ext, fileByte, "dev")
}

//上传文件流
func (obj *_oss) UploadFileByte(filename string, fileExt string, fileByte []byte, env string) (string, error) {
	if env == "" {
		env = "dev"
	}
	// oss配置
	client, err := oss.New(OssEndpointNet, OssAccessKeyId, OssAccessKeySecret)
	if err != nil {
		fmt.Println("oss初始化错误:", err)
		return "", err
	}

	bucket, err := client.Bucket(OssBucket)
	if err != nil {
		fmt.Println("oss-Bucket初始化错误:", err)
		return "", err
	}

	random := obj.CreateRandomString(8)

	//上传阿里云路径
	year, month, day := DateYmdInts()
	//fileExt := strings.ToLower(path.Ext(filename))
	//if fileExt != "" {
	//	fileExt = fileExt[1:len(fileExt)]
	//}
	folderName := fileExt + "/" + IntToString(year) + IntToString(month) + IntToString(day)
	//文件重命名
	w := md5.New()
	if filename == "" {
		filename = fmt.Sprintf("%x", w.Sum(nil)) + "." + fileExt // strings.ToLower(path.Ext(filename))
	}
	yunFileTmpPath := env + "/" + folderName + "/" + random + "/" + filename
	//如果常量配置的话 就直接则常量
	if OssPath != "" {
		yunFileTmpPath = OssPath + "/" + filename
	}

	fmt.Println("文件路径:", yunFileTmpPath)
	err = bucket.PutObject(yunFileTmpPath, bytes.NewReader(fileByte))
	if err != nil {
		return "", err
	}

	imgUrl := OssDomain + yunFileTmpPath

	return imgUrl, nil
}

func (obj *_oss) CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

//localPath 本地文件路径
//fileName 上传后的文件名
func FileToOSS(localPath string, fileName string, env string) (error, string) {
	if env == "" {
		env = "dev"
	}
	client, err := oss.New(OssEndpointNet, OssAccessKeyId, OssAccessKeySecret)
	if err != nil {
		return errors.New("oss初始化错误：err " + err.Error()), ""
	}

	bucket, err := client.Bucket(OssBucket)
	if err != nil {
		fmt.Println("oss-Bucket初始化错误:", err)
		return errors.New("oss-Bucket初始化错误：err " + err.Error()), ""
	}

	//上传阿里云路径
	year, month, day := DateYmdInts()
	folderName := IntToString(year) + IntToString(month) + IntToString(day)
	yunFileTmpPath := path.Join(env + "/" + folderName + "/" + fileName)
	fmt.Println("文件路径:", yunFileTmpPath)

	//test/20200617/100050/GoLang (1).docx
	err = bucket.PutObjectFromFile(yunFileTmpPath, localPath)
	if err != nil {
		fmt.Println("上传阿里云错误:", err)
		return errors.New("上传阿里云错误：err " + err.Error()), ""
	}

	return nil, OssDomain + yunFileTmpPath
}
