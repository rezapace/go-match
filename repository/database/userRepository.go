package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ReadToken(id int) (user *models.User, err error)
	GetUser() (user []models.User, err error)
	GetuserByEmail(email string) (*models.User, error)
	GetUseById(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	CreateUser(user *models.User) error
	DeleteUser(user *models.User) error
	CountUser() (res int64)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}

}

// ReadToken is a function to read token
func (u *userRepository) ReadToken(id int) (user *models.User, err error) {

	err = config.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser is a function to get all user
func (u *userRepository) GetUser() (user []models.User, err error) {
	if err := config.DB.Preload("Player.Participation").Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetUseById is a function to get user by id
func (u *userRepository) GetUseById(id int) (user *models.User, err error) {
	err = config.DB.Model(&user).Preload("Player.Participation").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CountUser is a function to count user
func (u *userRepository) CountUser() (res int64) {
	res = 0
	user := []models.User{}

	if err := config.DB.Model(&user).Count(&res).Error; err != nil {
		return 0
	}
	return res
}

// UpdateUser is a function to update user
func (u *userRepository) UpdateUser(user *models.User) error {

	if err := config.DB.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetuserByEmail is a function to get user by email
func (u *userRepository) GetuserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser is a function to create user
func (u *userRepository) CreateUser(user *models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser is a function to delete user
func (u *userRepository) DeleteUser(user *models.User) error {

	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
