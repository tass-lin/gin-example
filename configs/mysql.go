package configs

import "gin-example/lib"

func GetMysqlConfig() map[string]string {

	mysqlDsn := lib.GetEnv("MYSQL_DSN","nil")
	if mysqlDsn == "nil" {
		panic("MYSQL_DSN error")
	}
	mongoDbConfig := make(map[string]string)
	mongoDbConfig["MYSQL_DSN"] = mysqlDsn

	return mongoDbConfig
}

