package controllers

import (
	"browsafe_backend/services"
	"hash/fnv"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCall(ctx *gin.Context) {
	channelName := ctx.Query("channelName")
	//userIDstring := ctx.Query("userID")
	//userIDuint := HashStringToUint32(userIDstring)
	if channelName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "channel name can not be empty",
		})
		return
	}

	token, err := services.GenerateAgoraToken(channelName, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	callLink := "https://browsafe.com/join?channel=" + channelName + "&token" + token

	ctx.JSON(http.StatusOK, gin.H{
		"success":     true,
		"channelName": channelName,
		"token":       token,
		"callLink":    callLink,
	})

}

func HashStringToUint32(ID string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(ID))
	return h.Sum32()
}

func JoinCall(ctx *gin.Context) {
	channelName := ctx.Query("channelName")
	if channelName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "channelname can not be empty",
		})
		return
	}

	token, err := services.GenerateAgoraToken(channelName, -0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"channelname": channelName,
		"token":       token,
	})
}
