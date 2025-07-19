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
	questionCount int
}

func NewGameHandler(db *gorm.DB) *GameHandler {
	return &GameHandler{
		db: db,
		questionCount: 0,
	}
}

// RenderGamePage affiche la page de jeu
func (h *GameHandler) RenderGamePage(c *gin.Context) {
	log.Println("RenderGamePage appelé")
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "game.html", gin.H{
		"title": "Jouer - Sarouels & Mocassins",
		"questionCount": h.questionCount,
	})
}

// RenderNextQuestion affiche la prochaine question via HTMX
func (h *GameHandler) RenderNextQuestion(c *gin.Context) {
	log.Println("RenderNextQuestion appelé")
	var statement models.Statement
	result := h.db.Order("RANDOM()").First(&statement)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("Aucune question trouvée")
			c.HTML(http.StatusOK, "question.html", gin.H{
				"empty": true,
				"message": "Aucune question disponible pour le moment !",
			})
			return
		}
		log.Printf("Erreur lors de la récupération d'une phrase aléatoire: %v", result.Error)
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Erreur lors du chargement de la question",
		})
		return
	}

	h.questionCount++
	log.Printf("Question chargée: %+v", statement)
	c.HTML(http.StatusOK, "question.html", gin.H{
		"statement": statement,
		"questionCount": h.questionCount,
	})
}

// SubmitVote enregistre un vote utilisateur et affiche les résultats
func (h *GameHandler) SubmitVote(c *gin.Context) {
	var input struct {
		StatementID uint   `json:"statement_id"`
		Choice      string `json:"choice"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Données de vote invalides",
		})
		return
	}

	if input.Choice != "left" && input.Choice != "right" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Le choix doit être 'left' ou 'right'",
		})
		return
	}

	// Vérifier que la phrase existe
	var statement models.Statement
	if err := h.db.First(&statement, input.StatementID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"error": "Phrase non trouvée",
			})
			return
		}
		log.Printf("Erreur lors de la vérification de la phrase: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur serveur",
		})
		return
	}

	// Créer le vote
	vote := models.Vote{
		StatementID: input.StatementID,
		Choice:      input.Choice,
	}

	if err := h.db.Create(&vote).Error; err != nil {
		log.Printf("Erreur lors de l'enregistrement du vote: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur lors de l'enregistrement du vote",
		})
		return
	}

	// Calculer les statistiques
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
		GROUP BY statement_id`, input.StatementID).Scan(&stats)

	if result.Error != nil {
		log.Printf("Erreur lors de la récupération des statistiques: %v", result.Error)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur lors de la récupération des statistiques",
		})
		return
	}

	// Afficher les résultats
	c.HTML(http.StatusOK, "results.html", gin.H{
		"statement": statement,
		"stats":     stats,
	})
}

// RenderVoteResults affiche les résultats après un vote
func (h *GameHandler) RenderVoteResults(c *gin.Context) {
	statementID := c.Query("statement_id")
	if statementID == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "ID de la phrase manquant",
		})
		return
	}

	var statement models.Statement
	if err := h.db.First(&statement, statementID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"error": "Phrase non trouvée",
			})
			return
		}
		log.Printf("Erreur lors de la récupération de la phrase: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur serveur",
		})
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
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur lors de la récupération des statistiques",
		})
		return
	}

	c.HTML(http.StatusOK, "results.html", gin.H{
		"statement": statement,
		"stats":     stats,
	})
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