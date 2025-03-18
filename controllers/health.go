package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	models "hangover/models/utils"
	"log"
	"net/http"
)

type HealthRepo struct{}

func (h HealthRepo) Status(c *gin.Context) {
	plainPassword := "admin"

	// Hash the password
	hashedPassword, err := models.HashPassword(plainPassword)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	fmt.Println("Hashed Password:", hashedPassword)

	// Validate the password (correct one)
	isValid := models.ValidatePassword(hashedPassword, plainPassword)
	if isValid {
		fmt.Println("Password is valid!")
	} else {
		fmt.Println("Invalid password!")
	}

	// Validate with a wrong password
	isValid = models.ValidatePassword(hashedPassword, "wrongpassword")
	if isValid {
		fmt.Println("Password is valid!")
	} else {
		fmt.Println("Invalid password!")
	}

	c.String(http.StatusOK, "Gin-starter is working")
}
