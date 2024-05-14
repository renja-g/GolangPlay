package main

import (
	"net/http"

	"github.com/renja-g/GolangPlay/config"
	"github.com/renja-g/GolangPlay/controllers"
	"github.com/renja-g/GolangPlay/helper"
	"github.com/renja-g/GolangPlay/initializers"
	"github.com/renja-g/GolangPlay/models"
	"github.com/renja-g/GolangPlay/repository"
	"github.com/renja-g/GolangPlay/router"
	"github.com/renja-g/GolangPlay/services"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	log.Info().Msg("Started Server!")
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.User{})


	// Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	// Service
	userService := services.NewUserServiceImpl(userRepository, validate)

	// Controller
	userController := controllers.NewUserController(userService)

	// Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

	/*
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth ,controllers.Validate)
	*/
}