Name: {{.serviceName}}.rpc
ListenOn: 127.0.0.1:8080
Nacos:
  Hosts: 10.10.10.157
  Port: 8848
  NamespaceId: public  #默认 public
  ServiceName: base.rpc
  LogDir: "/tmp/nacos/log"
  CacheDir: "/tmp/nacos/cache"
#Etcd:
#  Hosts:
#  - 127.0.0.1:2379
#  Key: {{.serviceName}}.rpc
Mysql:
  DataSource: shenque:Sqqmall1234!@@tcp(10.10.10.159)/github.com/txbao/common-go?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: 10.10.10.157:6379
    Pass:
    Type: node
Logs:
  Level: debug