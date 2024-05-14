package services

import (
	"github.com/renja-g/GolangPlay/data/request"
	"github.com/renja-g/GolangPlay/helper"
	"github.com/renja-g/GolangPlay/models"
	"github.com/renja-g/GolangPlay/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) Create(user request.SignupUserRequest) {
	err := s.Validate.Struct(user)
	helper.ErrorPanic(err)
	hashedPassword, err := helper.HashPassword(user.Password)
	helper.ErrorPanic(err)
	userModel := models.User{
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}
	s.UserRepository.Save(userModel)
}
