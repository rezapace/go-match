package faker

import (
	"PongPedia/constants"
	"PongPedia/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	passwordhash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	return &models.User{
		Username: "admin",
		Email:    "admin@gmail.com",
		Password: string(passwordhash),
		Role:     constants.ADMIN,
	}
}
