package initializers

import "github.com/renja-g/GolangPlay/models"

func SyncDatabase(){
	DB.AutoMigrate(
		&models.User{},
	)
}