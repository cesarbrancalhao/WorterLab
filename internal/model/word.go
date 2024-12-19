package model

type Word struct {
	Model
	Word        string `gorm:"uniqueIndex;not null"`
	Difficulty  int    `gorm:"default:0"`
	Translation *string
	Examples    *[]string
	UserWords   []UserWord `gorm:"foreignKey:WordID"`
}
