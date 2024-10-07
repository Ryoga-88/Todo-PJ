package usecase

import (
	"time"

	"github.com/Ryoga-88/Todo-PJ/backend/config"
	"github.com/Ryoga-88/Todo-PJ/backend/entity"
	"github.com/Ryoga-88/Todo-PJ/backend/repository"
	"github.com/Ryoga-88/Todo-PJ/backend/validator"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user entity.User) (entity.UserResponse, error)
	Login(user entity.User) (string, error) // JWT token
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user entity.User) (entity.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return entity.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return entity.UserResponse{}, err
	}
	newUser := entity.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return entity.UserResponse{}, err
	}
	resUser := entity.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user entity.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := entity.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Conf.SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
