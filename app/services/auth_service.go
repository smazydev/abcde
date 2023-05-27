package services

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/smazydev/abcde/app/repositories"
	"github.com/smazydev/abcde/app/utils"
)

var secretKey = "your-secret-key"

type AuthService interface {
	Login(credentials Credentials) (string, error)
	ValidateJWT(token string) (string, error)
}

type Claims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(credentials Credentials) (string, error) {
	// Validate credentials
	user, err := s.userRepo.GetByEmail(credentials.Email)
	if err != nil {
		return "", err
	}
	if user == nil || user.Password != credentials.Password {
		return "", ErrInvalidCredentials
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *authService) ValidateJWT(token string) (string, error) {
	// Parse and validate the token
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	// Extract the user ID from the claims and return it
	return claims.UserID, nil
}
