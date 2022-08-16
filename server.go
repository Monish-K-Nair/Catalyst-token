package main

import (
	db "catalyst-token/config"
	admin "catalyst-token/controllers/auth-controllers"
	invite "catalyst-token/controllers/invite-token-controllers"
	_ "catalyst-token/docs"
	handler "catalyst-token/handlers"
	auth "catalyst-token/middleware"
	services "catalyst-token/services"
	task "catalyst-token/tasks"
	"time"
	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
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

func main() {
	router := gin.Default()
	db := db.SetupConnection()

	admin_handler := handler.HandlerRegister(services.ServiceRegister(admin.RepositoryRegister(db)))
	tk_handler := handler.InviteRegister(services.InviteServiceRegister(invite.InvRepositoryRegister(db)))

	v1 := router.Group("/api/v1")
	{
		token := v1.Group("/invite-token")
		{
			token.GET("", auth.AdminAuth(), tk_handler.ListTokens)
			token.POST("", auth.AdminAuth(), tk_handler.GenerateToken)
			token.PUT("", auth.AdminAuth(), tk_handler.RevokeToken)
			token.PATCH("", auth.AdminAuth(), tk_handler.RevokeToken)
			token.DELETE("", auth.AdminAuth(), tk_handler.DeleteToken)
			token.POST("/validate", limit.MaxAllowed(20), tk_handler.ValidateToken)
		}
		router.POST("/admin/login", admin_handler.RegisterNewToken)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	cron := gocron.NewScheduler(time.UTC)

	cron.Every(1).Seconds().Do(task.InvalidateToken, db)
	cron.StartAsync()

	router.Run()
}
