package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/maxime-louis14/Clone-site-Vilebrequin/controllers"
	"github.com/maxime-louis14/Clone-site-Vilebrequin/middleware"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.CorsMiddleware())
	
	incomingRoutes.POST("/users/register", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
