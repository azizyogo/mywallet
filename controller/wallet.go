package controller

import (
	"mywallet/dto/request"
	"mywallet/middleware"
	"mywallet/pkg/httputil"
	"mywallet/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		httputil.SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	wallet, err := server.WalletUsecase.GetBalance(userID)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, wallet)
}

func TopUp(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		httputil.SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var req request.TopUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.SendError(c, http.StatusBadRequest, "Validation failed", middleware.ValidationErrorResponse(err))
		return
	}

	result, err := server.WalletUsecase.TopUp(userID, req)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, result)
}
