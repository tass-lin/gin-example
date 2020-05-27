package main

import (
	"github.com/gin-gonic/gin"
	"gin-example/server"
)

var HttpServer *gin.Engine

func main() {
	server.Run(HttpServer)
}
