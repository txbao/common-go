package main

import (
	"flag"
	"fmt"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/txbao/common-go/utils"
    "github.com/txbao/common-go/errorx"
    "github.com/txbao/common-go/reponse"
    "github.com/txbao/common-go/service/{{.serviceName}}/api/internal/middleware"
    "net/http"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")
var env string

func main() {
    /*
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	*/
	flag.StringVar(&env, "env", "debug", "环境变量参数")
    	if !flag.Parsed() {
    		flag.Parse()
    	}

    	var c config.Config
    	//conf.MustLoad(*configFile, &c)

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
            fmt.Println("配置文件main.yaml错误：",err.Error())
            return
        }
        mainLocalYaml, err := utils.FileGetContents(fmt.Sprintf("%s/main-local.yaml", configPath))
        if err != nil {
            fmt.Println("配置文件main-local.yaml错误：",err.Error())
            return
        }
        commonYaml, err := utils.FileGetContents(fmt.Sprintf("%s/common.yaml", commConfigPath))
        if err != nil {
            fmt.Println("配置文件common.yaml错误：",err.Error())
            return
        }
    	etcYaml := fmt.Sprintf("%s\n%s\n%s", mainYaml, mainLocalYaml, commonYaml)
    	err = conf.LoadConfigFromYamlBytes([]byte(etcYaml), &c)
    	if err != nil {
            fmt.Println("配置文件错误：", err.Error())
            return
        }
    	fmt.Println("env:", env)

    	ctx := svc.NewServiceContext(c)
    	//server := rest.MustNewServer(c.RestConf)
    	server := rest.MustNewServer(c.RestConf,
    		rest.WithNotAllowedHandler(middleware.NewCorsMiddleware().Handler()),
    		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
    			// your custom callback function
    			//jwt鉴权失败处理
    			reponse.Response(w, nil, err,0)
    		}))
    	defer server.Stop()
    	server.Use(middleware.NewCorsMiddleware().Handle)

    	handler.RegisterHandlers(server, ctx)

    	// 自定义错误
    	httpx.SetErrorHandler(func(err error) (int, interface{}) {
    		switch e := err.(type) {
    		case *errorx.CodeError:
    			return http.StatusOK, e.Data()
    		default:
    			return http.StatusInternalServerError, nil
    		}
    	})

    	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    	server.Start()
}
