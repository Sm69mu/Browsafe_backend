package routes

import (
	"browsafe_backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello server",
		})
	})

	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.POST("/google", controllers.HandleGoogleSignIn)
		userGroup.GET("/getuser/:id", controllers.GetuserDetails)
		userGroup.PUT("/updateuser/:id", controllers.UpdateUser)
		userGroup.POST("/bookmarks/:id", controllers.AddBookmark)
		userGroup.GET("/bookmarks/:id", controllers.GetBookmarks)

		// protected := userGroup.Group("")
		// protected.Use(middlewares.AuthMiddleware())
		// {

		// }
	}

}
