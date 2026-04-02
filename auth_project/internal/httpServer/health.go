package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"service":   "Authentication API",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
