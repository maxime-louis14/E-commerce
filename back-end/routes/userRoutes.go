// routes/user_routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/maxime-louis14/Clone-site-Vilebrequin/controllers"
	"github.com/maxime-louis14/Clone-site-Vilebrequin/middleware"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.CorsMiddleware())

	incomingRoutes.POST("/api/users/register", controller.SignUp())
	incomingRoutes.POST("/api/users/login", controller.Login())

	// Utilisez le middleware ParseMultipartFormMiddleware ici
	incomingRoutes.POST("/api/users/avatar", middleware.Authentication(), middleware.ParseMultipartFormMiddleware(), controller.UploadAvatar())
	incomingRoutes.GET("/api/users/avatar", middleware.Authentication(), controller.GetAvatarImage())
}
