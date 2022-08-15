package main

import (
	db "catalyst-token/config"
	handler "catalyst-token/handlers"
	auth "catalyst-token/middleware"

	_ "catalyst-token/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files
// @title          Swagger API for Catalyst Token
// @version        1.0
// @description    This is an API for Catalyst Token.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

func main() {
	router := gin.Default()
	db.SetupConnection()
	v1 := router.Group("/api/v1")
	{
		token := v1.Group("/invite-token")
		{
			token.GET("", auth.AdminAuth(), handler.ListTokens)
			token.POST("", auth.AdminAuth(), handler.GenerateToken)
			token.PUT("", auth.AdminAuth(), handler.RevokeToken)
			token.PATCH("", auth.AdminAuth(), handler.RevokeToken)
			token.DELETE("", auth.AdminAuth(), handler.DeleteToken)
		}
		router.POST("/admin/login", handler.RegisterNewToken)
		router.POST("/validate", handler.ValidateToken)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
