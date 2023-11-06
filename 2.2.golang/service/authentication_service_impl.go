package service

import (
	"errors"
	"log"
	"os"

	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in the database
	newUsers, usersErr := a.UsersRepository.FindByUsername(users.Username)
	if usersErr != nil {
		return "", errors.New("invalid username or password")
	}
    if err := godotenv.Load("app.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	// Load environment variables
	tokenSecret := os.Getenv("TOKEN_SECRET")
	tokenExpiredIn := os.Getenv("TOKEN_EXPIRED_IN")

	// Check if TOKEN_SECRET and TOKEN_EXPIRED_IN are set
	if tokenSecret == "" || tokenExpiredIn == "" {
		return "", errors.New("TOKEN_SECRET or TOKEN_EXPIRED_IN environment variable not set")
	}

	// Verify the user's password
	verifyError := utils.VerifyPassword(newUsers.Password, users.Password)
	if verifyError != nil {
		return "", errors.New("invalid username or password")
	}

	// Parse the token expiration duration
	tokenDuration, err := time.ParseDuration(tokenExpiredIn)
	if err != nil {
		return "", errors.New("error parsing TOKEN_EXPIRED_IN")
	}

	// Generate the token
	token, errToken := utils.GenerateToken(tokenDuration, newUsers.Id, tokenSecret)
	if errToken != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}


// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(users request.CreateUsersRequest) {

	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UsersRepository.Save(newUser)
}
