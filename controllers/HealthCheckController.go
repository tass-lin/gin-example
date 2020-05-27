package controllers

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) {
	c.Set("statusCode",200)
}
