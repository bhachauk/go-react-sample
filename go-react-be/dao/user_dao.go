package dao

import (
	"bhachauk.github.io/go-react-sample/go-react-be/config"
	"bhachauk.github.io/go-react-sample/go-react-be/models"
	"errors"
	"gorm.io/gorm"
)

// UserDAO provides methods for interacting with the User model in the database
type UserDAO struct {
	DB *gorm.DB
}

// NewUserDAO creates a new UserDAO instance
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{DB: db}
}

// CreateUser creates a new user in the database
func (dao *UserDAO) CreateUser(user *models.User) error {
	return dao.DB.Create(user).Error
}

// GetAllUsers retrieves all users from the database
func (dao *UserDAO) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := dao.DB.Find(&users).Error
	return users, err
}

// GetUserByID retrieves a user by their ID
func (dao *UserDAO) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := dao.DB.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil // Return nil if not found
	}
	return &user, err
}

// UpdateUser updates an existing user in the database
func (dao *UserDAO) UpdateUser(user *models.User) error {
	return dao.DB.Save(user).Error
}

// DeleteUser deletes a user by their ID
func (dao *UserDAO) DeleteUser(id uint) error {
	return dao.DB.Delete(&models.User{}, id).Error
}

// GetUserByEmail retrieves a user by their email
func (dao *UserDAO) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// GetUserByUsername retrieves a user by their username
func (dao *UserDAO) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}
