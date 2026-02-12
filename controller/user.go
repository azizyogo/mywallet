package controller

import (
	"mywallet/dto/request"
	"mywallet/middleware"
	"mywallet/pkg/httputil"
	"mywallet/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.SendError(c, http.StatusBadRequest, "Validation failed", middleware.ValidationErrorResponse(err))
		return
	}

	user, err := server.UserUsecase.Register(req)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	httputil.SendSuccess(c, http.StatusCreated, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.SendError(c, http.StatusBadRequest, "Validation failed", middleware.ValidationErrorResponse(err))
		return
	}

	userResp, err := server.UserUsecase.Login(req)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, userResp)
}

func GetProfile(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		httputil.SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	user, err := server.UserUsecase.GetProfile(userID)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, gin.H{
		"user": user,
	})
}
