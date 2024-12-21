package models

import "github.com/google/uuid"


type User struct{
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
