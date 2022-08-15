package handler

import (
	inviteService "catalyst-token/services"
	"net/http"

	// models "catalyst-token/models/invite-token-models"

	"github.com/gin-gonic/gin"
)

type SwaggerAPIResponse struct {
	Message string `json:"message"`
	Data string `json:"data"`
}

type TokenForm struct {
	Token string `json:"token" binding:"required"`
}

type inv_handler struct {
	service inviteService.InviteService
}

func InviteRegister(srv inviteService.InviteService) *inv_handler {
	return &inv_handler{service: srv}
}

// Gettokens     godoc
// @Summary      Get tokens list
// @Description  Responds with the list of all tokens as JSON.
// @Tags         tokens
// @Produce      json
// @Success      200  {object} SwaggerAPIResponse
// @Failure      401
// @Failure      400
// @Router       /api/v1/invite-token/validate [POST]
func (h inv_handler) ValidateToken(c *gin.Context) {

	token := TokenForm{}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tk := token.Token
	val := h.service.ValidateToken(tk)
	if val {
		c.JSON(200, gin.H{
			"message": "Token Valid",
			"data": tk,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Token Invalid",
			"data": tk,
		})
	}
}

// Gettokens     godoc
// @Summary      Get tokens list
// @Description  Responds with the list of all tokens as JSON.
// @Tags         tokens
// @Produce      json
// @Success      200  {object} SwaggerAPIResponse
// @Failure      401  
// @Failure      400
// @Router       /api/v1/invite-token [get]
func (h inv_handler) ListTokens(c *gin.Context) {

	val := h.service.ListToken()

	c.JSON(200, gin.H{
		"message": "List of All Tokens",
		"data":    val,
	})
}

// Gettokens     godoc
// @Summary      Create a new token
// @Description  Responds with the Successful token inside dict
// @Tags         tokens
// @Produce      json
// @Success      200  {object} models.InviteToken
// @Failure      401 
// @Failure      400 
// @Failure      500 {object} SwaggerAPIResponse
// @Router       /api/v1/invite-token [POST]
func (h inv_handler) GenerateToken(c *gin.Context) {

	val, str := h.service.GenerateToken()
	if str == "" {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully Created Invite Token",
			"data":    val,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invite Token Creation Failed",
			"data": str,
		})
	}

}

// Gettokens     godoc
// @Summary      Create a new token
// @Description  Responds with the status of revoking Token
// @Tags         tokens
// @Produce      json
// @Success      200  {object} SwaggerAPIResponse
// @Failure      401 
// @Failure      400 
// @Router       /api/v1/invite-token [PUT]
func (h inv_handler) RevokeToken(c *gin.Context) {

	token := TokenForm{}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	val, error_message := h.service.RevokeToken(token.Token)
	if val {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Revoked the token",
			"data" : token.Token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cannot Revoked the token",
			"data" : error_message,
		})
	}
}

// Gettokens     godoc
// @Summary      Create a new token
// @Description  Responds with the Deletion of token
// @Tags         tokens
// @Produce      json
// @Success      200  {object} SwaggerAPIResponse
// @Failure      401  {object} SwaggerAPIResponse
// @Failure      400  {object} SwaggerAPIResponse
// @Failure      500  {object} SwaggerAPIResponse
// @Router       /api/v1/invite-token [DELETE]
func (h inv_handler) DeleteToken(c *gin.Context) {
	var service inviteService.InviteService

	token := TokenForm{}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	val,error_message := service.DeleteToken(token.Token)
	if val {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Deleted Token",
			"data" : token.Token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cannot Delet token",
			"data" : error_message,
		})
	}
}
