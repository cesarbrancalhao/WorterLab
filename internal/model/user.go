package model

type User struct {
	Model
	Email     string     `gorm:"uniqueIndex;not null"`
	Password  string     `gorm:"not null"`
	UserWords []UserWord `gorm:"foreignKey:UserID"`
}
