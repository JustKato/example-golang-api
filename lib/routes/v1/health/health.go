package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(r *gin.RouterGroup) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
