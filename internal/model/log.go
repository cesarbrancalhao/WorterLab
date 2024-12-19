package model

import "github.com/google/uuid"

type Log struct {
	Model
	UserID uuid.UUID `gorm:"index;foreignKey:UserID;not null"`
	WordID uuid.UUID `gorm:"index;foreignKey:WordID;not null"`
}
