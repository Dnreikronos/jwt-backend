package migration

import (
	"github.com/Dnreikronos/jwt-backend/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	createTables(db)
}

func createTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
