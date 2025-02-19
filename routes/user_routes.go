package routes

import (
	"browsafe_backend/controllers"
	"browsafe_backend/middlewares"
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

		protected := userGroup.Group("")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/getuser/:id", controllers.GetuserDetails)
			protected.PUT("/updateuser/:id", controllers.UpdateUser)
			protected.POST("/bookmarks/:id", controllers.AddBookmark)
			protected.GET("/bookmarks/:id", controllers.GetBookmarks)
		}
	}

}
