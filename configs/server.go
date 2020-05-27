package configs

import "gin-example/lib"

func GetServerConfig() (serverConfig map[string]string) {
	host := lib.GetEnv("HOST","0.0.0.0")
	port := lib.GetEnv("PORT","9999")
	mode := lib.GetEnv("GIN_MODE","release")

	serverConfig = make(map[string]string)

	serverConfig["HOST"] = host
	serverConfig["PORT"] = port
	serverConfig["GIN_MODE"] = mode             //  release/debug/test
	return
}
