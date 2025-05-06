package usecases

import (
	"errors"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/middlewares"
	"test-indonesia-cakap-digital/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository entities.UserRepositoryInterface
}

func NewUserUsecase(userRepository entities.UserRepositoryInterface) UserUsecase {
	return UserUsecase{
		userRepository: userRepository,
	}
}

func (u UserUsecase) Register(user entities.User) (entities.User, error) {
	if(user.Username == "" || user.Password == "") {
		return entities.User{}, utils.ErrEmptyField
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)
	
	newUser, err := u.userRepository.Register(user)
	if err != nil {
		return entities.User{}, err
	}

	token, err := middlewares.CreateToken(newUser.ID)
	if err != nil {
		return entities.User{}, errors.New("failed to create token")
	}

	newUser.Token = token

	return newUser, nil
}

func (u UserUsecase) Login(user entities.User) (entities.User, error) {
	if (user.Username == "" || user.Password == "") {
		return entities.User{}, utils.ErrEmptyField
	}

	userLogin, err := u.userRepository.Login(user)
	if err != nil {
		return entities.User{}, err
	}

	token, err := middlewares.CreateToken(userLogin.ID)
	if err != nil {
		return entities.User{}, errors.New("failed to create token")
	}

	userLogin.Token = token

	return userLogin, nil
}