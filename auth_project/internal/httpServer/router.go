package httpserver

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", healthCheckHandler)

	return r
}
