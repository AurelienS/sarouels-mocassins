package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"sarouels-mocassins/config"
	"sarouels-mocassins/models"
)

var (
	adminMutex sync.RWMutex
	db         *gorm.DB
)

// SetDB initialise la connexion à la base de données
func SetDB(database *gorm.DB) {
	db = database
}

// Fonctions de base de données
func GetAllStatements() ([]models.Statement, error) {
	var statements []models.Statement
	result := db.Find(&statements)
	return statements, result.Error
}

func GetStatement(id int) (*models.Statement, error) {
	var statement models.Statement
	result := db.First(&statement, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &statement, nil
}

func CreateStatement(statement *models.Statement) error {
	return db.Create(statement).Error
}

func UpdateStatement(statement *models.Statement) error {
	return db.Save(statement).Error
}

func DeleteStatement(id int) error {
	return db.Delete(&models.Statement{}, id).Error
}

// Middleware d'authentification
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Vérifier d'abord le cookie
		_, err := c.Cookie("admin_auth")
		if err == nil {
			c.Next()
			return
		}

		// Si pas de cookie, vérifier le header
		password := c.GetHeader("X-Admin-Password")
		if password == config.AdminPassword {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

// Affichage de la page d'administration
func AdminPageHandler(c *gin.Context) {
	// Vérification du cookie d'authentification
	_, err := c.Cookie("admin_auth")
	if err != nil {
		c.HTML(http.StatusOK, "admin_login.html", nil)
		return
	}

	statements, err := GetAllStatements()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"Statements":    statements,
		"AdminPassword": config.AdminPassword,
	})
}

// Login administrateur
func AdminLoginHandler(c *gin.Context) {
	password := c.PostForm("password")
	if password != config.AdminPassword {
		c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{
			"error": "Mot de passe incorrect",
		})
		return
	}

	c.SetCookie("admin_auth", "true", 3600, "/", "", false, true)
	c.Redirect(http.StatusFound, "/admin")
}

// Logout administrateur
func AdminLogoutHandler(c *gin.Context) {
	c.SetCookie("admin_auth", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/admin")
}

// Création d'un statement
func CreateStatementHandler(c *gin.Context) {
	statement := models.Statement{
		Text:          c.PostForm("text"),
		AIChoice:      c.PostForm("ai_choice"),
		AIExplanation: c.PostForm("ai_explanation"),
	}

	err := CreateStatement(&statement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/admin")
}

// Récupération d'un statement
func GetStatementHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	statement, err := GetStatement(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Statement non trouvé"})
		return
	}

	c.JSON(http.StatusOK, statement)
}

// Mise à jour d'un statement
func UpdateStatementHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var updateData struct {
		Text          string `json:"text"`
		AIChoice      string `json:"ai_choice"`
		AIExplanation string `json:"ai_explanation"`
	}

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statement, err := GetStatement(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Statement non trouvé"})
		return
	}

	statement.Text = updateData.Text
	statement.AIChoice = updateData.AIChoice
	statement.AIExplanation = updateData.AIExplanation

	err = UpdateStatement(statement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statement)
}

// Suppression d'un statement
func DeleteStatementHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	err = DeleteStatement(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Statement supprimé"})
}

// Réinitialisation des votes
func ResetVotesHandler(c *gin.Context) {
	adminMutex.Lock()
	defer adminMutex.Unlock()

	// Supprimer tous les votes
	if err := db.Delete(&models.Vote{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Votes réinitialisés"})
}

// Export des données en JSON
func ExportDataHandler(c *gin.Context) {
	statements, err := GetAllStatements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=statements.json")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, statements)
}