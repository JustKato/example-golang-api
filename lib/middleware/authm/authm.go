package authm

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {

	// Get the authorization token
	token := c.GetHeader("Authorization")

	// Check if the provided token exists
	if token == "" {
		// No token was provided, send out an unathorized error code
		c.AbortWithError(http.StatusUnauthorized, errors.New("empty or missing Authorization header"))
	}

	// Check the validity of the token and make sure that it is authorized.
	if !checkToken(token) {
		// The token is invalid, send out the error as a response
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid authentication token was provided"))
	}

	// Send the request to the next token
	c.Next()

}

func checkToken(token string) bool {
	// TODO: Implement actual logic with this
	return false
}
