package handler

import (
	"fmt"
	"net/http"

	utils "catalyst-token/utils"

	adminService "catalyst-token/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterNewToken(c *gin.Context) {

	var service adminService.Service
	body := adminService.InputLogin{}
	fmt.Println(body)
	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, new_err := service.RegisterNewToken(&body)

	if new_err != "" {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"id": res.ID, "email": res.Email}

	token, tk_error := utils.CreateJWT(data, "12345678")

	if tk_error != nil {
		defer logrus.Error(tk_error.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
