package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/renja-g/GolangPlay/data/request"
	"github.com/renja-g/GolangPlay/data/response"
	"github.com/renja-g/GolangPlay/helper"
	"github.com/renja-g/GolangPlay/initializers"
	"github.com/renja-g/GolangPlay/models"
	"github.com/renja-g/GolangPlay/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController{
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) Signup(ctx *gin.Context){
	log.Info().Msg("Signup")
	signupUserRequest := request.SignupUserRequest{}
	err := ctx.ShouldBindJSON(&signupUserRequest)
	helper.ErrorPanic(err)

	controller.userService.Create(signupUserRequest)
	webResponse := response.Response{
		Code: http.StatusCreated,
		Status: "Success",
		Data: gin.H{
			"message": "User created successfully",
		},
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)

}

func Login (c *gin.Context){
	// Get email & pass off req body
	var body struct{
		Email string    `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Look up for requested user
	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in password with saved users password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {	
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context){
	user, _ := c.Get("user")

	// user.(models.User).Email    -->   to access specific data
	respone := gin.H{
		"id": user.(models.User).ID,
		"createdAt": user.(models.User).CreatedAt,
		"UpdatedAt": user.(models.User).UpdatedAt,
		"email": user.(models.User).Email,
		"username": user.(models.User).Username,
	}

	c.JSON(http.StatusOK, respone)
}
