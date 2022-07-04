package main

import (
	"bank-activity/common/utils"
    "bank-activity/common/utils/logs"
    "flag"
    "fmt"
    "context"
    "log"
    "time"

	{{.imports}}

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

    "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")
var env string

func main() {
    /*
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.New{{.serviceNew}}Server(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		{{.pkg}}.Register{{.service}}Server(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
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
    	    //删除Omitempty标签
        	utils.RemoveOmitemptyTag("./{{.pkg}}/{{.pkg}}.pb.go")

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
    	//logs.Init(ctx.Config.Logs.Level)
    	srv := server.New{{.serviceNew}}Server(ctx)

    	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
    		{{.pkg}}.Register{{.serviceNew}}Server(grpcServer, srv)
    	})
    	defer s.Stop()
    	// Nacos Start
    	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
    		st := time.Now()
    		resp, err = handler(ctx, req)
    		log.Printf("method: %s time: %v\n", info.FullMethod, time.Since(st))
    		return resp, err
    	}

    	s.AddUnaryInterceptors(interceptor)

    	sc := []constant.ServerConfig{
    		*constant.NewServerConfig(c.Nacos.Hosts, c.Nacos.Port),
    	}

    	cc := &constant.ClientConfig{
    		NamespaceId:         c.Nacos.NamespaceId, // namespace id
    		TimeoutMs:           5000,
    		NotLoadCacheAtStart: true,
    		LogDir:              c.Nacos.LogDir,   //"/tmp/nacos/log",
    		CacheDir:            c.Nacos.CacheDir, //"/tmp/nacos/cache",
    		RotateTime:          "1h",
    		MaxAge:              3,
    		LogLevel:            "debug",
    	}

    	opts := nacos.NewNacosConfig(c.Nacos.ServiceName, c.ListenOn, sc, cc)
    	_ = nacos.RegitserService(opts)
    	// Nacos End

    	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
    	s.Start()
}