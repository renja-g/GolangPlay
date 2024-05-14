package services

import "github.com/renja-g/GolangPlay/data/request"


type UserService interface {
	Create(user request.SignupUserRequest)
}
