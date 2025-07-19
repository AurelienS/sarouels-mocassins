package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Statement struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Text          string `json:"text" gorm:"not null"`
	AIChoice      string `json:"ai_choice" gorm:"not null"`
	AIExplanation string `json:"ai_explanation" gorm:"not null"`
}

type Vote struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	StatementID uint   `json:"statement_id"`
	Choice      string `json:"choice" gorm:"not null"`
}

func main() {
	// Configuration de la base de données
	db, err := gorm.Open(sqlite.Open("../database/app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}

	// Auto-migration des modèles
	err = db.AutoMigrate(&Statement{}, &Vote{})
	if err != nil {
		log.Fatal("Erreur lors de la migration de la base de données:", err)
	}

	// Supprimer toutes les données existantes
	db.Exec("DELETE FROM statements")
	db.Exec("DELETE FROM votes")

	// Questions de test
	statements := []Statement{
		{
			Text:          "Porter un sarouel",
			AIChoice:      "left",
			AIExplanation: "Le sarouel est souvent associé à une mode alternative et à un style de vie bohème, caractéristiques traditionnellement associées à la gauche.",
		},
		{
			Text:          "Porter des mocassins",
			AIChoice:      "right",
			AIExplanation: "Les mocassins sont souvent associés à un style vestimentaire classique et traditionnel, plus typique de la droite.",
		},
		{
			Text:          "Écouter du reggae",
			AIChoice:      "left",
			AIExplanation: "Le reggae est souvent associé à des valeurs de paix, de justice sociale et de résistance, traditionnellement liées à la gauche politique.",
		},
		{
			Text:          "Acheter une BMW",
			AIChoice:      "right",
			AIExplanation: "BMW est une marque de luxe allemande, souvent associée au statut social et à la réussite financière, valeurs plus traditionnellement liées à la droite.",
		},
		{
			Text:          "Faire du vélo en ville",
			AIChoice:      "left",
			AIExplanation: "Le vélo en ville est associé à la protection de l'environnement, aux transports durables et à la critique de la voiture, valeurs écologistes souvent liées à la gauche.",
		},
		{
			Text:          "Investir en bourse",
			AIChoice:      "right",
			AIExplanation: "L'investissement boursier est traditionnellement associé au capitalisme et aux valeurs économiques libérales, plutôt liées à la droite.",
		},
		{
			Text:          "Faire du bénévolat",
			AIChoice:      "left",
			AIExplanation: "Le bénévolat et l'engagement solidaire sont souvent associés aux valeurs d'entraide et de justice sociale, plutôt liées à la gauche.",
		},
		{
			Text:          "Collectionner des montres de luxe",
			AIChoice:      "right",
			AIExplanation: "La collection d'objets de luxe est associée à la richesse et au statut social, valeurs plus traditionnellement liées à la droite.",
		},
	}

	// Insertion des questions
	for _, stmt := range statements {
		if err := db.Create(&stmt).Error; err != nil {
			log.Printf("Erreur lors de l'insertion de '%s': %v", stmt.Text, err)
		} else {
			log.Printf("Question ajoutée: %s", stmt.Text)
		}
	}

	log.Println("Base de données initialisée avec succès!")
}