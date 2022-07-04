package proto

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

//生成proto
func Test(t *testing.T) {
	dbType := "mysql"
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "shenque", "Sqqmall1234!@", "10.10.10.159", 3306, "bank_activity")
	pkg := "my_package"
	goPkg := "./my_package"
	table := "goods_activity_spu_detail"
	serviceName := "coupon"

	db, err := sql.Open(dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//s, err := core.GenerateSchema(db, table,nil,serviceName, goPkg, pkg)
	s, err := GenerateSchema(db, table, nil, serviceName, goPkg, pkg, false)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}
