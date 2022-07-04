Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
Timeout: 6000
MaxBytes: 8388608
Mysql:
  DataSource: shenque:Sqqmall1234!@@tcp(10.10.10.159)/bank_activity?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: 10.10.10.157:6379
    Pass:
    Type: node
BaseRpc:
  #NACOS
  Name: client
  Target: nacos://10.10.10.157:8848/base.rpc?namespaceId=public&timeout=5000s