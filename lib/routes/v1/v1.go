package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/JustKato/example-golang-api/lib/routes/v1/auth"
	"github.com/JustKato/example-golang-api/lib/routes/v1/health"
)

func HandleRoutes(r *gin.RouterGroup) {
	// Bind the Auth routes
	auth.HandleRoutes(r.Group("auth"))
	health.HandleRoutes(r.Group("health"))
}
