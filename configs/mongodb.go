package configs

import "gin-example/lib"

func GetMongoDbConfig() map[string]string {

	dbDsn := lib.GetEnv("DB_DSN","nil")
	if dbDsn == "nil" {
		panic("dbDsn error")
	}
	mongoDbConfig := make(map[string]string)
	mongoDbConfig["DB_DSN"] = dbDsn

	return mongoDbConfig
}

