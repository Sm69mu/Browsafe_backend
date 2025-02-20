package routes

import (
	"browsafe_backend/controllers"

	"github.com/gin-gonic/gin"
)

func VideoCallRoutes(router *gin.Engine) {
	videoCallGroup := router.Group("/videoCall")
	{
		videoCallGroup.GET("/create-call", controllers.CreateCall)
		videoCallGroup.GET("/join-call", controllers.JoinCall)
	}
}