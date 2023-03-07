package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x/multiservice/parsejson/types"
)

// ParseJSON godoc
//
//	@Summary		auth login as a user
//	@Description	login for auth users
//	@Tags           auth
//	@Accept		    json
//	@Produce		json
//	@Success		200		{array}    	types.Person	"success"
//	@Failure		403		{string}	string			"forbidden"
//	@Failure		500		{string}	string			"server error"
//	@Router			/json  [get]
func ParseJSON(c *gin.Context) {
	persons := &types.Person{}
	res, err := persons.FetchUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"persons": res})
}
