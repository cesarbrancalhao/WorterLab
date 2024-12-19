package model

import (
	"time"

	"github.com/google/uuid"
)

type UserWord struct {
	Model
	UserID     uuid.UUID  `gorm:"index;foreignKey:UserID;not null"`
	WordID     uuid.UUID  `gorm:"index;foreignKey:WordID;not null"`
	LastSeenAt *time.Time `gorm:"default:autoUpdateTime"`
	ViewCount  int
}
