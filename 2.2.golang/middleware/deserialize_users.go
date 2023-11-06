package middleware

import (
	"fmt"
	"golang-crud-gin/helper"
	"golang-crud-gin/repository"
	"golang-crud-gin/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//

// func DeserializeUser(userRepository repository.UsersRepository) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var token string
// 		authorizationHeader := ctx.Request.Header.Get("Authorization")
// 		fields := strings.Fields(authorizationHeader)

// 		if len(fields) != 2 || fields[0] != "Bearer" {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid or missing Authorization header"})
// 			return
// 		}

// 		token = fields[1]

// 		if token == "" {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
// 			return
// 		}

// 		sub, err := utils.ValidateToken(token, token)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
// 			return
// 		}

// 		id, err_id := strconv.Atoi(fmt.Sprint(sub))
// 		helper.ErrorPanic(err_id)
// 		result, err := userRepository.FindById(id)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no longer exists"})
// 			return
// 		}

// 		ctx.Set("currentUser", result.Username)
// 		ctx.Next()
// 	}
// }







func DeserializeUser(userRepository repository.UsersRepository) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			var token string
			authorizationHeader := ctx.Request.Header.Get("Authorization")
			fields := strings.Fields(authorizationHeader)
	
			if len(fields) != 0 && fields[0] == "Bearer" {
				token = fields[1]
			}
	
			if token == "" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
				return
			}
			if err := godotenv.Load("app.env"); err != nil {
				log.Fatalf("Error loading .env file: %v", err)
			}
			tokenSecret := os.Getenv("TOKEN_SECRET")
		
			sub, err := utils.ValidateToken(token, tokenSecret)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
				return
			}
	
			id, err_id := strconv.Atoi(fmt.Sprint(sub))
			helper.ErrorPanic(err_id)
			result, err := userRepository.FindById(id)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
				return
			}
	
			ctx.Set("currentUser", result.Username)
			ctx.Next()
	
		}
	}
	