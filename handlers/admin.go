package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"sarouels-mocassins/models"
)

type AdminHandler struct {
	db *gorm.DB
}

func NewAdminHandler(db *gorm.DB) *AdminHandler {
	return &AdminHandler{db: db}
}

// GetAllStatements liste toutes les phrases
func (h *AdminHandler) GetAllStatements(c *gin.Context) {
	var statements []models.Statement
	if err := h.db.Find(&statements).Error; err != nil {
		log.Printf("Erreur lors de la récupération des phrases: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des phrases"})
		return
	}

	c.JSON(http.StatusOK, statements)
}

// AddStatement ajoute une nouvelle phrase
func (h *AdminHandler) AddStatement(c *gin.Context) {
	var statement models.Statement
	if err := c.ShouldBindJSON(&statement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	// Validation des champs obligatoires
	if statement.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le texte de la phrase est requis"})
		return
	}

	if statement.AIChoice != "left" && statement.AIChoice != "right" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La réponse de l'IA doit être 'left' ou 'right'"})
		return
	}

	if statement.AIExplanation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "L'explication de l'IA est requise"})
		return
	}

	if err := h.db.Create(&statement).Error; err != nil {
		log.Printf("Erreur lors de la création de la phrase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de la phrase"})
		return
	}

	c.JSON(http.StatusCreated, statement)
}

// UpdateStatement modifie une phrase existante
func (h *AdminHandler) UpdateStatement(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID manquant"})
		return
	}

	var statement models.Statement
	if err := h.db.First(&statement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Phrase non trouvée"})
			return
		}
		log.Printf("Erreur lors de la récupération de la phrase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	var updateData models.Statement
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	// Validation des champs obligatoires
	if updateData.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le texte de la phrase est requis"})
		return
	}

	if updateData.AIChoice != "left" && updateData.AIChoice != "right" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La réponse de l'IA doit être 'left' ou 'right'"})
		return
	}

	if updateData.AIExplanation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "L'explication de l'IA est requise"})
		return
	}

	// Mise à jour des champs
	statement.Text = updateData.Text
	statement.AIChoice = updateData.AIChoice
	statement.AIExplanation = updateData.AIExplanation

	if err := h.db.Save(&statement).Error; err != nil {
		log.Printf("Erreur lors de la mise à jour de la phrase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour"})
		return
	}

	c.JSON(http.StatusOK, statement)
}

// DeleteStatement supprime une phrase
func (h *AdminHandler) DeleteStatement(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID manquant"})
		return
	}

	// Vérifier si la phrase existe
	var statement models.Statement
	if err := h.db.First(&statement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Phrase non trouvée"})
			return
		}
		log.Printf("Erreur lors de la récupération de la phrase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	// Supprimer les votes associés
	if err := h.db.Where("statement_id = ?", id).Delete(&models.Vote{}).Error; err != nil {
		log.Printf("Erreur lors de la suppression des votes: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression des votes"})
		return
	}

	// Supprimer la phrase
	if err := h.db.Delete(&statement).Error; err != nil {
		log.Printf("Erreur lors de la suppression de la phrase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Phrase supprimée avec succès"})
}