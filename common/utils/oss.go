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
	ossAccessKeyId     string
	ossAccessKeySecret string
	ossEndpointNet     string
	ossEndpointLoc     string
	ossBucket          string
	ossDomain          string
}

func NewOss(ossAccessKeyId string, ossAccessKeySecret string, ossEndpointNet string, ossEndpointLoc string, ossBucket string, ossDomain string) *_oss {
	return &_oss{
		ossAccessKeyId:     ossAccessKeyId,
		ossAccessKeySecret: ossAccessKeySecret,
		ossEndpointNet:     ossEndpointNet,
		ossEndpointLoc:     ossEndpointLoc,
		ossBucket:          ossBucket,
		ossDomain:          ossDomain,
	}
}

//是否重命名
const IsRename = false

//固定OOS文件路径，为空则按系统规则
const OssPath = "" //

//判断切片中是否包含某元素
func (o *_oss) IsInSlice(slice []string, str string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}

//上传文件
func (o *_oss) Upload(file *multipart.FileHeader, fileExt []string) (string, error) {
	if len(fileExt) == 0 {
		fileExt = []string{".jpg", ".png", ".gif"}
	}
	ext := strings.ToLower(path.Ext(file.Filename))
	if !o.IsInSlice(fileExt, ext) {
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
	return o.UploadFileByte(fileName, ext, fileByte, "dev")
}

//上传文件流
func (o *_oss) UploadFileByte(filename string, fileExt string, fileByte []byte, env string) (string, error) {
	if env == "" {
		env = "dev"
	}
	// oss配置
	client, err := oss.New(o.ossEndpointNet, o.ossAccessKeyId, o.ossAccessKeySecret)
	if err != nil {
		fmt.Println("oss初始化错误:", err)
		return "", err
	}

	bucket, err := client.Bucket(o.ossBucket)
	if err != nil {
		fmt.Println("oss-Bucket初始化错误:", err)
		return "", err
	}

	random := o.CreateRandomString(8)

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

	imgUrl := o.ossDomain + yunFileTmpPath

	return imgUrl, nil
}

func (o *_oss) CreateRandomString(len int) string {
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
func (o *_oss) FileToOSS(localPath string, fileName string, env string) (error, string) {
	if env == "" {
		env = "dev"
	}
	client, err := oss.New(o.ossEndpointNet, o.ossAccessKeyId, o.ossAccessKeySecret)
	if err != nil {
		return errors.New("oss初始化错误：err " + err.Error()), ""
	}

	bucket, err := client.Bucket(o.ossBucket)
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

	return nil, o.ossDomain + yunFileTmpPath
}
