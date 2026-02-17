package user

import (
	"mywallet/apperror"
	"mywallet/dto/request"
	"mywallet/dto/response"
	"mywallet/model"
	"mywallet/shared/utils/auth"
	"mywallet/shared/utils/converter"
	"mywallet/shared/utils/hash"
)

func (uc *UserUsecase) Register(req request.RegisterRequest) (*response.UserResponse, error) {
	// Check if user already exists
	existing, _ := uc.u.FindByEmail(req.Email)
	if existing != nil {
		return nil, apperror.ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &model.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}
	if err := uc.u.Create(user); err != nil {
		return nil, err
	}

	// Create wallet for the user
	_, err = uc.w.CreateWallet(user.ID)
	if err != nil {
		return nil, err
	}

	userResp := converter.ModelUserToResponse(user)
	return &userResp, nil
}

func (uc *UserUsecase) Login(req request.LoginRequest) (*response.AuthResponse, error) {
	// Get user by email
	user, err := uc.u.FindByEmail(req.Email)
	if err != nil {
		return nil, apperror.ErrInvalidCredentials
	}

	// Verify password
	if !hash.VerifyPassword(user.PasswordHash, req.Password) {
		return nil, apperror.ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.ID, user.Email, uc.cfg.JWTSecret, uc.cfg.JWTExpirationHours)
	if err != nil {
		return nil, err
	}

	userResp := converter.ModelUserToResponse(user)
	return &response.AuthResponse{
		Token: token,
		User:  userResp,
	}, nil
}

func (uc *UserUsecase) GetProfile(userID uint) (*response.UserResponse, error) {
	user, err := uc.u.FindByID(userID)
	if err != nil {
		return nil, apperror.ErrUserNotFound
	}

	userResp := converter.ModelUserToResponse(user)
	return &userResp, nil
}

func (uc *UserUsecase) ValidateToken(tokenString string) (*auth.JWTClaims, error) {
	return auth.ParseJWT(tokenString, uc.cfg.JWTSecret)
}
