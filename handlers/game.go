package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"sarouels-mocassins/models"
)

type GameHandler struct {
	db *gorm.DB
}

// EnrichedStats représente les statistiques enrichies
type EnrichedStats struct {
	models.Stats
	TotalVotes      int     `json:"total_votes"`
	AIChoice        string  `json:"ai_choice"`
	PopularityScore float64 `json:"popularity_score"`   // Écart avec 50/50
	ConsensusLevel  string  `json:"consensus_level"`    // "fort", "modéré", "faible"
	MajorityChoice  string  `json:"majority_choice"`    // "left", "right", ou "tie"
	MatchesAI       bool    `json:"matches_ai"`         // La majorité correspond à l'IA
}

func NewGameHandler(db *gorm.DB) *GameHandler {
	return &GameHandler{
		db: db,
	}
}

// selectIntelligentStatement sélectionne une phrase de manière intelligente
func (h *GameHandler) selectIntelligentStatement() (*models.Statement, error) {
	var statements []models.Statement

	// Récupérer toutes les phrases disponibles
	if err := h.db.Find(&statements).Error; err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des phrases: %w", err)
	}

	if len(statements) == 0 {
		return nil, fmt.Errorf("aucune phrase disponible")
	}

	// Appliquer l'algorithme de sélection intelligente
	return h.applySelectionAlgorithm(statements), nil
}

// applySelectionAlgorithm applique l'algorithme de sélection variée
func (h *GameHandler) applySelectionAlgorithm(statements []models.Statement) *models.Statement {
	if len(statements) == 1 {
		return &statements[0]
	}

	// Calculer les poids pour chaque phrase
	weights := make(map[uint]float64)

	for _, stmt := range statements {
		weight := 1.0

		// 1. Favoriser l'équilibre des choix IA
		leftCount, rightCount := h.countAIChoices()
		if stmt.AIChoice == "left" && leftCount < rightCount {
			weight += 0.4 // Favoriser l'équilibre
		} else if stmt.AIChoice == "right" && rightCount < leftCount {
			weight += 0.4
		}

		// 2. Favoriser les phrases avec moins de votes (pour équilibrer les stats)
		voteCount := h.getStatementVoteCount(stmt.ID)
		if voteCount < 5 {
			weight += 0.6 // Favoriser les phrases peu votées
		} else if voteCount < 20 {
			weight += 0.3
		}

		// 3. Ajouter une composante aléatoire pour la variété
		weight += rand.Float64() * 0.5

		weights[stmt.ID] = weight
	}

	// Sélection pondérée
	return h.weightedSelection(statements, weights)
}

// countAIChoices compte la distribution des choix IA dans toutes les phrases
func (h *GameHandler) countAIChoices() (int, int) {
	var results []struct {
		AIChoice string
		Count    int
	}

	h.db.Model(&models.Statement{}).
		Select("ai_choice, COUNT(*) as count").
		Group("ai_choice").
		Scan(&results)

	leftCount, rightCount := 0, 0
	for _, result := range results {
		if result.AIChoice == "left" {
			leftCount = result.Count
		} else {
			rightCount = result.Count
		}
	}

	return leftCount, rightCount
}

// getStatementVoteCount récupère le nombre de votes pour une phrase
func (h *GameHandler) getStatementVoteCount(statementID uint) int {
	var count int64
	h.db.Model(&models.Vote{}).Where("statement_id = ?", statementID).Count(&count)
	return int(count)
}

// weightedSelection effectue une sélection pondérée
func (h *GameHandler) weightedSelection(statements []models.Statement, weights map[uint]float64) *models.Statement {
	totalWeight := 0.0
	for _, weight := range weights {
		totalWeight += weight
	}

	if totalWeight == 0 {
		return &statements[rand.Intn(len(statements))]
	}

	r := rand.Float64() * totalWeight
	currentWeight := 0.0

	for _, stmt := range statements {
		currentWeight += weights[stmt.ID]
		if r <= currentWeight {
			return &stmt
		}
	}

	return &statements[len(statements)-1]
}

// calculateEnrichedStats calcule des statistiques enrichies
func (h *GameHandler) calculateEnrichedStats(statementID uint) (*EnrichedStats, error) {
	var statement models.Statement
	if err := h.db.First(&statement, statementID).Error; err != nil {
		return nil, fmt.Errorf("phrase non trouvée: %w", err)
	}

	// Calculer les statistiques de base
	var result struct {
		LeftVotes  int `json:"left_votes"`
		RightVotes int `json:"right_votes"`
	}

	err := h.db.Model(&models.Vote{}).
		Select("SUM(CASE WHEN choice = 'left' THEN 1 ELSE 0 END) as left_votes, SUM(CASE WHEN choice = 'right' THEN 1 ELSE 0 END) as right_votes").
		Where("statement_id = ?", statementID).
		Scan(&result).Error

	if err != nil {
		return nil, fmt.Errorf("erreur lors du calcul des statistiques: %w", err)
	}

	totalVotes := result.LeftVotes + result.RightVotes
	leftPercent := float64(0)
	rightPercent := float64(0)

	if totalVotes > 0 {
		leftPercent = float64(result.LeftVotes) / float64(totalVotes) * 100
		rightPercent = float64(result.RightVotes) / float64(totalVotes) * 100
	}

	// Calculs enrichis
	popularityScore := calculatePopularityScore(leftPercent, rightPercent)
	consensusLevel := calculateConsensusLevel(leftPercent, rightPercent)
	majorityChoice := calculateMajorityChoice(leftPercent, rightPercent)
	matchesAI := (majorityChoice == statement.AIChoice) || (majorityChoice == "tie")

	return &EnrichedStats{
		Stats: models.Stats{
			StatementID:  statementID,
			LeftVotes:    result.LeftVotes,
			RightVotes:   result.RightVotes,
			LeftPercent:  leftPercent,
			RightPercent: rightPercent,
		},
		TotalVotes:      totalVotes,
		AIChoice:        statement.AIChoice,
		PopularityScore: popularityScore,
		ConsensusLevel:  consensusLevel,
		MajorityChoice:  majorityChoice,
		MatchesAI:       matchesAI,
	}, nil
}

// calculatePopularityScore calcule un score de popularité (écart avec l'équilibre)
func calculatePopularityScore(leftPercent, rightPercent float64) float64 {
	diff := leftPercent - 50.0
	if diff < 0 {
		diff = -diff
	}
	return diff
}

// calculateConsensusLevel détermine le niveau de consensus
func calculateConsensusLevel(leftPercent, rightPercent float64) string {
	maxPercent := leftPercent
	if rightPercent > leftPercent {
		maxPercent = rightPercent
	}

	if maxPercent >= 75 {
		return "fort"
	} else if maxPercent >= 60 {
		return "modéré"
	} else {
		return "faible"
	}
}

// calculateMajorityChoice détermine le choix majoritaire
func calculateMajorityChoice(leftPercent, rightPercent float64) string {
	if leftPercent > rightPercent+5 { // Marge de 5% pour éviter les égalités trop sensibles
		return "left"
	} else if rightPercent > leftPercent+5 {
		return "right"
	} else {
		return "tie"
	}
}

// validateVoteInput valide les données de vote de manière robuste
func (h *GameHandler) validateVoteInput(input map[string]interface{}) (uint, string, error) {
	// Validation de statement_id
	statementIDRaw, exists := input["statement_id"]
	if !exists {
		return 0, "", fmt.Errorf("statement_id manquant")
	}

	var statementID uint
	switch v := statementIDRaw.(type) {
	case float64:
		statementID = uint(v)
	case int:
		statementID = uint(v)
	case string:
		id, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0, "", fmt.Errorf("statement_id invalide: %w", err)
		}
		statementID = uint(id)
	default:
		return 0, "", fmt.Errorf("statement_id doit être un nombre")
	}

	if statementID == 0 {
		return 0, "", fmt.Errorf("statement_id doit être supérieur à 0")
	}

	// Validation du choix
	choiceRaw, exists := input["choice"]
	if !exists {
		return 0, "", fmt.Errorf("choice manquant")
	}

	choice, ok := choiceRaw.(string)
	if !ok {
		return 0, "", fmt.Errorf("choice doit être une chaîne")
	}

	choice = strings.ToLower(strings.TrimSpace(choice))
	if choice != "left" && choice != "right" {
		return 0, "", fmt.Errorf("choice doit être 'left' ou 'right'")
	}

	return statementID, choice, nil
}

// RenderGamePage affiche la page de jeu
func (h *GameHandler) RenderGamePage(c *gin.Context) {
	log.Println("RenderGamePage appelé")
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "game.html", gin.H{
		"title": "Jouer - Sarouels & Mocassins",
		"questionCount": 1,
	})
}

// RenderNextQuestion affiche la prochaine question avec sélection intelligente
func (h *GameHandler) RenderNextQuestion(c *gin.Context) {
	log.Println("RenderNextQuestion appelé")

	// Sélection intelligente de la prochaine phrase
	statement, err := h.selectIntelligentStatement()
	if err != nil {
		log.Printf("Erreur lors de la sélection de phrase: %v", err)
		if err.Error() == "aucune phrase disponible" {
			c.HTML(http.StatusOK, "question.html", gin.H{
				"empty":   true,
				"message": "Aucune question disponible pour le moment !",
			})
			return
		}
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Erreur lors du chargement de la question",
		})
		return
	}

	log.Printf("Question sélectionnée intelligemment: ID=%d, Text=%s, AIChoice=%s",
		statement.ID, statement.Text, statement.AIChoice)

	c.HTML(http.StatusOK, "question.html", gin.H{
		"statement": statement,
	})
}

// SubmitVote enregistre un vote et affiche les résultats enrichis
func (h *GameHandler) SubmitVote(c *gin.Context) {
	log.Println("SubmitVote appelé")

	// Parser le formulaire manuellement
	if err := c.Request.ParseForm(); err != nil {
		log.Printf("Erreur de parsing du formulaire: %v", err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Données de vote invalides",
		})
		return
	}

	// Récupérer les données du formulaire
	statementIDStr := c.Request.PostFormValue("statement_id")
	choice := c.Request.PostFormValue("choice")

	log.Printf("Données reçues: statement_id=%s, choice=%s", statementIDStr, choice)

	// Vérifier que les données sont présentes
	if statementIDStr == "" || choice == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Données de vote manquantes",
		})
		return
	}

	// Convertir statement_id en uint
	statementID, err := strconv.ParseUint(statementIDStr, 10, 32)
	if err != nil {
		log.Printf("Erreur de conversion de statement_id: %v", err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "ID de statement invalide",
		})
		return
	}

	// Valider le choix
	choice = strings.ToLower(strings.TrimSpace(choice))
	if choice != "left" && choice != "right" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Le choix doit être 'left' ou 'right'",
		})
		return
	}

	// Vérifier que la phrase existe
	var statement models.Statement
	if err := h.db.First(&statement, uint(statementID)).Error; err != nil {
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

	// Créer le vote avec horodatage
	vote := models.Vote{
		StatementID: uint(statementID),
		Choice:      choice,
		CreatedAt:   time.Now(),
	}

	if err := h.db.Create(&vote).Error; err != nil {
		log.Printf("Erreur lors de l'enregistrement du vote: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur lors de l'enregistrement du vote",
		})
		return
	}

	// Calculer les statistiques enrichies
	stats, err := h.calculateEnrichedStats(uint(statementID))
	if err != nil {
		log.Printf("Erreur lors du calcul des statistiques: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur lors du calcul des statistiques",
		})
		return
	}

	log.Printf("Vote enregistré - Statement: %d, Choice: %s, Stats: Left=%d%%, Right=%d%%, Consensus=%s",
		statementID, choice, int(stats.LeftPercent), int(stats.RightPercent), stats.ConsensusLevel)

	// Afficher les résultats enrichis
	c.HTML(http.StatusOK, "results.html", gin.H{
		"statement":  statement,
		"stats":     stats,
		"userChoice": choice,
	})
}

// RenderVoteResults affiche les résultats après un vote (endpoint de compatibilité)
func (h *GameHandler) RenderVoteResults(c *gin.Context) {
	statementIDStr := c.Query("statement_id")
	if statementIDStr == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "ID de la phrase manquant",
		})
		return
	}

	statementID, err := strconv.ParseUint(statementIDStr, 10, 32)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "ID de phrase invalide",
		})
		return
	}

	var statement models.Statement
	if err := h.db.First(&statement, uint(statementID)).Error; err != nil {
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

	stats, err := h.calculateEnrichedStats(uint(statementID))
	if err != nil {
		log.Printf("Erreur lors du calcul des statistiques: %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Erreur lors du calcul des statistiques",
		})
		return
	}

	c.HTML(http.StatusOK, "results.html", gin.H{
		"statement": statement,
		"stats":     stats,
	})
}

// GetStats retourne les statistiques globales enrichies
func (h *GameHandler) GetStats(c *gin.Context) {
	var statements []models.Statement
	if err := h.db.Find(&statements).Error; err != nil {
		log.Printf("Erreur lors de la récupération des phrases: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	var allStats []EnrichedStats
	var totalVotes int
	consensusDistribution := map[string]int{"fort": 0, "modéré": 0, "faible": 0}
	aiAccuracy := 0

	for _, stmt := range statements {
		stats, err := h.calculateEnrichedStats(stmt.ID)
		if err != nil {
			log.Printf("Erreur lors du calcul des stats pour la phrase %d: %v", stmt.ID, err)
			continue
		}

		allStats = append(allStats, *stats)
		totalVotes += stats.TotalVotes
		consensusDistribution[stats.ConsensusLevel]++

		if stats.MatchesAI && stats.TotalVotes > 0 {
			aiAccuracy++
		}
	}

	aiAccuracyPercent := float64(0)
	if len(allStats) > 0 {
		aiAccuracyPercent = float64(aiAccuracy) / float64(len(allStats)) * 100
	}

	c.JSON(http.StatusOK, gin.H{
		"stats":                allStats,
		"summary": gin.H{
			"total_statements":       len(statements),
			"total_votes":           totalVotes,
			"consensus_distribution": consensusDistribution,
			"ai_accuracy_percent":   aiAccuracyPercent,
		},
		"timestamp": time.Now(),
	})
}