package handler

import (
	inviteService "catalyst-token/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenForm struct {
	Token string `json:"token" binding:"required"`
}

type inv_handler struct {
	service inviteService.InviteService
}

func InviteRegister(srv inviteService.InviteService) *inv_handler {
	return &inv_handler{service: srv}
}

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
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Token Invalid",
		})
	}
}

func (h inv_handler) ListTokens(c *gin.Context) {

	val, str := h.service.ListToken()
	fmt.Println(val)
	fmt.Println(str)

	c.JSON(200, gin.H{
		"message": "List of All Tokens",
		"data":    val,
	})
}

func (h inv_handler) GenerateToken(c *gin.Context) {

	val, str := h.service.GenerateToken()
	if str {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully Created Invite Token",
			"data":    val,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invite Token Creation Failed",
		})
	}

}

func (h inv_handler) RevokeToken(c *gin.Context) {

	token := TokenForm{}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	val := h.service.RevokeToken(token.Token)
	if val {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot Revoke Token",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Revoked the token",
		})
	}
}

func (h inv_handler) DeleteToken(c *gin.Context) {
	var service inviteService.InviteService

	token := TokenForm{}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	val := service.DeleteToken(token.Token)
	if val {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot Delete Token",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted token",
		})
	}
}
