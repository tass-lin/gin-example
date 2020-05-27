package server

import (
	"github.com/gin-gonic/gin"
	"gin-example/configs"
	"gin-example/routes"
)


func Run(httpServer *gin.Engine) {

	serverConfig := configs.GetServerConfig()

	// gin release debug test
	gin.SetMode(serverConfig["GIN_MODE"])

	switch ginMode := serverConfig["GIN_MODE"] ; ginMode{
	case "release" :
		httpServer = gin.New()
	default:
		httpServer = gin.Default()
	}

	//if "" != serverConfig["VIEWS_PATTERN"] {
	//	httpServer.LoadHTMLGlob(serverConfig["VIEWS_PATTERN"])
	//}
	routes.RegisterRoutes(httpServer)

	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	err := httpServer.Run(serverAddr)

	if nil != err {
		panic("server run error: " + err.Error())
	}
}

