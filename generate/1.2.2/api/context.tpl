package svc

import (
	{{.configImport}}

	"github.com/zeromicro/go-zero/core/stores/redis"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
    "fmt"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
    mysql2 "gorm.io/driver/mysql"
)

type ServiceContext struct {
	Config {{.config}}
	Rds    *redis.Redis
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
    conn := sqlx.NewMysql(c.Mysql.DataSource)

	//启动Gorm支持
	db, err := gorm.Open(mysql2.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "base_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,    // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger:logger.Default.LogMode(logger.Info),
	})

	//如果出错就GameOver了
	if err != nil {
		//panic(err)
		fmt.Println("初始化db出错：",err.Error())
	}
	fmt.Println(conn,db)

	//Redis
	cr := c.CacheRedis[0]
	redisObj := redis.NewRedis(cr.Host, cr.Type, cr.Pass)

	return &ServiceContext{
		Config: c,
		Rds:    redisObj,
		{{.middlewareAssignment}}
	}
}
