package database

import (
	"gin-example/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error

	mysqlConfig := configs.GetMysqlConfig()
	var mysqlDsn = mysqlConfig["MYSQL_DSN"]
	Eloquent, err = gorm.Open("mysql", mysqlDsn)

	if err != nil {
		panic(err)
	}

	if Eloquent.Error != nil {
		panic(Eloquent.Error)
	}
}

