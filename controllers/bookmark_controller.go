package controllers

import (
	"browsafe_backend/models"
	"browsafe_backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//-----------------------------------add bookmark controller --------------------------------------

func AddBookmark(ctx *gin.Context) {
	userID := ctx.Param("id")
	var bookmark models.Bookmark
	if err := ctx.ShouldBind(&bookmark); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format: " + err.Error(),
		})
		return
	}
	if bookmark.URL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is required",
		})
		return
	}
	createdBookmark, err := services.AddBookmarkService(userID, bookmark)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    createdBookmark,
	})
}

//-----------------------------------get bookmarks controller --------------------------------------

func GetBookmarks(ctx *gin.Context) {
	userID := ctx.Param("id")
	bookmarks, err := services.GetBookmarksByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bookmarks,
	})
}
