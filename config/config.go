package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AdminPassword string
)

func Init() {
	// Charger le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Fichier .env non trouvé: %v", err)
	}

	// Récupérer les variables d'environnement avec des valeurs par défaut
	AdminPassword = getEnvWithDefault("ADMIN_PASSWORD", "admin123")
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}