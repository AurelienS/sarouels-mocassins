package main

import (
	"log"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)

func main() {
	// Initialisation de la base de données
	// db, err := gorm.Open(sqlite.Open("database/app.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Erreur de connexion à la base de données:", err)
	// }

	// Configuration du routeur Gin
	r := gin.Default()

	// Configuration des dossiers statiques
	r.Static("/static", "./static")

	// Configuration du dossier des templates
	r.LoadHTMLGlob("templates/*")

	// Route de base
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Sarouels & Mocassins",
		})
	})

	// Démarrage du serveur
	log.Println("Serveur démarré sur http://localhost:8080")
	r.Run(":8080")
}