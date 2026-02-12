package controller

import (
	"mywallet/dto/request"
	"mywallet/middleware"
	"mywallet/pkg/httputil"
	"mywallet/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		httputil.SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var req request.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httputil.SendError(c, http.StatusBadRequest, "Validation failed", middleware.ValidationErrorResponse(err))
		return
	}

	result, err := server.TransactionUsecase.Transfer(userID, req)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	httputil.SendSuccess(c, http.StatusOK, result)
}

func GetHistory(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		httputil.SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	// Parse pagination params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	transactions, pagination, err := server.TransactionUsecase.GetHistory(userID, page, limit)
	if err != nil {
		middleware.HandleAppError(c, err)
		return
	}

	httputil.SendSuccessWithMeta(c, http.StatusOK, transactions, pagination)
}
