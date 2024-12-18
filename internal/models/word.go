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
	/* Examples    []string  `json:"examples"`
	 * Currently disabled because I can't find a free API that
	 * returns a list or at least a random example.
	 */
	CreatedAt time.Time
}
