package routes

import (
	"github.com/gin-gonic/gin"
	"gin-example/controllers"
	"gin-example/middlewares"
)

func RegisterRoutes(router *gin.Engine) {


	router.GET("/healthcheck", controllers.HealthCheck)

	api := router.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.Use(middlewares.ResponseHandlerMiddleware)
			v1.GET("/sample", controllers.SampleGet)
			v1.POST("/sample", controllers.SamplePost)
			v1.PATCH("/sample", controllers.SamplePatch)
			v1.DELETE("/sample", controllers.SampleDelete)


			v1.GET("/mysql", controllers.MysqlGet)
			v1.POST("/mysql", controllers.MysqlPost)
			v1.PATCH("/mysql", controllers.MysqlPatch)
			v1.DELETE("/mysql", controllers.MysqlDelete)
		}

	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": false, "message": "Page not found", "data": nil})
	})

}
