package main

import (
	"fmt"

	"github.com/JustKato/example-golang-api/lib/conf"
	v1 "github.com/JustKato/example-golang-api/lib/routes/v1"
	"github.com/gin-gonic/gin"
)

// @title           Example Golang API
// @version         1.0
// @description     This is an example golang API setup project.
// @termsOfService  https://justkato.me/

// @contact.name   Kato Twofold
// @contact.url    https://justkato.me
// @contact.email  contact@justkato.me

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.oauth2  OAuth2
func main() {

	// Load the configuration file
	conf := conf.GetConfig(true)

	// Initialize the gin server
	r := gin.Default()

	// Register v1 Routes
	v1.HandleRoutes(r.Group("/v1"))

	// You can add routing here for v2, v3, etc...

	// Parse the server-side run
	r.Run(fmt.Sprintf("%s:%v", conf.Server.Host, conf.Server.Port))
}
