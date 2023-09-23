package main

import (
	"os"

	middleware "github.com/maxime-louis14/Clone-site-Vilebrequin/middleware"
	routes "github.com/maxime-louis14/Clone-site-Vilebrequin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	router := gin.Default()

	// Appel de la fonction UserRoutes pour définir les routes liées aux utilisateurs
	routes.UserRoutes(router)
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	// API-2
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// API-1
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
