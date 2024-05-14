package repository

import (
	"github.com/renja-g/GolangPlay/helper"
	"github.com/renja-g/GolangPlay/models"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &userRepositoryImpl{Db: Db}
}

func (r *userRepositoryImpl) Save(user models.User) {
	result := r.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}