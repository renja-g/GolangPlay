package router

import (
	"net/http"

	"github.com/renja-g/GolangPlay/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	userRouter := baseRouter.Group("/users")
	userRouter.POST("/signup", userController.Signup)

	return router
}
