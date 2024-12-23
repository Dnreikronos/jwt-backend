package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func FilteredResponse(user User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

type UserResponse struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Email string    `json:"email,omitempty"`
}

func (u *User) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
