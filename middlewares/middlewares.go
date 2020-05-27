package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gin-example/lib"
	"os"
	"time"
)



func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func ResponseHandlerMiddleware(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		log.WithFields(log.Fields{
			"tracingId": c.Request.Header.Get("X-Request-Id"),
			"ip": c.ClientIP(),
			"datetime": time.Now(),
			"url": c.Request.Host+c.Request.RequestURI,
			"error": c.Errors,
		}).Warn("500 err")
		c.JSON(c.GetInt("statusCode"), lib.Response{false,nil, c.Errors})
		return
	}
	if _, exists := c.Get("data"); exists{
		c.JSON(c.GetInt("statusCode"), lib.Response{true,nil, c.MustGet("data")})
	} else {
		c.JSON(c.GetInt("statusCode"), lib.Response{true,nil, nil})
	}
}
