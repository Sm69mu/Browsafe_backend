package controllers

import (
	"browsafe_backend/models"
	"browsafe_backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type GoogleSignIn struct {
	IDToken string `json:"idtoken" binding:"required"`
}

func RegisterUser(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}
	user, err := services.RegisterUser(req.Email, req.Password, req.Name)
	if err != nil {
		statucode := http.StatusInternalServerError
		//check for error types 
		if err.Error()=="error creating Firebase user: email already exists" {
			statucode= http.StatusConflict
		}
		ctx.JSON(statucode, gin.H{
			"success":false,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func LoginUser(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if err.Error()=="invalid password" || err.Error() == "error getting user details: no user record found" {
			statusCode=http.StatusUnauthorized
			
		}
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    token,
	})
}


func HandleGoogleSignIn(ctx *gin.Context)  {
	var req GoogleSignIn
	if err:= ctx.ShouldBind(&req);err!= nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}
	user ,token, err := services.HandleGoogleSignIn(req.IDToken)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if err.Error()=="error verifying Google ID token" {
			statusCode=http.StatusUnauthorized
		}
		ctx.JSON(statusCode, gin.H{
			"success":false,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":gin.H{
			"user":user,
			"token": token,
		},
	})
}


func GetAuthUser(ctx *gin.Context)  {
	//the user ID should be set by the middleware 
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error": "user not authenticated",
		})
		return
	}
	user ,err := services.GetUserDetailsByID(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": user,
	})
}



func UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var updates models.Users
	if err := ctx.ShouldBind(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	updateduser, err := services.UpdateUserService(userID, updates)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    updateduser,
	})

}

func GetuserDetails(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "UserID is empty",
		})
		return
	}
	userDetails, err := services.GetUserDetailsByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userDetails,
	})
}
