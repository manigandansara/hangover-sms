package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hangover/models"
	"hangover/structs"
	"hangover/utils"
)

type AuthRepo struct{}

func (a AuthRepo) Login(c *gin.Context) {
	var loginRequest structs.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		// If validation fails, return a bad request status with the error details
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, e := range validationErrors {
				// Collect all validation error messages
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' %s", e.Field(), e.Tag()))
			}
			utils.ErrorResponse(c, "Validation error", 400, errorMessages)
		} else {
			// If it’s another error (not validation error)
			utils.ErrorResponse(c, "Invalid request", 400, err.Error())
		}
		return
	}

	user, err := models.ValidateUserByUserNameAndPassword(loginRequest.Username, loginRequest.Password)
	if err != nil {
		utils.ErrorResponse(c, "Invalid username or password", 400, err.Error())
		return
	}

	Token, err := utils.GenerateJWT(user)
	if err != nil {
		utils.ErrorResponse(c, "Error generating token", 500, err.Error())
		return
	}
	utils.SuccessResponse(c, "Login successful", 200, Token)
	return
}

func (a AuthRepo) Logout(c *gin.Context) {}

func (a AuthRepo) Register(c *gin.Context) {}

func (a AuthRepo) Refresh(c *gin.Context) {}

func (a AuthRepo) Verify(c *gin.Context) {}

func (a AuthRepo) Forgot(c *gin.Context) {}

func (a AuthRepo) Reset(c *gin.Context) {}
