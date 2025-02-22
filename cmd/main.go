package main

import (
	"browsafe_backend/configs"
	"browsafe_backend/routes"
	"os"
	"strings"

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
    
    // Get port from environment variable for Render compatibility
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000" // fallback port
    }
    
    // Ensure port has ":" prefix
    if !strings.HasPrefix(port, ":") {
        port = ":" + port
    }
    
    r.Run(port)
}