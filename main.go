package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"sarouels-mocassins/config"
	"sarouels-mocassins/handlers"
	"sarouels-mocassins/models"
)

func main() {
	// Initialisation de la configuration
	config.Init()

	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Configuration de la base de données
	db, err := gorm.Open(sqlite.Open("db/app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Auto-migration des modèles
	err = db.AutoMigrate(&models.Statement{}, &models.Vote{})
	if err != nil {
		log.Fatal("Erreur lors de la migration de la base de données:", err)
	}

	// Initialisation des handlers
	gameHandler := handlers.NewGameHandler(db)
	handlers.SetDB(db) // Initialisation de la connexion pour le handler admin

	// Configuration du routeur Gin
	r := gin.Default()

	// Configuration CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:          12 * time.Hour,
	}))

	// Middleware de récupération des erreurs
	r.Use(gin.Recovery())

	// Configuration des dossiers statiques
	r.Static("/static", "./static")

	// Configuration du dossier des templates
	r.LoadHTMLGlob("templates/*")

	// Routes des pages
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Sarouels & Mocassins",
		})
	})

	r.GET("/game", gameHandler.RenderGamePage)
	r.GET("/game/:id", gameHandler.RenderGamePage) // Nouvelle route pour les questions spécifiques

	// Routes de l'API
	api := r.Group("/api")
	{
		// Routes du jeu (pas d'authentification requise)
		game := api.Group("/game")
		{
			game.GET("/next-question", gameHandler.RenderNextQuestion)
			game.POST("/vote", gameHandler.SubmitVote)
			game.GET("/vote-results", gameHandler.RenderVoteResults)
			game.GET("/stats", gameHandler.GetStats)
		}
	}

	// Routes d'administration
	admin := r.Group("/admin")
	{
		// Routes publiques
		admin.GET("/", handlers.AdminPageHandler)
		admin.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin_login.html", nil)
		})
		admin.POST("/login", handlers.AdminLoginHandler)
		admin.GET("/logout", handlers.AdminLogoutHandler)

		// Routes protégées
		protected := admin.Group("")
		protected.Use(handlers.AdminAuthMiddleware())
		{
			// CRUD Statements
			protected.POST("/statements", handlers.CreateStatementHandler)
			protected.GET("/statements/:id", handlers.GetStatementHandler)
			protected.PUT("/statements/:id", handlers.UpdateStatementHandler)
			protected.DELETE("/statements/:id", handlers.DeleteStatementHandler)

			// Autres fonctionnalités
			protected.POST("/reset-votes", handlers.ResetVotesHandler)
			protected.GET("/export", handlers.ExportDataHandler)
		}
	}

	// Démarrage du serveur
	log.Println("Serveur démarré sur http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}