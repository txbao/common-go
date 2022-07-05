package cfg

import (
	"fmt"
	"github.com/txbao/common-go/utils"
	"os"
)

func GetEtcYaml(env string) (string, error) {
	commConfigPath := utils.GetExcPath()
	var configPath string = ""

	switch env {
	case "debug":
		commConfigPath = "./etc"
		configPath = commConfigPath + "/debug"
	default:
		commConfigPath = fmt.Sprintf("%s/etc", commConfigPath)
		configPath = fmt.Sprintf("%s/%s", commConfigPath, env)
	}

	mainYaml, err := utils.FileGetContents(fmt.Sprintf("%s/main.yaml", configPath))
	if err != nil {
		fmt.Println("配置文件main.yaml错误：", err.Error())
		return "", err
	}
	mainLocalYaml, err := utils.FileGetContents(fmt.Sprintf("%s/main-local.yaml", configPath))
	if err != nil {
		fmt.Println("配置文件main-local.yaml错误：", err.Error())
		return "", err
	}
	commonYaml, err := utils.FileGetContents(fmt.Sprintf("%s/common.yaml", commConfigPath))
	if err != nil {
		fmt.Println("配置文件common.yaml错误：", err.Error())
		return "", err
	}
	etcYaml := fmt.Sprintf("%s\n%s\n%s", mainYaml, mainLocalYaml, commonYaml)

	return replaceStr(etcYaml), nil
}

func replaceStr(etcYaml string) string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "err"
	}
	//替换计算机名称
	etcYaml = utils.StrReplace("${HOST_NAME}", hostname, etcYaml, -1)
	return etcYaml
}
