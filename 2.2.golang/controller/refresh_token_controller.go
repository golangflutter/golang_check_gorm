package controller

import (
	"fmt"
	"golang-crud-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenRequest struct {
    RefreshToken string `json:"refresh_token" binding:"required"`
}


type RefreshTokenResponse struct {
    AccessToken  string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

// RefreshTokenController handles token refresh requests
func RefreshTokenController(c *gin.Context) {
    var request RefreshTokenRequest
	fmt.Println("This is RefreshTokenController 111111")
    // Bind the request data to the RefreshTokenRequest struct
	if err := c.ShouldBind(&request); err != nil {
		fmt.Println("Error binding request:", err.Error()) // Add this line for detailed error information
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}
	
	fmt.Println("This is RefreshTokenController before newAccessToken")
    // Validate and refresh the access token using the provided refresh token
    newAccessToken, err := utils.RefreshAccessToken(request.RefreshToken, "my-ultra-secure-json-web-token-string") // Replace "your-secret-key" with your actual secret key
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid refresh token",
        })
        return
    }
    fmt.Println("This is RefreshTokenController newAccessToken")
    fmt.Println(newAccessToken)
    // You can also generate a new refresh token if needed

    // Respond with the new access token and refresh token (if applicable)
    response := RefreshTokenResponse{
        AccessToken:  newAccessToken,
        RefreshToken: "", // Set the new refresh token here, if applicable
    }

    c.JSON(http.StatusOK, response)
}
