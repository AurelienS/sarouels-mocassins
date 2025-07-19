package models

import (
	"time"
)

type Statement struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Text            string    `json:"text" gorm:"not null"`
	AIChoice        string    `json:"ai_choice" gorm:"not null"`  // "left" ou "right"
	AIExplanation   string    `json:"ai_explanation" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Vote struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	StatementID uint      `json:"statement_id"`
	Choice      string    `json:"choice" gorm:"not null"` // "left" ou "right"
	CreatedAt   time.Time `json:"created_at"`
}

type Stats struct {
	StatementID   uint    `json:"statement_id"`
	LeftVotes     int     `json:"left_votes"`
	RightVotes    int     `json:"right_votes"`
	LeftPercent   float64 `json:"left_percent"`
	RightPercent  float64 `json:"right_percent"`
}