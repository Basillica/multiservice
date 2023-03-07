package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authenticate godoc
//
//	@Summary		authenticate as a user
//	@Description	authentication endpoint for the gateway logic
//	@Tags           auth
//	@Accept		    json
//	@Produce		json
//	@Success		200		{string}    "status: authorized"	"success"
//	@Failure		403		{string}	string					"forbidden"
//	@Failure		401		{string}	string					"unauthorized"
//	@Failure		500		{string}	string					"server error"
//	@Router			/authenticate  [get]
func Authenticate(c *gin.Context) {
	// The login is handled as a middleware within the JWT Middleware
	c.JSON(http.StatusOK, gin.H{"status": "authorzed"})
}
