# 此目录下是自己生成代码目录

数据库文件目录
```
generate\sql
```
模板目录
```
generate\template
```

# goctl使用
Model生成
```
-- 有缓存
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_activity" -table="*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"
goctl model mysql ddl -src="../../generate/sql/*.sql" -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"

======= 指定前缀model
基础服务
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_activity" -table="base_*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"
用户服务
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_activity" -table="user_*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"
订单服务
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_activity" -table="order_*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"
商品服务
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_activity" -table="goods_*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"
卡券服务
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_activity" -table="coupon_*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"
支付服务
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/pay" -table="gb_*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"

新卡券服务
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_coupon" -table="cp_*"  -dir="./model" -cache -style goZero  --home "../../generate/1.3.8"


-- 无缓存
goctl model mysql datasource -url="shenque:Sqqmall1234!@@tcp(10.10.10.159:3306)/bank_activity" -table="*"  -dir="./model" -style goZero  --home "../../generate/1.3.8"
goctl model mysql ddl -src="../../generate/sql/*.sql" -dir="./model" -style goZero  --home "../../generate/1.3.8"
```

Api生成
```
goctl api go -api admin.api -dir . -style goZero  --home "../../../generate/1.3.8"
```

Rpc生成
```
goctl rpc protoc base.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero --home="../../../generate/1.3.8"
```

goctl升级到1.3.8方法
```
GOPROXY=https://goproxy.cn/,direct 
go install github.com/zeromicro/go-zero/tools/goctl@latest
-- 代码升级不需要操作
//goctl migrate —verbose —version v1.3.8
```