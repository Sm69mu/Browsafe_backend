package routes

import (
	"browsafe_backend/controllers"

	"github.com/gin-gonic/gin"
)

func BookMarkRoutes(router *gin.Engine) {
	bookMarkGroup := router.Group("/bookmarks")
	{
		bookMarkGroup.GET("/getbookmark/:id", controllers.GetBookmarks)
		bookMarkGroup.POST("/addbookmark/:id", controllers.AddBookmark)
	}
}
