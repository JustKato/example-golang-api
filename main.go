package main

import (
	"fmt"

	"github.com/JustKato/example-golang-api/lib/conf"
	v1 "github.com/JustKato/example-golang-api/lib/routes/v1"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load the configuration file
	conf := conf.GetConfig()

	r := gin.Default()

	// Register v1 Routes
	v1.HandleRoutes(r.Group("/v1"))

	// Parse the server-side run
	r.Run(fmt.Sprintf("%s:%v", conf.Server.Host, conf.Server.Port))
}
