package handler

import (
	model "catalyst-token/models/invite-token-models"
	inviteService "catalyst-token/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type inv_handler struct {
	service inviteService.InviteService
}

func InviteRegister(srv inviteService.InviteService) *inv_handler {
	return &inv_handler{service: srv}
}

func (h inv_handler) ValidateToken(c *gin.Context) {

	var service inviteService.InviteService

	token := model.InviteToken{}
	err := c.BindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	val := service.ValidateToken(&token)
	if val == true {
		c.JSON(200, gin.H{
			"message": "Token InvalidValid",
		})

	}
	c.JSON(200, gin.H{
		"message": "Token InvalidValid",
	})
}

func (h inv_handler) ListTokens(c *gin.Context) {
	var service inviteService.InviteService
	val, str := service.ListToken()
	fmt.Println(val)
	fmt.Println(str)

	c.JSON(200, gin.H{
		"message": "List of All Tokens",
		"data":    val,
	})
}

func (h inv_handler) GenerateToken(c *gin.Context) {
	var service inviteService.InviteService
	val, str := service.GenerateToken()
	if str == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invite Token Creation Failed",
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully Created Invite Token",
		"data":    val,
	})
}

func (h inv_handler) RevokeToken(c *gin.Context) {
	var service inviteService.InviteService
	token := model.InviteToken{}
	err := c.BindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	val := service.RevokeToken(&token)
	if val != true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot Revoke Token",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Revoked the token",
	})
}

func (h inv_handler) DeleteToken(c *gin.Context) {
	var service inviteService.InviteService

	token := model.InviteToken{}
	err := c.BindJSON(&token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	val := service.DeleteToken(&token)
	if val != true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot Delete Token",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted token",
	})
}
