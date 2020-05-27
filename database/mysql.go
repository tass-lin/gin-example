package database

import (
	"gin-example/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var Eloquent *gorm.DB

func init() {
	var err error

	mysqlConfig := configs.GetMysqlConfig()
	var mysqlDsn = mysqlConfig["MYSQL_DSN"]
	Eloquent, err = gorm.Open("mysql", mysqlDsn)

	Eloquent.LogMode(false)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	Eloquent.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	Eloquent.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	Eloquent.DB().SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic(err)
	}

	if Eloquent.Error != nil {
		panic(Eloquent.Error)
	}
}

