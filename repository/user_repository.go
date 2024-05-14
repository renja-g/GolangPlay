package repository

import "github.com/renja-g/GolangPlay/models"

type UserRepository interface {
	Save(user models.User)
}