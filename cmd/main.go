package main

import (
	"browsafe_backend/configs"
	"browsafe_backend/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVariables()
	configs.LoadAgoraConfigs()
	configs.InitFirebase()
}

func main() {
	r := gin.Default()
	routes.UserRoutes(r)
	routes.BookMarkRoutes(r)
	routes.VideoCallRoutes(r)
	port := configs.Port
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
