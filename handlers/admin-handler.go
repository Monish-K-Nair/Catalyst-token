package handler

import (
	"net/http"

	utils "catalyst-token/utils"

	adminService "catalyst-token/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service adminService.Service
}

func HandlerRegister(service adminService.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterNewToken(c *gin.Context) {

	body := adminService.InputLogin{}
	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, new_err := h.service.RegisterNewToken(&body)

	if new_err != nil {
		c.AbortWithError(http.StatusBadRequest, new_err)
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
