package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"sarouels-mocassins/models"
)

type GameHandler struct {
	db *gorm.DB
}

func NewGameHandler(db *gorm.DB) *GameHandler {
	return &GameHandler{db: db}
}

// GetRandomStatement retourne une phrase aléatoire
func (h *GameHandler) GetRandomStatement(c *gin.Context) {
	var statement models.Statement
	result := h.db.Order("RANDOM()").First(&statement)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucune phrase disponible"})
			return
		}
		log.Printf("Erreur lors de la récupération d'une phrase aléatoire: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	c.JSON(http.StatusOK, statement)
}

// SubmitVote enregistre un vote utilisateur
func (h *GameHandler) SubmitVote(c *gin.Context) {
	var vote models.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données de vote invalides"})
		return
	}

	if vote.Choice != "left" && vote.Choice != "right" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le choix doit être 'left' ou 'right'"})
		return
	}

	// Vérifier que la phrase existe
	var statement models.Statement
	if err := h.db.First(&statement, vote.StatementID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Phrase non trouvée"})
			return
		}
		log.Printf("Erreur lors de la vérification de la phrase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	if err := h.db.Create(&vote).Error; err != nil {
		log.Printf("Erreur lors de l'enregistrement du vote: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'enregistrement du vote"})
		return
	}

	c.JSON(http.StatusCreated, vote)
}

// GetStats retourne les statistiques globales
func (h *GameHandler) GetStats(c *gin.Context) {
	var stats []models.Stats
	result := h.db.Raw(`
		SELECT
			statement_id,
			SUM(CASE WHEN choice = 'left' THEN 1 ELSE 0 END) as left_votes,
			SUM(CASE WHEN choice = 'right' THEN 1 ELSE 0 END) as right_votes,
			CAST(SUM(CASE WHEN choice = 'left' THEN 1 ELSE 0 END) AS FLOAT) / COUNT(*) * 100 as left_percent,
			CAST(SUM(CASE WHEN choice = 'right' THEN 1 ELSE 0 END) AS FLOAT) / COUNT(*) * 100 as right_percent
		FROM votes
		GROUP BY statement_id`).Scan(&stats)

	if result.Error != nil {
		log.Printf("Erreur lors de la récupération des statistiques: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des statistiques"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetResults retourne les résultats après vote
func (h *GameHandler) GetResults(c *gin.Context) {
	statementID := c.Query("statement_id")
	if statementID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de la phrase manquant"})
		return
	}

	var statement models.Statement
	if err := h.db.First(&statement, statementID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Phrase non trouvée"})
			return
		}
		log.Printf("Erreur lors de la récupération de la phrase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	var stats models.Stats
	result := h.db.Raw(`
		SELECT
			statement_id,
			SUM(CASE WHEN choice = 'left' THEN 1 ELSE 0 END) as left_votes,
			SUM(CASE WHEN choice = 'right' THEN 1 ELSE 0 END) as right_votes,
			CAST(SUM(CASE WHEN choice = 'left' THEN 1 ELSE 0 END) AS FLOAT) / COUNT(*) * 100 as left_percent,
			CAST(SUM(CASE WHEN choice = 'right' THEN 1 ELSE 0 END) AS FLOAT) / COUNT(*) * 100 as right_percent
		FROM votes
		WHERE statement_id = ?
		GROUP BY statement_id`, statementID).Scan(&stats)

	if result.Error != nil {
		log.Printf("Erreur lors de la récupération des statistiques: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statement": statement,
		"stats":     stats,
	})
}