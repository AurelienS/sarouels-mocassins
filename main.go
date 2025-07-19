package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"sarouels-mocassins/handlers"
	"sarouels-mocassins/middleware"
	"sarouels-mocassins/models"
)

func main() {
	// Configuration de la base de données
	db, err := gorm.Open(sqlite.Open("database/app.db"), &gorm.Config{
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
	adminHandler := handlers.NewAdminHandler(db)

	// Configuration du routeur Gin
	r := gin.Default()

	// Configuration CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
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

	// Routes de l'API
	api := r.Group("/api")
	{
		// Routes du jeu (pas d'authentification requise)
		game := api.Group("/game")
		{
			game.GET("/statement/random", gameHandler.GetRandomStatement)
			game.POST("/vote", gameHandler.SubmitVote)
			game.GET("/stats", gameHandler.GetStats)
			game.GET("/results", gameHandler.GetResults)
		}

		// Routes admin (avec authentification)
		admin := api.Group("/admin")
		admin.Use(middleware.AdminAuth())
		{
			admin.GET("/statements", adminHandler.GetAllStatements)
			admin.POST("/statements", adminHandler.AddStatement)
			admin.PUT("/statements/:id", adminHandler.UpdateStatement)
			admin.DELETE("/statements/:id", adminHandler.DeleteStatement)
		}
	}

	// Route par défaut pour servir l'application frontend
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Sarouels & Mocassins",
		})
	})

	// Démarrage du serveur
	log.Println("Serveur démarré sur http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}