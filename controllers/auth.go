package controllers

import (
	"fmt"
	"hangover/models"
	"hangover/structs"
	"hangover/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthRepo struct{}

func (a AuthRepo) Login(c *gin.Context) {
	var loginRequest structs.LoginRequest

	// Bind and validate the JSON request
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		handleValidationError(c, err)
		return
	}

	user, err := models.ValidateUserByUserNameAndPassword(loginRequest.Username, loginRequest.Password)
	if err != nil {
		utils.JsonResponse(c, "Invalid username or password", http.StatusBadRequest, err.Error())
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		utils.JsonResponse(c, "Failed to generate authentication token", http.StatusInternalServerError, nil)
		return
	}
	utils.JsonResponse(c, "Login successful", http.StatusOK, gin.H{"token": token})
}

func (a AuthRepo) Logout(c *gin.Context) {}

func (a AuthRepo) Register(c *gin.Context) {}

func (a AuthRepo) Refresh(c *gin.Context) {}

func (a AuthRepo) Verify(c *gin.Context) {}

func (a AuthRepo) Forgot(c *gin.Context) {}

func (a AuthRepo) Reset(c *gin.Context) {}

// Helper function to handle validation errors
func handleValidationError(c *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' %s", e.Field(), e.Tag()))
		}
		utils.JsonResponse(c, "Validation error", http.StatusBadRequest, errorMessages)
	} else {
		utils.JsonResponse(c, "Invalid request payload", http.StatusBadRequest, err.Error())
	}
}
