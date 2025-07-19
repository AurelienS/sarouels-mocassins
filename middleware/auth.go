package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultAdminPassword = "admin123" // À changer en production !

// AdminAuth vérifie le mot de passe admin dans le header Authorization
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		if adminPassword == "" {
			adminPassword = defaultAdminPassword
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader != adminPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Accès non autorisé"})
			c.Abort()
			return
		}

		c.Next()
	}
}