package models

import (
	"time"

	"github.com/google/uuid"
)

type Word struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Word        string    `json:"word"`
	Translation string    `json:"translation"`
	Difficulty  int       `json:"difficulty"`
	Examples    []string  `json:"examples"`
	CreatedAt   time.Time
}
