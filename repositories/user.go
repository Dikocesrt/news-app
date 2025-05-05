package repositories

import (
	"errors"
	"makanan-app/entities"
	"makanan-app/models"
	"makanan-app/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (r UserRepository) Register(user entities.User) (entities.User, error) {
	var userDB models.User

	if err := r.DB.Where("username = ?", user.Username).First(&userDB).Error; err == nil {
		return entities.User{}, utils.ErrUsernameAlreadyRegistered
	}

	userDB = models.NewUser(user.Password, user.Username)
	if err := r.DB.Create(&userDB).Error; err != nil {
		return entities.User{}, errors.New("failed to register user")
	}

	return entities.User{
		Password: userDB.Password,
		Username: userDB.Username,
	}, nil
}

func (r UserRepository) Login(user entities.User) (entities.User, error) {
	var userDB models.User

	if err := r.DB.Where("username = ?", user.Username).First(&userDB).Error; err != nil {
		return entities.User{}, utils.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		return entities.User{}, utils.ErrInvalidCredentials
	}

	return entities.User{
		Password: userDB.Password,
		Username: userDB.Username,
	}, nil
}