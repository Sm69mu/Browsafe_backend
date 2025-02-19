package middlewares

import (
	"browsafe_backend/configs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "authorize header required",
			})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "invalid authorization header format",
			})
			ctx.Abort()
			return
		}

		//Get id token
		idToken := parts[1]

		//verify firebase token

		token, err := configs.AuthClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "invalid or expired token",
			})
			ctx.Abort()
			return
		}

		//set the verified userID in the context
		ctx.Set("userID", token.UID)
		ctx.Next()
	}
}
