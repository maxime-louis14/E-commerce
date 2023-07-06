package routes

import (
	controller "github.com/maxime-louis14/Clone-site-Vilebrequin/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/register", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
