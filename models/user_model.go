package models

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"size:255;not null" json:"first_name"`
	LastName  string `gorm:"size:255;not null" json:"last_name"`
	Email     string `gorm:"size:255;unique;not null" json:"email"`
}
