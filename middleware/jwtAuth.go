package middleware

import (
	"net/http"

	utils "catalyst-token/utils"

	"github.com/gin-gonic/gin"
)

type AccessError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AdminAuth() gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {

		var resp AccessError

		resp.Status = ""

		if c.GetHeader("Authorization") == "" {
			resp.Code = http.StatusForbidden
			resp.Status = "Forbidden"
			resp.Message = "Auth required"
			c.JSON(http.StatusForbidden, resp)
			defer c.AbortWithStatus(http.StatusForbidden)
		}

		token, err := utils.VerifyJWTToken(c, "JWT_SECRET")

		if err != nil {
			resp.Status = "Unathorizated"
			resp.Code = http.StatusUnauthorized
			resp.Message = "Access Denied"
			c.JSON(http.StatusUnauthorized, resp)
			defer c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", token.Claims)
		c.Next()

	})
}
