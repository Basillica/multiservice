package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x/multiservice/algorithm/helpers"
	"github.com/x/multiservice/algorithm/types"
)

// Algorithm godoc
//
//	@Summary		auth login as a user
//	@Description	login for auth users
//	@Tags           auth
//	@Accept		    json
//	@Produce		json
//	@Success		200		{array}    types.ArrayRequest	"success"
//	@Failure		403		{string}	string			"forbidden"
//	@Failure		401		{string}	string			"unauthorized"
//	@Failure		500		{string}	string			"server error"
//	@Router			/json  [get]
func Algorithm(c *gin.Context) {
	var req types.ArrayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FormatError(c, err)
		return
	}
	arr := &types.ArrayRequest{
		Array: req.Array,
		Int:   req.Int,
	}
	c.JSON(http.StatusOK, gin.H{"result": arr.SortInts()})
}
