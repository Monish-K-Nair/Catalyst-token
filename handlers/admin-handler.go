package handler

import (
	"net/http"

	utils "catalyst-token/utils"

	adminService "catalyst-token/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterNewToken(c *gin.Context) {

	var service adminService.Service
	body := adminService.InputLogin{}
	err := c.BindJSON(&body);
	if  err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, new_err := service.RegisterNewToken(&body)

	if new_err != "" {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"id": res.ID, "email": res.Email}

	token, tk_error := utils.CreateJWT(data, utils.GodotEnv("JWT_SECRET"))

	if tk_error != nil {
		defer logrus.Error(tk_error.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
