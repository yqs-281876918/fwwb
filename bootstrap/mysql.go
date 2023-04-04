package bootstarp

import (
	"database/sql"
	"fmt"

	"fwwb/library"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() {
	var err error
	connStr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", "root", "Upc202304!", "tcp", "8.130.22.183", 3306, "fwwb")
	library.MysqlDB, err = sql.Open("mysql", connStr)
	if err != nil {
		panic("init mysql failed:" + err.Error())
	}
	//设置数据库最大连接数
	library.MysqlDB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	library.MysqlDB.SetMaxIdleConns(10)
	//验证连接
	if err := library.MysqlDB.Ping(); err != nil {
		panic("open database fail")
	}
	fmt.Println("init mysql succeed")
}
