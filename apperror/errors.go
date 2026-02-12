package apperror

import (
	"errors"
	"net/http"
)

type AppError struct {
	Err        error
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Err.Error()
}

func NewAppError(err error, message string, statusCode int) *AppError {
	return &AppError{
		Err:        err,
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	ErrUserAlreadyExists    = &AppError{errors.New("user exists"), "User with this email already exists", http.StatusConflict}
	ErrUserNotFound         = &AppError{errors.New("user not found"), "User not found", http.StatusNotFound}
	ErrInvalidCredentials   = &AppError{errors.New("invalid credentials"), "Invalid email or password", http.StatusUnauthorized}
	ErrWalletNotFound       = &AppError{errors.New("wallet not found"), "Wallet not found", http.StatusNotFound}
	ErrInsufficientBalance  = &AppError{errors.New("insufficient balance"), "Insufficient balance for this transaction", http.StatusConflict}
	ErrInvalidAmount        = &AppError{errors.New("invalid amount"), "Amount must be greater than zero", http.StatusBadRequest}
	ErrSelfTransfer         = &AppError{errors.New("self transfer"), "Cannot transfer to yourself", http.StatusBadRequest}
	ErrUnauthorized         = &AppError{errors.New("unauthorized"), "Unauthorized access", http.StatusUnauthorized}
	ErrForbidden            = &AppError{errors.New("forbidden"), "Access forbidden", http.StatusForbidden}
	ErrDuplicateTransaction = &AppError{errors.New("duplicate transaction"), "Duplicate transaction detected", http.StatusConflict}
	ErrOptimisticLock       = &AppError{errors.New("optimistic lock"), "Concurrent modification detected, please retry", http.StatusConflict}
)
