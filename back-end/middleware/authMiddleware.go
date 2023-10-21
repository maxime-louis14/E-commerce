package middleware

import (
		
	"net/http"

	helper "github.com/maxime-louis14/Clone-site-Vilebrequin/helpers"

	"github.com/gin-gonic/gin"
)

// Authz valide le jeton et autorise les utilisateurs
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("nom", claims.Nom)
		c.Set("prenom", claims.Prenom)
		c.Set("uid", claims.Uid)

		c.Next()

	}
}
