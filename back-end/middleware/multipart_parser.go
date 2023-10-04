// middleware/multipart_parser.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http" // Importer le package "net/http" pour utiliser MaxBytesReader
)

const maxSize = 20 * 1024 * 1024 // Limite la taille du corps de la requête à 20 Mo

// ParseMultipartFormMiddleware est une fonction middleware pour analyser les requêtes multipart/form-data.
func ParseMultipartFormMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Limiter la taille du corps de la requête
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)

		err := c.Request.ParseMultipartForm(maxSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur d'analyse du formulaire multipart"})
			c.Abort()
			return
		}

		c.Next()
	}
}
